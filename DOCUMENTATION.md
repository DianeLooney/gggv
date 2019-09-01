# Overview
This is the documentation on the conceptual organization of this tool as well as the OSC endpoints that allow you to control it. YOu can hopefully use this as a reference for what is happening in the bindings, or as a guid on how to make your own binding

## Basic Architecture
You can imagine gggv as a video synthesizer. You push it's hypothetical buttons and feed it video files and out comes pixels on your screen!


The core logic of the application is writen in Go.
Media decoding is provided by dynamically linking to [FFmpeg](https://ffmpeg.org).
Filters are writen in GLSL.

The basic design of gggv is a server which recieves commands from a client and then produces video. It is heavily inspired by the architecture of SuperCollider's SCSynth and SCLang.

Other influences include: [media lovin toolkit](https://mltframwork.org), [Olive](https://www.olivevideoeditor.org), and Kodelife. We love these tools dearly, you should use them too!

## Sources, Filters
### ffvideo.source
gggv treats each video a one "ffvideo.source" you can think of this a factory that turns video files on your hard drive into pixels you can manipulate and put on your screen.

### shader.source
Where do the pixels go when they roll off that factory floor? Well into a Shader of course! 
For our purposes you can consider a shader to be a program that runs once for each pixel of each frame. 

So the ffv

## Intro to the Media Graph and OSC
It is helpful to think of all this pixel passing in terms of a graph or flow chart. This is called our Media Graph. 
The Media Graph is conceptually very similar to Blender's material graphs, Unity's Shader Graph, or the graph in a compositing program like Nuke or Natron.
All ffvideo.source and shader.source are each a node on this graph and when you connect the output of one source to the input of a shader.source you are drawing an arrow between the two.
Lets look at a basic example:
```
 +-myVideo-+     +-default------+     +-window------+
 |video.mp4|---->|default shader|---->|window shader|
 +---------+     +--------------+     +-------------+
```
In this graph we have on video source node which is making pixels and passing them to the default shader node (the built in default shader does nothing to the video) and then out to the window shader (for our purposes we can ignore the fact that window is a shader).
On the top of the box you can see the name of each node. So if we want to reference a node later we need to know it's name.
Let's find out what messages we would need to send our server to get to this state.
```
/source.ffvideo/create ,ss myVideo /path/to/video.mp4
/program/create ,sss defShader /shaders/vert/default.glsl /shaders/frag/default.glsl
/source.shader/create ,s default
/source.shader/set/input ,sis default 0 myVideo
/source.shader/set/input ,sis window 0 default
```
If you were to send this exact text to UDP port 4200 on the machine running the server you would find that the server would begin to play the video!! Assuming that the server computer has a video at that path, and you have the daemon installed with the shaders folder it came with.
you might accomplish this test with netcat for example. However, very few people are willing to hand write all that OSC. Generally you would use a library in your language of choice. So fo the remainder of this guide we will write our messages in shorthand. If you would like a full list of osc messages that documentation is soon to come. (for now go look at /cmd/daemon/main.go)

