#!/usr/bin/env bash

set -euo pipefail

function get_deps() {
	# Generates a list of any OS level dependencies that need to be installed
	# as well as any additional repositories to install those

	local ostype="${1}"
	local dep_fd="${2}"
	local repo_fd="${3}"

	local deps=()
	local repos=()

	# https://github.com/ggordonklaus/portaudio
	# https://github.com/go-gl/glfw
	# https://github.com/giorgisio/goav
	case "${ostype}" in
	debian)
		deps+=(
			# portaudio
			portaudio19-dev

			# glfw
			libgl1-mesa-dev
			xorg-dev

			# goav
			libav{codec,device,filter,format,util}-dev
			libsw{resample,scale}-dev
		)
		;;
	fedora)
		deps+=(
			# portaudio
			portaudio-devel

			# glfw
			libX{11,cursor,randr,inerama,i}-devel
			mesa-libGL-devel

			# goav
			ffmpeg-devel
		)
		repos+=(
			# ffmpeg-devel
			https://download1.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm
		)
		;;

	Windows_NT)
		deps+=(
			# goav
			ffmpeg-latest-win64-dev:libav{codec,device,filter,format,util}
			ffmpeg-latest-win64-dev:libsw{resample,scale}
			ffmpeg-latest-win64-dev:swresample
		)
		;;
	esac

	echo ${deps[@]} >&${dep_fd}
	echo ${repos[@]} >&${repo_fd}
}

function detect_os() {
	[ -f /etc/os-release ] && source /etc/os-release

	local ostype="${ID:-}${OS:-}"

	echo ${ostype}
}

function install_deps() {
	ostype=${1}

	exec {dep_fd}<>$(mktemp)
	exec {repo_fd}<>$(mktemp)

	get_deps ${ostype} ${dep_fd} ${repo_fd}

	deps=$(</dev/fd/${dep_fd})
	repos=$(</dev/fd/${repo_fd})

	case "${ostype:-}" in
	debian)
		apt-get install --no-install-recommends ${deps[@]}
		;;
	fedora)
		dnf install -y ${repos[@]}
		dnf install -y --setopt=install_weak_deps=False --best ${deps[@]}
		;;
	Windows_NT)
		mkdir -p $PKG_CONFIG_PATH

		for i in ${deps[@]}; do
			IFS=: read prefix lib <<<${i}
			(
				cat <<-EOM
					prefix=/tmp/${prefix}
					name=${lib}
				EOM
				echo
				cat <<-'EOM'
					exec_prefix=${prefix}
					includedir=${prefix}/include
					libdir=${exec_prefix}/lib

					Name: ${name}
					Description: ${name}
					Version: 1.0.0
					Cflags: -I${includedir}/${name}
					Libs: -L${libdir} -l${name}
				EOM
			) >${PKG_CONFIG_PATH}/${lib}.pc
		done
		;;
	esac
}

os=$(detect_os)
install_deps ${os}
