# VHS

`vhs` is used to record OSC messages and to replay them at a later date

There are inputs and outputs.

Inputs currently supported are udp and recordings.

Outputs currently supported are udp and recordings.

## Examples

Record data to only stdout, `iu` defines an *i*nput of *u*dp

* `vhs iu :4201`

Simple passthrough, `ou` defines an *o*utput of *u*dp

* `vhs iu :4201 ou :4200`

Recording for later, `of` defines an *o*utput of *f*ile.

* `vhs iu :4201 ou :4200 of recording.txt`

List whats in a file to stdout, `if` defines an *i*nput of *f*ile.

* `vhs if recording.txt`

Replay contents of a file to udp

* `vhs if recording.txt ou :4200`
