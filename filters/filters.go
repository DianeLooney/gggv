package filters

type Interface interface {
	Apply(img []uint8)
}
