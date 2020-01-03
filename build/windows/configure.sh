set -euo pipefail

mkdir -p $PKG_CONFIG_PATH

declare -A libs

# github.com/giorgisio/goav
libs["ffmpeg-latest-win64-dev"]="
	libavcodec
	libavfilter
	libavutil
	libavdevice
	libavformat
	libswscale
	libswresample
	swresample
"

for prefix in "${!libs[@]}"; do
	for lib in ${libs[$prefix]}; do
		cat <<-EOM >${PKG_CONFIG_PATH}/${lib}.pc
			prefix=/tmp/${prefix}
			name=${lib}

			$(<build/windows/base.pc)
		EOM
	done
done

# ( echo "prefix=/tmp/portaudio"; echo "name=portaudio-2.0" ; cat build/windows/base.pc ) > $PKG_CONFIG_PATH/portaudio-2.0.pc;
