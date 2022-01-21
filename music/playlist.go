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
	GetPlaylist(guildID string) Playlist
	GetPlaylists() map[string]Playlist

	AddToPlaylist(track lavalink.Track, guildID string) Playlist
	RemoveLastTrack(guildID string) Playlist
}

type playlistManagertmpl struct {
	playlists map[string]Playlist
}

var _ PlaylistManager = (*playlistManagertmpl)(nil)

func (p *playlistManagertmpl) GetPlaylists() map[string]Playlist {
	return p.playlists
}

func (p *playlistManagertmpl) GetPlaylist(guildID string) Playlist {
	if playlist, ok := p.playlists[guildID]; ok {
		return playlist
	}
	playlist := Playlist{GuildID: guildID}
	p.playlists[guildID] = playlist
	return playlist
}

func (p *playlistManagertmpl) AddToPlaylist(track lavalink.Track, guildID string) Playlist {

	playlist := p.GetPlaylist(guildID)

	playlist = Playlist{playlist.GuildID, append(playlist.Tracks, track)}

	p.playlists[playlist.GuildID] = playlist

	return playlist
}

func (p *playlistManagertmpl) RemoveLastTrack(guildID string) Playlist {
	playlist := p.GetPlaylist(guildID)

	playlist = Playlist{GuildID: guildID, Tracks: removeIndex(playlist.Tracks, 1)}

	p.playlists[playlist.GuildID] = playlist

	return playlist

}

func removeIndex(s []lavalink.Track, index int) []lavalink.Track {
	ret := make([]lavalink.Track, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
