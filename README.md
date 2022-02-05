[![Go Reference](https://pkg.go.dev/badge/github.com/DisgoOrg/source-extensions-plugin.svg)](https://pkg.go.dev/github.com/DisgoOrg/source-extensions-plugin)
[![Go Report](https://goreportcard.com/badge/github.com/DisgoOrg/source-extensions-plugin)](https://goreportcard.com/report/github.com/DisgoOrg/source-extensions-plugin)
[![Go Version](https://img.shields.io/github/go-mod/go-version/DisgoOrg/source-extensions-plugin)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/DisgoOrg/source-extensions-plugin/blob/master/LICENSE)
[![Disgo Version](https://img.shields.io/github/v/tag/DisgoOrg/source-extensions-plugin?label=release)](https://github.com/DisgoOrg/source-extensions-plugin/releases/latest)
[![Disgo Discord](https://discord.com/api/guilds/817327181659111454/widget.png)](https://discord.gg/NFmvZYmZMF)

# source-extensions-plugin

source-extensions-plugin is a collection of additional source extension for [disgolink](https://github.com/DisgoOrg/disgolink)

## Getting Started

### Installing

```sh
go get github.com/DisgoOrg/source-extensions-plugin
```

## Usage

```go
import (
    "github.com/DisgoOrg/disgolink/lavalink"
    "github.com/DisgoOrg/source-extensions-plugin"
)

// create new lavalink and add the spotify plugin
link := lavalink.New(
    lavalink.WithUserID("user_id_here"),
    lavalink.WithPlugins(
        source_extensions.NewNewSpotifyPlugin(),
        source_extensions.NewAppleMusicPlugin(),
    ),
)

// when loading track you can type cast the track to an ISRCAudioTrack to access extra data
_ = link.BestRestClient().LoadItemHandler("https://open.spotify.com/track/3yk51U329nwdpeIHV0O5ez", lavalink.NewResultHandler(
    func (track lavalink.AudioTrack) {
        if isrcTrack, ok := track.(*source_extensions.ISRCAudioTrack); ok {
            println("ISRC: ", isrcTrack.ISRC)
            println("ArtworkURL: ", isrcTrack.ArtworkURL)
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

Distributed under the [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/DisgoOrg/source-extensions-plugin/blob/master/LICENSE). See LICENSE for more information.
