package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type song struct {
	Name   string
	Artist string
	Length int
}

func innitLibrary() {
	songs := []song{{Name: "Mo Bamba", Artist: "Sheck Wes", Length: 185}}

	song, err := json.Marshal(songs)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("library.json", song, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func loadLibrary() []song {
	songsData, err := ioutil.ReadFile("library.json")
	if err != nil {
		log.Fatal(err)
	}
	var songs []song
	err = json.Unmarshal(songsData, &songs)
	if err != nil {
		log.Fatal(err)
	}

	return songs
}

func main() {
	innitLibrary()

	loadLibrary()

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/songs", songsHandler)
	fmt.Println("Starting server")

	log.Fatal(http.ListenAndServe(":80", mux))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("go to /songs")); err != nil {
		log.Fatal(err)
	}
}

func songsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(loadLibrary()); err != nil {
			log.Fatal(err)
		}

		return
	}
	http.Error(w, "wrong method jack ash", 500)
}

func addSongsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := json.NewEncoder(w).Encode(loadLibrary()); err != nil {
			log.Fatal(err)
		}

		return
	}
	http.Error(w, "wrong method jack ash", 500)
}
