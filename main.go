package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"playground/controllers"
	"playground/models"

	"github.com/urfave/cli/v2"
)

const (
	flagArtist    = "artist"
	flagSongTitle = "song"
)

func main() {
	app := &cli.App{
		Authors: []*cli.Author{
			{
				Name: "Joaquin Diaz",
			},
		},
		Version: "0.0.1",
		Name:    "playground",
		Usage:   "Get Lyrics for a given Artist and Song",
		Action:  getLyrics,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     flagArtist,
				Required: true,
			},
			&cli.StringFlag{
				Name:     flagSongTitle,
				Required: true,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// getLyrics : Get song lyrics from a public api
func getLyrics(ctx *cli.Context) error {
	artist := ctx.String(flagArtist)
	title := ctx.String(flagSongTitle)

	apiURL := "https://api.lyrics.ovh/v1/%s/%s"
	url := fmt.Sprintf(apiURL, artist, title)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Lyrics not found.")
	} else {
		var song models.SongLyrics
		json.NewDecoder(resp.Body).Decode(&song)

		if song.Lyrics == "" {
			fmt.Println("Lyrics not found.")
		} else {
			fmt.Println("Artist:", artist)
			fmt.Println()
			fmt.Println("Title:", title)
			fmt.Println()
			fmt.Println("Lyrics:")
			fmt.Println()
			fmt.Println(song.Lyrics)
			fmt.Println()
		}
	}

	return nil
}

/*
Previous playground methods
*/

// helloGo : Print 'Hello Go!' to the console
func helloGo() {
	fmt.Println("Hello Go!")
}

// printUser : Print a harcoded User struct to the console
func printUser() {
	user := models.User{
		ID:        1,
		FirstName: "Mr",
		LastName:  "Spock",
	}

	fmt.Println(user)
}

// runUsersWebService : Run a simplified web server to fetch harcoded users
func runUsersWebService() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
