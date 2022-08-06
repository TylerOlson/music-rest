package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func clearLibrary() {
	if err := ioutil.WriteFile("library.json", []byte("[]"), 0644); err != nil {
		log.Fatal(err)
	}
}

func getLibrary() []Song {
	encodedSongs, err := ioutil.ReadFile("library.json")
	if err != nil {
		log.Fatal(err)
	}

	var songs []Song
	err = json.Unmarshal(encodedSongs, &songs)
	if err != nil {
		log.Fatal(err)
	}

	return songs
}

func addSong(newSong Song) {
	songs := getLibrary()

	newSong.ID = len(songs)

	songs = append(songs, newSong)

	song, err := json.Marshal(songs)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("library.json", song, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getSong(id int) Song {
	songs := getLibrary()

	for i, s := range songs {
		if i == id {
			return s
		}
	}

	return Song{ID: -1}
}
