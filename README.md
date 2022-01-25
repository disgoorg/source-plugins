[![Go Reference](https://pkg.go.dev/badge/github.com/DisgoOrg/spotify-plugin.svg)](https://pkg.go.dev/github.com/DisgoOrg/spotify-plugin)
[![Go Report](https://goreportcard.com/badge/github.com/DisgoOrg/spotify-plugin)](https://goreportcard.com/report/github.com/DisgoOrg/spotify-plugin)
[![Go Version](https://img.shields.io/github/go-mod/go-version/DisgoOrg/spotify-plugin)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/DisgoOrg/spotify-plugin/blob/master/LICENSE)
[![Disgo Version](https://img.shields.io/github/v/tag/DisgoOrg/spotify-plugin?label=release)](https://github.com/DisgoOrg/spotify-plugin/releases/latest)
[![Disgo Discord](https://discord.com/api/guilds/817327181659111454/widget.png)](https://discord.gg/NFmvZYmZMF)

# spotify-plugin

spotify-plugin is a [Lavalink](https://github.com/freyacodes/Lavalink) Client which supports the latest Lavalink 3.4 release

## Getting Started

### Installing

```sh
go get github.com/DisgoOrg/spotify-plugin
```

## Usage

```go
import (
    "github.com/DisgoOrg/disgolink/lavalink"
    "github.com/DisgoOrg/spotify-plugin"
)
// create new lavalink and add the spotify plugin
link := lavalink.New(
    lavalink.WithUserID("user_id_here"),
    lavalink.WithPlugins(spotify.New()),
)

// when loading track you can type cast the track to a spotify track to access extra data
_ = link.BestRestClient().LoadItemHandler("https://open.spotify.com/track/3yk51U329nwdpeIHV0O5ez", lavalink.NewResultHandler(
    func (track lavalink.AudioTrack) {
        if spotifyTrack, ok := track.(*spotify.AudioTrack); ok {
            println("ISRC: ", spotifyTrack.AudioTrackInfo.ISRC)
            println("ArtworkURL: ", spotifyTrack.AudioTrackInfo.ArtworkURL)
        }
    },
    func (playlist lavalink.AudioPlaylist) {},
    func (tracks []lavalink.AudioTrack) {},
    func () {},
    func (ex lavalink.FriendlyException) {},
))
```

## Example

You can find an example under [_example](https://github.com/DisgoOrg/spotify-plugin/tree/master/_example)

## Troubleshooting

For help feel free to open an issues or reach out on [Discord](https://discord.gg/NFmvZYmZMF)

## Contributing

Contributions are welcomed but for bigger changes please first reach out via [Discord](https://discord.gg/NFmvZYmZMF) or create an issue to discuss your intentions and ideas.

## License

Distributed under the [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/DisgoOrg/spotify-plugin/blob/master/LICENSE). See LICENSE for more information.
