//Returns the playlist name for a given playlist ID

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {

	const PLAYLIST spotify.ID = "4OyKDT6cLw96G7bd8nTfxD"
	client := newClient()
	playlistName := getPlaylistName(client, PLAYLIST)
	fmt.Println(playlistName)
}

type spotifyClient interface {
	GetPlaylist(playlistID spotify.ID) (*spotify.FullPlaylist, error)
}

func newClient() *spotify.Client { //Return a new spotify client
	conf := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"), //Set SPOTIFY_ID andSPOTIFY_SECRET
		TokenURL:     spotify.TokenURL,
	}

	token, err := conf.Token(context.Background())
	if err != nil {
		log.Fatal("could not get token")
	}

	client := spotify.Authenticator{}.NewClient(token)

	return &client
}

func getPlaylistName(client spotifyClient, PLAYLIST spotify.ID) string {
	results, err := client.GetPlaylist(PLAYLIST)
	if err != nil {
		log.Fatal("could not get playlist")
	}

	return (results.Name)
}
