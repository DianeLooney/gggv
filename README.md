# GO GO GADGET: VIDEO [![CircleCI](https://circleci.com/gh/DianeLooney/gvd/tree/master.svg?style=svg)](https://circleci.com/gh/DianeLooney/gvd/tree/master)

Command-line video editing environment. Not intended for use as a package.

How to build:

* Linux: follow along in the dockerfile at `build/gvd-base` to install the dependencies from source, and `build/gvd` to install gvd.

* OSX: builds are similar to the Linux builds, but exact steps are not given.

* Windows: Binaries are available through CircleCI as a build artifact. View build artifacts by appending `#artifacts/containers/0` to the CircleCI link of a `win-build` job. It is left to the user to install the various ffmpeg dlls. If you wish to build from source, follow along with the `win-build` entry in `.circleci/config.yml`.
