package source_extensions

import "github.com/DisgoOrg/disgolink/lavalink"

const (
	SearchTypeSpotify    lavalink.SearchType = "spsearch"
	SearchTypeAppleMusic lavalink.SearchType = "amsearch"
)

var (
	_ lavalink.SourceExtension = (*SpotifyPlugin)(nil)
	_ lavalink.SourceExtension = (*AppleMusicPlugin)(nil)
)

func NewSpotifyPlugin() *SpotifyPlugin {
	return &SpotifyPlugin{
		ISRCSourceExtension: ISRCSourceExtension{},
	}
}

type SpotifyPlugin struct {
	ISRCSourceExtension
}

func (p *SpotifyPlugin) SourceName() string {
	return "spotify"
}

func NewAppleMusicPlugin() *AppleMusicPlugin {
	return &AppleMusicPlugin{
		ISRCSourceExtension: ISRCSourceExtension{},
	}
}

type AppleMusicPlugin struct {
	ISRCSourceExtension
}

func (p *AppleMusicPlugin) SourceName() string {
	return "applemusic"
}
