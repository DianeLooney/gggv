package playlist

type Playlist struct {
	Videos []PlaylistVideo
}

type PlaylistVideo struct {
	Path     string
	Duration float64
}
