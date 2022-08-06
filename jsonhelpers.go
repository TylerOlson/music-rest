package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func clearLibrary() {
	library := make(map[int]Song)

	encodedLibrary, err := json.Marshal(library)
	if err != nil {
		fmt.Println(err)
	}

	if err := ioutil.WriteFile("library.json", encodedLibrary, 0644); err != nil {
		log.Fatal(err)
	}
}

func getLibrary() map[int]Song {
	encodedLibrary, err := ioutil.ReadFile("library.json")
	if err != nil {
		log.Fatal(err)
	}

	var library = make(map[int]Song)
	err = json.Unmarshal(encodedLibrary, &library)
	if err != nil {
		log.Fatal(err)
	}

	return library
}

func addSong(newSong Song) {
	library := getLibrary()

	newSong.ID = len(library)

	library[newSong.ID] = newSong

	song, err := json.Marshal(library)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("library.json", song, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getSong(id int) Song {
	return getLibrary()[id]
}
