[![Go Reference](https://pkg.go.dev/badge/github.com/disgoorg/source-extensions-plugin.svg)](https://pkg.go.dev/github.com/disgoorg/source-extensions-plugin)
[![Go Report](https://goreportcard.com/badge/github.com/disgoorg/source-extensions-plugin)](https://goreportcard.com/report/github.com/disgoorg/source-extensions-plugin)
[![Go Version](https://img.shields.io/github/go-mod/go-version/disgoorg/source-extensions-plugin)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/disgoorg/source-extensions-plugin/blob/master/LICENSE)
[![Disgo Version](https://img.shields.io/github/v/tag/disgoorg/source-extensions-plugin?label=release)](https://github.com/disgoorg/source-extensions-plugin/releases/latest)
[![Disgo Discord](https://discord.com/api/guilds/817327181659111454/widget.png)](https://discord.gg/NFmvZYmZMF)

# source-plugins 

source-plugins is a collection of additional audio sources for [disgolink](https://github.com/disgoorg/disgolink)

## Getting Started

### Installing

```sh
go get github.com/disgoorg/source-plugins
```

## Usage

```go
import (
    "github.com/disgoorg/disgolink/lavalink"
    "github.com/disgoorg/source-plugins"
)

// create new lavalink and add the spotify plugin
link := lavalink.New(
    lavalink.WithUserID("user_id_here"),
    lavalink.WithPlugins(
        source_plugins.NewNewSpotifyPlugin(),
        source_plugins.NewAppleMusicPlugin(),
    ),
)

// when loading track you can type cast the track to an ISRCAudioTrack to access extra data
_ = link.BestRestClient().LoadItemHandler("https://open.spotify.com/track/3yk51U329nwdpeIHV0O5ez", lavalink.NewResultHandler(
    func (track lavalink.AudioTrack) {
        if spotifyTrack, ok := track.(*source_plugins.SpotifyAudioTrack); ok {
            println("Spotify ISRC: ", spotifyTrack.ISRC)
            println("Spotify ArtworkURL: ", spotifyTrack.ArtworkURL)
        } else if appleMusicTrack, ok := track.(*source_plugins.AppleMusicAudioTrack); ok {
            println("AppleMusic ISRC: ", appleMusicTrack.ISRC)
            println("AppleMusic ArtworkURL: ", appleMusicTrack.ArtworkURL)
        }
    },
    func (playlist lavalink.AudioPlaylist) {},
    func (tracks []lavalink.AudioTrack) {},
    func () {},
    func (ex lavalink.FriendlyException) {},
))
```

## Troubleshooting

For help feel free to open an issues or reach out on [Discord](https://discord.gg/NFmvZYmZMF)

## Contributing

Contributions are welcomed but for bigger changes please first reach out via [Discord](https://discord.gg/NFmvZYmZMF) or create an issue to discuss your intentions and ideas.

## License

Distributed under the [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/disgoorg/source-extensions-plugin/blob/master/LICENSE). See LICENSE for more information.
