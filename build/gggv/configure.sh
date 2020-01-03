#!/usr/bin/env bash

set -euo pipefail

libs=(
	# github.com/ggordonklaus/portaudio
	portaudio19-dev

	# github.com/go-gl/gl
	libgl1-mesa-dev
	xorg-dev

	# github.com/giorgisio/goav
	libavcodec-dev
	libavformat-dev
	libavutil-dev
	libavfilter-dev
	libavdevice-dev
	libswresample-dev
	libswscale-dev
)

apt-get install --no-install-recommends ${libs[@]}
