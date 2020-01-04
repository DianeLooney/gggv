FROM diane/gggv:latest
### Cross Compilation
RUN apt-get install -y mingw-w64 unzip
RUN curl https://ffmpeg.zeranoe.com/builds/win64/dev/ffmpeg-latest-win64-dev.zip > ${HOME}/ffmpeg-latest-win64-dev.zip
RUN ls ${HOME}
RUN unzip -o ${HOME}/ffmpeg-latest-win64-dev.zip -d ${HOME}


ENV CGO_ENABLED=1 \
    CC="/usr/bin/x86_64-w64-mingw32-gcc" \
    GOOS=windows \
    GOARCH=amd64 \
    CGO_LDFLAGS="-L/root/ffmpeg-latest-win64-dev/lib -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter -L/usr/x86_64-w64-mingw32/lib" \
    CGO_CFLAGS="-I/root/ffmpeg-latest-win64-dev/include -w" \
    LD_LIBRARY_PATH="/root/ffmpeg-latest-win64-dev/lib"
RUN go build cmd/daemon/main.go
