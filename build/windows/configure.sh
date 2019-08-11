mkdir /tmp/pkg-config

( echo "name=libavcodec" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libavcodec.pc;
( echo "name=libavfilter" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libavfilter.pc;
( echo "name=libavutil" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libavutil.pc;
( echo "name=libswscale" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libswscale.pc;
( echo "name=libavdevice" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libavdevice.pc;
( echo "name=libavformat" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libavformat.pc;
( echo "name=libswresample" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libswresample.pc;
( echo "name=swresample" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/swresample.pc;

curl https://ffmpeg.zeranoe.com/builds/win64/dev/ffmpeg-latest-win64-dev.zip > /tmp/ffmpeg-latest-win64-dev.zip
unzip /tmp/ffmpeg-latest-win64-dev.zip -d /tmp/
