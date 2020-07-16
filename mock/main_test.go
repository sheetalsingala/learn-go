package main

import (
	"testing"

	"github.com/zmb3/spotify"
)

type mockSpotifyClient struct{}

func (m *mockSpotifyClient) GetPlaylist(playlistID spotify.ID) (*spotify.FullPlaylist, error) {
	// 3
	return &spotify.FullPlaylist{
		SimplePlaylist: spotify.SimplePlaylist{
			Name: "whatever",
		},
	}, nil
}

func Test_NewGetPlaylistName(t *testing.T) {
	// 4
	client := &mockSpotifyClient{}

	// 5
	name := getPlaylistName(client, "whatever")

	if name != "whatever" {
		t.Errorf("expected %s, got %s", "whatever", name)
	}
}
