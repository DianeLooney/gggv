mkdir $PKG_CONFIG_PATH


( echo "prefix=/tmp/ffmpeg-latest-win64-dev"; echo "name=libavcodec" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libavcodec.pc;
( echo "prefix=/tmp/ffmpeg-latest-win64-dev"; echo "name=libavfilter" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libavfilter.pc;
( echo "prefix=/tmp/ffmpeg-latest-win64-dev"; echo "name=libavutil" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libavutil.pc;
( echo "prefix=/tmp/ffmpeg-latest-win64-dev"; echo "name=libswscale" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libswscale.pc;
( echo "prefix=/tmp/ffmpeg-latest-win64-dev"; echo "name=libavdevice" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libavdevice.pc;
( echo "prefix=/tmp/ffmpeg-latest-win64-dev"; echo "name=libavformat" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libavformat.pc;
( echo "prefix=/tmp/ffmpeg-latest-win64-dev"; echo "name=libswresample" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/libswresample.pc;
( echo "prefix=/tmp/ffmpeg-latest-win64-dev"; echo "name=swresample" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/swresample.pc;
# ( echo "prefix=/tmp/portaudio"; echo "name=portaudio-2.0" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/portaudio-2.0.pc;
