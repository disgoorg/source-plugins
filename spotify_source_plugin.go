package source_plugins

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/disgoorg/disgolink/lavalink"
)

const SearchTypeSpotify lavalink.SearchType = "spsearch"

var (
	_ lavalink.SourcePlugin = (*SpotifyPlugin)(nil)
	_ lavalink.AudioTrack   = (*SpotifyAudioTrack)(nil)
)

func NewSpotifyPlugin() *SpotifyPlugin {
	return &SpotifyPlugin{}
}

type SpotifyPlugin struct{}

func (p *SpotifyPlugin) SourceName() string {
	return "spotify"
}

func (p *SpotifyPlugin) Encode(track lavalink.AudioTrack, w io.Writer) error {
	spotifyTrack, ok := track.(*SpotifyAudioTrack)
	if !ok {
		return errors.New("track is not a *SpotifyAudioTrack")
	}
	return EncodeISRCArtworkURL(spotifyTrack.ISRC, spotifyTrack.ArtworkURL, w)
}

func (p *SpotifyPlugin) Decode(info lavalink.AudioTrackInfo, r io.Reader) (lavalink.AudioTrack, error) {
	isrc, artworkURL, err := DecodeISRCArtworkURL(r)
	if err != nil {
		return nil, err
	}
	return &SpotifyAudioTrack{
		AudioTrack: &lavalink.BasicAudioTrack{
			AudioTrackInfo: info,
		},
		ISRC:       isrc,
		ArtworkURL: artworkURL,
	}, nil
}

type SpotifyAudioTrack struct {
	lavalink.AudioTrack
	ISRC       *string `json:"isrc"`
	ArtworkURL *string `json:"artwork_url"`
}

func (t *SpotifyAudioTrack) UnmarshalJSON(data []byte) error {
	var v struct {
		*lavalink.BasicAudioTrack
		ISRC       *string `json:"isrc"`
		ArtworkURL *string `json:"artwork_url"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	t.AudioTrack = v.BasicAudioTrack
	t.ISRC = v.ISRC
	t.ArtworkURL = v.ArtworkURL
	return nil
}

func (t *SpotifyAudioTrack) Clone() lavalink.AudioTrack {
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
	return &SpotifyAudioTrack{
		AudioTrack: t.AudioTrack.Clone(),
		ISRC:       isrc,
		ArtworkURL: artworkURL,
	}
}
