package source_plugins

import (
	"io"

	"github.com/disgoorg/disgolink/lavalink"
)

func EncodeISRCArtworkURL(isrc *string, artworkURL *string, w io.Writer) (err error) {
	if err = lavalink.WriteNullableString(w, isrc); err != nil {
		return
	}
	return lavalink.WriteNullableString(w, artworkURL)
}

func DecodeISRCArtworkURL(r io.Reader) (isrc *string, artworkURL *string, err error) {
	if isrc, err = lavalink.ReadNullableString(r); err != nil {
		return
	}
	if artworkURL, err = lavalink.ReadNullableString(r); err != nil {
		return
	}

	return
}
