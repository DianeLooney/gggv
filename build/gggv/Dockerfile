# Base Image with OS level dependencies
FROM golang:1.13 as base

RUN apt-get update \
    && apt-get -y install --no-install-recommends \
    # begin github.com/giorgisio/goav
    autoconf \
    automake \
    build-essential \
    libass-dev \
    libavcodec-dev \ 
    libavdevice-dev \ 
    libavfilter-dev \ 
    libavformat-dev \ 
    libavutil-dev \
    libfreetype6-dev \
    libsdl1.2-dev \
    libswresample-dev \ 
    libswscale-dev \ 
    libtheora-dev \
    libtool \
    libva-dev \
    libvdpau-dev \
    libvorbis-dev \
    libxcb-shm0-dev \
    libxcb-xfixes0-dev \
    libxcb1-dev \
    pkg-config \
    texi2html \
    yasm \
    zlib1g-dev \
    # begin github.com/go-gl/gl
    libgl1-mesa-dev \
    xorg-dev \
    # begin github.com/ggordonklaus/portaudio
    portaudio19-dev \
    # cleanup
    && rm -rf /var/list/apt/lists/*

ENV FFMPEG_ROOT=$HOME/ffmpeg

RUN mkdir -p ${FFMPEG_ROOT} && git clone https://github.com/FFmpeg/FFmpeg ${FFMPEG_ROOT}

WORKDIR ${FFMPEG_ROOT}

ENV LD_LIBRARY_PATH=${FFMPEG_ROOT}/lib \
    CGO_LDFLAGS="-L${FFMPEG_ROOT}/lib -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter" \
    CGO_CFLAGS="-I${FFMPEG_ROOT}/include" 

RUN ./configure && make && make install

### Main Image
FROM base

WORKDIR /go/src/app
COPY . .

ENV GO111MODULE=on
RUN go get -d -v ./...
RUN go build cmd/daemon/main.go

ENV DISPLAY=:0
