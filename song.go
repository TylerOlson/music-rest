package main

type Song struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Length string `json:"length"`
}

func newSong(name string, artist string, length string) Song {
	return Song{Name: name, Artist: artist, Length: length}
}
