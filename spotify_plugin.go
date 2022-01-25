package spotify

import (
	"errors"
	"io"

	"github.com/DisgoOrg/disgolink/lavalink"
)

const SearchTypeSpotify lavalink.SearchType = "spsearch"

var (
	_ lavalink.SourceExtension = (*Plugin)(nil)
)

func New() *Plugin {
	return &Plugin{}
}

type Plugin struct{}

func (p *Plugin) SourceName() string {
	return "spotify"
}

func (p *Plugin) Encode(track lavalink.AudioTrack, w io.Writer) (err error) {
	spotifyTrack, ok := track.(*AudioTrack)
	if !ok {
		return errors.New("track is not a SpotifyAudioTrack")
	}

	if err = lavalink.WriteNullableString(w, spotifyTrack.AudioTrackInfo.ISRC); err != nil {
		return
	}
	return lavalink.WriteNullableString(w, spotifyTrack.AudioTrackInfo.ArtworkURL)
}

func (p *Plugin) Decode(track string, info lavalink.AudioTrackInfo, r io.Reader) (spotifyTrack lavalink.AudioTrack, err error) {
	var isrc, artworkURL *string

	if isrc, err = lavalink.ReadNullableString(r); err != nil {
		return
	}
	if artworkURL, err = lavalink.ReadNullableString(r); err != nil {
		return
	}

	return &AudioTrack{
		AudioTrack: track,
		AudioTrackInfo: &AudioTrackInfo{
			AudioTrackInfo: info,
			ISRC:           isrc,
			ArtworkURL:     artworkURL,
		},
	}, nil
}

var (
	_ lavalink.AudioTrack     = (*AudioTrack)(nil)
	_ lavalink.AudioTrackInfo = (*AudioTrackInfo)(nil)
)

type AudioTrack struct {
	AudioTrack     string          `json:"track"`
	AudioTrackInfo *AudioTrackInfo `json:"info"`
}

func (t *AudioTrack) Track() string {
	return t.AudioTrack
}

func (t *AudioTrack) Info() lavalink.AudioTrackInfo {
	return t.AudioTrackInfo
}

type AudioTrackInfo struct {
	lavalink.AudioTrackInfo
	ISRC       *string `json:"isrc"`
	ArtworkURL *string `json:"artwork_url"`
}
