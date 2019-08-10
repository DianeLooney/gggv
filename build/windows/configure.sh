mkdir ~/pkg-config

( echo "name=libavcodec" ; cat base.pc ) > $PKG_CONFIG_PATH/libavcodec.pc;
( echo "name=libavfilter" ; cat base.pc ) > $PKG_CONFIG_PATH/libavfilter.pc;
( echo "name=libavutil" ; cat base.pc ) > $PKG_CONFIG_PATH/libavutil.pc;
( echo "name=libswscale" ; cat base.pc ) > $PKG_CONFIG_PATH/libswscale.pc;
( echo "name=libavdevice" ; cat base.pc ) > $PKG_CONFIG_PATH/libavdevice.pc;
( echo "name=libavformat" ; cat base.pc ) > $PKG_CONFIG_PATH/libavformat.pc;
( echo "name=libswresample" ; cat base.pc ) > $PKG_CONFIG_PATH/libswresample.pc;
( echo "name=swresample" ; cat base.pc ) > $PKG_CONFIG_PATH/swresample.pc;

curl https://ffmpeg.zeranoe.com/builds/win64/dev/ffmpeg-latest-win64-dev.zip > $HOME/ffmpeg-latest-win64-dev.zip
unzip $HOME/ffmpeg-latest-win64-dev.zip -d $HOME/
