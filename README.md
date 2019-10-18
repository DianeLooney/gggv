# GO GO GADGET: VIDEO [![CircleCI](https://circleci.com/gh/DianeLooney/gggv/tree/master.svg?style=svg)](https://circleci.com/gh/DianeLooney/gggv/tree/master)

OSC based video editing environment. Not intended for use as a package.

## How to build

### Linux

Follow along in the dockerfile at `build/gggv-base` to install the dependencies from source, and `build/gggv` to install gggv.

### OSX

```bash
go get github.com/dianelooney/gggv
brew install ffmpeg

go run cmd/shadertest/main.go # for a simple shader live-reload environment
go run cmd/daemon/main.go # for the full OSC-controlled daemon
```

### Windows

Binaries are available through CircleCI as a build artifact. View build artifacts by appending `#artifacts/containers/0` to the CircleCI link of a `win-build` job. It is left to the user to install the various ffmpeg dlls. If you wish to build from source, follow along with the `win-build` entry in `.circleci/config.yml`.

## Frontends

[Clojure](https://github.com/DianeLooney/gggv-clojure)
