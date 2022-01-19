package music

import (
	"github.com/DisgoOrg/disgolink/lavalink"
)

func New() PlaylistManager {
	playlistManager := &playlistManagertmpl{
		playlists: map[string]Playlist{},
	}
	return playlistManager
}

type Playlist struct {
	GuildID string
	Tracks  []lavalink.Track
}

type PlaylistManager interface {
	Playlist(guildID string) Playlist
	Playlists() map[string]Playlist

	AddToPlaylist(track lavalink.Track, guildID string) Playlist
	GetPlaylist(guildID string) Playlist
}

type playlistManagertmpl struct {
	playlists map[string]Playlist
}

var _ PlaylistManager = (*playlistManagertmpl)(nil)

func (p *playlistManagertmpl) Playlists() map[string]Playlist {
	return p.playlists
}

func (p *playlistManagertmpl) Playlist(guildID string) Playlist {
	if playlist, ok := p.playlists[guildID]; ok {
		return playlist
	}
	playlist := Playlist{GuildID: guildID}
	p.playlists[guildID] = playlist
	return playlist
}

func (p *playlistManagertmpl) AddToPlaylist(track lavalink.Track, guildID string) Playlist {

	playlist := p.Playlist(guildID)

	playlist = Playlist{playlist.GuildID, append(playlist.Tracks, track)}

	p.playlists[playlist.GuildID] = playlist

	return playlist
}

func (p *playlistManagertmpl) GetPlaylist(guildID string) Playlist {

	playlist := p.Playlist(guildID)

	return playlist
}
