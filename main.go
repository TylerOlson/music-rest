package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type song struct {
	Id     string
	Name   string
	Artist string
	Length string
}

func innitLibrary() {
	songs := []song{{Id: "0", Name: "Mo Bamba", Artist: "Sheck Wes", Length: "185"}}

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

func getSong(id int) song {
	songs := loadLibrary()

	for i, s := range songs {
		if i == id {
			return s
		}
	}

	return song{Id: "-1"}
}

func main() {
	innitLibrary()

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/songs", songsHandler)
	mux.HandleFunc("/songs/", singleSongHandler)
	fmt.Println("Starting server")

	log.Fatal(http.ListenAndServe(":80", mux))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_ = r
	if _, err := w.Write([]byte("<a href='/songs'>go to /songs</a>")); err != nil {
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

func singleSongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/songs/")
		if idInt, err := strconv.Atoi(id); err == nil {
			w.Header().Set("Content-Type", "application/json")
			song := getSong(idInt)
			if song.Id == "-1" {
				w.WriteHeader(http.StatusBadRequest)
			}

			if err := json.NewEncoder(w).Encode(getSong(idInt)); err != nil {
				log.Fatal(err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("not a number bub")); err != nil {
				log.Fatal(err)
			}
		}

		return
	}
	http.Error(w, "wrong method jack ash", 500)
}
