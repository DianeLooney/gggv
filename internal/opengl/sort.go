package opengl

import (
	"strings"

	"github.com/dianelooney/gggv/internal/errors"
)

type sorter struct {
	sources   map[SourceName]Source
	ordered   map[SourceName]bool
	loopCheck map[SourceName]bool
	ordering  []SourceName
}

func (s *sorter) order(name SourceName) error {
	this := s.sources[name]
	if this == nil {
		if name == SourceName("-") {
			return nil
		}
		return errors.SourceMissing(name)
	}

	if s.ordered[name] {
		return nil
	}
	if s.loopCheck[name] {
		return s.genError(name)
	}

	s.loopCheck[name] = true
	for _, child := range this.Children() {
		if err := s.order(child); err != nil {
			return err
		}
	}
	s.ordering = append(s.ordering, name)
	s.ordered[name] = true

	return nil
}

func (s *sorter) genError(name SourceName) error {
	loop := []string{string(name)}
	visitedNames := make(map[SourceName]bool, len(s.sources))
	for {
		next := s.sources[name]
		for _, child := range next.Children() {
			if s.ordered[child] {
				continue
			}
			if !s.loopCheck[child] {
				continue
			}

			loop = append(loop, string(child))

			if !visitedNames[child] {
				visitedNames[name] = true
				name = child
				break
			}

			names := strings.Join(loop, " → ")
			return errors.SourceDependencyLoop(names)
		}
	}
}

func (s *sorter) checkDeps(name SourceName) error {
	checked := make(map[SourceName]bool, len(s.sources))
	sources := []SourceName{name}
	for _, name := range sources {
		if checked[name] {
			continue
		}
		source := s.sources[name]

		for _, child := range source.Children() {
			sources = append(sources, child)
			if s.sources[child] != nil {
				continue
			}

			var srcs []SourceName
			for src := range s.sources {
				srcs = append(srcs, src)
			}
			return errors.SourceMissing(name, child, child, srcs)
		}
		checked[name] = true
	}
	return nil
}

func Order(target SourceName, sources map[SourceName]Source) (order []SourceName, err error) {
	s := sorter{
		sources:   sources,
		ordered:   make(map[SourceName]bool, len(sources)),
		loopCheck: make(map[SourceName]bool, len(sources)),
		ordering:  nil,
	}
	err = s.checkDeps(target)
	if err != nil {
		return nil, err
	}
	err = s.order(target)
	order = s.ordering
	return
}
