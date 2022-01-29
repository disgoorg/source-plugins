package spotify

import (
	"errors"
	"io"
	"time"

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

	if err = lavalink.WriteNullableString(w, spotifyTrack.ISRC); err != nil {
		return
	}
	return lavalink.WriteNullableString(w, spotifyTrack.ArtworkURL)
}

func (p *Plugin) Decode(info lavalink.AudioTrackInfo, r io.Reader) (spotifyTrack lavalink.AudioTrack, err error) {
	var isrc, artworkURL *string

	if isrc, err = lavalink.ReadNullableString(r); err != nil {
		return
	}
	if artworkURL, err = lavalink.ReadNullableString(r); err != nil {
		return
	}

	return &AudioTrack{
		AudioTrackInfo: info,
		ISRC:           isrc,
		ArtworkURL:     artworkURL,
	}, nil
}

var (
	_ lavalink.AudioTrack = (*AudioTrack)(nil)
)

type AudioTrack struct {
	AudioTrackInfo lavalink.AudioTrackInfo `json:"info"`
	ISRC           *string                 `json:"isrc"`
	ArtworkURL     *string                 `json:"artwork_url"`
}

func (t *AudioTrack) Info() lavalink.AudioTrackInfo {
	return t.AudioTrackInfo
}

func (t *AudioTrack) SetPosition(position time.Duration) {
	t.AudioTrackInfo.Position = position
}

func (t *AudioTrack) Clone() lavalink.AudioTrack {
	info := t.AudioTrackInfo
	info.Position = 0
	var (
		isrc, artworkURL *string
	)
	if t.ISRC != nil {
		isrc = new(string)
		*isrc = *t.ISRC
	}
	if t.ArtworkURL != nil {
		artworkURL = new(string)
		*artworkURL = *t.ArtworkURL
	}
	return &AudioTrack{
		AudioTrackInfo: info,
		ISRC:           isrc,
		ArtworkURL:     artworkURL,
	}
}
