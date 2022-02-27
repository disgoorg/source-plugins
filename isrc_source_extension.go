package source_extensions

import (
	"io"

	"github.com/DisgoOrg/disgolink/lavalink"
	"github.com/pkg/errors"
)

type ISRCSourceExtension struct{}

func (p ISRCSourceExtension) Encode(track lavalink.AudioTrack, w io.Writer) (err error) {
	isrcAudioTrack, ok := track.(*ISRCAudioTrack)
	if !ok {
		return errors.New("track is not a SpotifyAudioTrack")
	}

	if err = lavalink.WriteNullableString(w, isrcAudioTrack.ISRC); err != nil {
		return
	}
	return lavalink.WriteNullableString(w, isrcAudioTrack.ArtworkURL)
}

func (p ISRCSourceExtension) Decode(info lavalink.AudioTrackInfo, r io.Reader) (isrcTrack lavalink.AudioTrack, err error) {
	var isrc, artworkURL *string

	if isrc, err = lavalink.ReadNullableString(r); err != nil {
		return
	}
	if artworkURL, err = lavalink.ReadNullableString(r); err != nil {
		return
	}

	return &ISRCAudioTrack{
		AudioTrack: &lavalink.BasicAudioTrack{
			AudioTrackInfo: info,
		},
		ISRC:       isrc,
		ArtworkURL: artworkURL,
	}, nil
}

var (
	_ lavalink.AudioTrack = (*ISRCAudioTrack)(nil)
)

type ISRCAudioTrack struct {
	lavalink.AudioTrack
	ISRC       *string `json:"isrc"`
	ArtworkURL *string `json:"artwork_url"`
}

func (t *ISRCAudioTrack) Clone() lavalink.AudioTrack {
	info := t.Info()
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
	return &ISRCAudioTrack{
		AudioTrack: t.AudioTrack.Clone(),
		ISRC:       isrc,
		ArtworkURL: artworkURL,
	}
}
