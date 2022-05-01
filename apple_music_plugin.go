package source_plugins

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/disgoorg/disgolink/lavalink"
)

const SearchTypeAppleMusic lavalink.SearchType = "amsearch"

var (
	_ lavalink.SourcePlugin = (*AppleMusicPlugin)(nil)
	_ lavalink.AudioTrack   = (*AppleMusicAudioTrack)(nil)
)

func NewAppleMusicPlugin() *AppleMusicPlugin {
	return &AppleMusicPlugin{}
}

type AppleMusicPlugin struct{}

func (p *AppleMusicPlugin) SourceName() string {
	return "applemusic"
}

func (p *AppleMusicPlugin) Encode(track lavalink.AudioTrack, w io.Writer) error {
	spotifyTrack, ok := track.(*AppleMusicAudioTrack)
	if !ok {
		return errors.New("track is not a *AppleMusicAudioTrack")
	}
	return EncodeISRCArtworkURL(spotifyTrack.ISRC, spotifyTrack.ArtworkURL, w)
}

func (p *AppleMusicPlugin) Decode(info lavalink.AudioTrackInfo, r io.Reader) (lavalink.AudioTrack, error) {
	isrc, artworkURL, err := DecodeISRCArtworkURL(r)
	if err != nil {
		return nil, err
	}
	return &AppleMusicAudioTrack{
		AudioTrack: &lavalink.BasicAudioTrack{
			AudioTrackInfo: info,
		},
		ISRC:       isrc,
		ArtworkURL: artworkURL,
	}, nil
}

type AppleMusicAudioTrack struct {
	lavalink.AudioTrack
	ISRC       *string `json:"isrc"`
	ArtworkURL *string `json:"artwork_url"`
}

func (t *AppleMusicAudioTrack) UnmarshalJSON(data []byte) error {
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

func (t *AppleMusicAudioTrack) Clone() lavalink.AudioTrack {
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
	return &AppleMusicAudioTrack{
		AudioTrack: t.AudioTrack.Clone(),
		ISRC:       isrc,
		ArtworkURL: artworkURL,
	}
}
