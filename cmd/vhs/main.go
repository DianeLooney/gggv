package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/hypebeast/go-osc/osc"
)

type Msg struct {
	Time time.Time
	Addr string
	Args []interface{}
}

var exit sync.WaitGroup

func inRedis(s string) <-chan Msg {
	pw := os.Getenv("REDIS_AUTH")
	client := redis.NewClient(&redis.Options{
		Addr:     s,
		Password: pw, // no password set
		DB:       0,  // use default DB
	})
	sub := client.Subscribe("osc")
	ch := make(chan Msg)

	go func() {
		defer fmt.Println("Done with things")

		for {
			msg, err := sub.ReceiveMessage()
			if err != nil {
				panic(err)
			}
			var m Msg
			if err := json.Unmarshal([]byte(msg.Payload), &m); err != nil {
				panic(err)
			}
			ch <- m
		}
	}()

	return ch
}

func outRedis(s string) chan<- Msg {
	pw := os.Getenv("REDIS_AUTH")
	redis.NewClient(&redis.Options{})
	client := redis.NewClient(&redis.Options{
		Addr:     s,
		Password: pw, // no password set
		DB:       0,  // use default DB
	})
	ch := make(chan Msg)
	exit.Add(1)

	go func() {
		defer exit.Done()

		for {
			m, ok := <-ch
			if !ok {
				fmt.Println("Done")
				return
			}
			data, err := json.Marshal(m)
			if err != nil {
				log.Fatalln(err)
				fmt.Println("Couldnt marshal json")
				return
			}

			if err := client.Publish("osc", string(data)); err.Err() != nil {
				fmt.Printf("Couldn't publish:\n%v\n", err)
				log.Fatalln(err)
			}
		}
	}()

	return ch
}

func inUDP(s string) <-chan Msg {
	ch := make(chan Msg)

	d := osc.NewStandardDispatcher()
	d.AddMsgHandler("*", func(msg *osc.Message) {
		ch <- Msg{
			Time: time.Now(),
			Addr: msg.Address,
			Args: msg.Arguments,
		}
	})
	srv := osc.Server{Addr: s, Dispatcher: d}
	go srv.ListenAndServe()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			close(ch)
		}
	}()

	return ch
}

func outUDP(s string) chan<- Msg {
	fmt.Println(s)
	exit.Add(1)
	strs := strings.Split(s, ":")
	n, err := strconv.Atoi(strs[1])
	if err != nil {
		panic(err)
	}

	cl := osc.NewClient(strs[0], n)
	ch := make(chan Msg)

	go func() {
		defer exit.Done()

		for m := range ch {
			err := cl.Send(osc.NewMessage(m.Addr, m.Args...))
			if err != nil {
				fmt.Fprintf(os.Stderr, "unable to send: %s\n", err)
			}
		}
	}()

	return ch
}

func inFile(s string) <-chan Msg {
	ch := make(chan Msg)

	data, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}
	msgs := make([]Msg, 0)
	if err := json.Unmarshal(data, &msgs); err != nil {
		panic(err)
	}

	start := time.Now()
	base := msgs[0].Time

	go func() {
		for _, m := range msgs {
			diff := m.Time.Sub(base)
			<-time.After(time.Until(start.Add(diff)))
			ch <- m
		}
		close(ch)
	}()

	return ch
}

func outFile(s string) chan<- Msg {
	exit.Add(1)
	ch := make(chan Msg)

	var bundle []Msg
	go func() {
		for m := range ch {
			bundle = append(bundle, m)
		}
	}()

	go func() {
		defer exit.Done()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		data, err := json.Marshal(bundle)
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(s, data, 0666); err != nil {
			fmt.Fprintf(os.Stderr, "unable to write to file: %v\n", err)
		}
	}()

	return ch
}

func outStdout() chan<- Msg {
	exit.Add(1)
	ch := make(chan Msg)

	go func() {
		defer exit.Done()

		for m := range ch {
			fmt.Printf("%s\t%s\t%v\n", m.Time.Format("15:04:05.999"), m.Addr, m.Args)
		}
	}()

	return ch
}

func main() {
	args := os.Args[1:]
	var ins []<-chan Msg
	var outs []chan<- Msg

	for len(args) > 0 {
		kind, target := args[0], args[1]
		args = args[2:]

		switch kind {
		case "if":
			ins = append(ins, inFile(target))
		case "of":
			outs = append(outs, outFile(target))
		case "iu":
			ins = append(ins, inUDP(target))
		case "ou":
			outs = append(outs, outUDP(target))
		case "ir":
			ins = append(ins, inRedis(target))
		case "or":
			outs = append(outs, outRedis(target))
		default:
			panic("unrecognized handler " + kind)
		}
	}

	outs = append(outs, outStdout())

	ch := make(chan Msg)
	wg := sync.WaitGroup{}
	wg.Add(len(ins))

	for _, in := range ins {
		go func(in <-chan Msg) {
			defer wg.Done()
			for {
				x, ok := <-in
				if !ok {
					return
				}
				ch <- x
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for x := range ch {
		for _, out := range outs {
			out <- x
		}
	}

	for _, out := range outs {
		close(out)
	}

	exit.Wait()
}
