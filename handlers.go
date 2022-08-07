package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_ = r
	if _, err := w.Write([]byte("<a href='/songs'>go to /songs</a>")); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func songsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(getLibrary()); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if r.Method == http.MethodPost {
		var newSong Song

		err := json.NewDecoder(r.Body).Decode(&newSong)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		addSong(newSong) //need to implement error check here
		w.WriteHeader(http.StatusCreated)

		return
	}

	http.Error(w, "wrong method jack ash", http.StatusMethodNotAllowed)
}

func singleSongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idStr := strings.TrimPrefix(r.URL.Path, "/songs/")
		idInt, err := strconv.Atoi(idStr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			if _, err := w.Write([]byte("not a number bub")); err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		song := getSong(idInt)

		if song.ID == -1 {
			w.WriteHeader(http.StatusBadRequest)
		}

		if err := json.NewEncoder(w).Encode(getSong(idInt)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}
	http.Error(w, "wrong method jack ash", http.StatusMethodNotAllowed)
}
