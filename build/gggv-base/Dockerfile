FROM golang:1.13

# begin github.com/giorgisio/goav
RUN apt-get update
RUN apt-get -y install autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev
RUN apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev
RUN apt-get install yasm

ENV FFMPEG_ROOT=$HOME/ffmpeg
ENV CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"
ENV CGO_CFLAGS="-I$FFMPEG_ROOT/include"
ENV LD_LIBRARY_PATH=$HOME/ffmpeg/lib
RUN cd $HOME/
RUN git clone https://github.com/FFmpeg/FFmpeg ffmpeg
RUN cd ffmpeg && ./configure && make && make install
# end

# begin github.com/go-gl/gl
RUN apt-get install -y libgl1-mesa-dev xorg-dev
# end

# begin github.com/ggordonklaus/portaudio
RUN apt-get install -y portaudio19-dev
# end
