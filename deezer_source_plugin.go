package source_plugins

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/disgoorg/disgolink/lavalink"
)

const (
	SearchTypeDeezer     lavalink.SearchType = "dzsearch"
	SearchTypeDeezerISRC lavalink.SearchType = "dzisrc"
)

var (
	_ lavalink.SourcePlugin = (*DeezerPlugin)(nil)
	_ lavalink.AudioTrack   = (*DeezerAudioTrack)(nil)
)

func NewDeezerPlugin() *DeezerPlugin {
	return &DeezerPlugin{}
}

type DeezerPlugin struct{}

func (p *DeezerPlugin) SourceName() string {
	return "deezer"
}

func (p *DeezerPlugin) Encode(track lavalink.AudioTrack, w io.Writer) error {
	deezerTrack, ok := track.(*DeezerAudioTrack)
	if !ok {
		return errors.New("track is not a *DeezerAudioTrack")
	}
	return EncodeISRCArtworkURL(deezerTrack.ISRC, deezerTrack.ArtworkURL, w)
}

func (p *DeezerPlugin) Decode(info lavalink.AudioTrackInfo, r io.Reader) (lavalink.AudioTrack, error) {
	isrc, artworkURL, err := DecodeISRCArtworkURL(r)
	if err != nil {
		return nil, err
	}
	return &DeezerAudioTrack{
		AudioTrack: &lavalink.BasicAudioTrack{
			AudioTrackInfo: info,
		},
		ISRC:       isrc,
		ArtworkURL: artworkURL,
	}, nil
}

type DeezerAudioTrack struct {
	lavalink.AudioTrack
	ISRC       *string `json:"isrc"`
	ArtworkURL *string `json:"artwork_url"`
}

func (t *DeezerAudioTrack) UnmarshalJSON(data []byte) error {
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

func (t *DeezerAudioTrack) Clone() lavalink.AudioTrack {
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
	return &DeezerAudioTrack{
		AudioTrack: t.AudioTrack.Clone(),
		ISRC:       isrc,
		ArtworkURL: artworkURL,
	}
}
