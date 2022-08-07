package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	clearLibrary()

	addSong(NewSong("test", "dasd", "ad"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/songs", songsHandler)
	mux.HandleFunc("/songs/", singleSongHandler)

	fmt.Println("Starting server")
	log.Fatal(http.ListenAndServe(":80", mux))
}
