package music

import "github.com/DisgoOrg/disgolink/lavalink"

// func (q *Queue) Add(track lavalink.Track) (queue Queue, err error) {

// 	tracks := append(q.Tracks, track)

// 	return Queue{tracks, q.GuildId}, nil
// }

// func (q *Queue) RemoveFirst() (queue Queue, err error) {

// 	return
// }

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
