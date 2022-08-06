package main

type Song struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Length string `json:"length"`
}

func NewSong(name string, artist string, length string) Song {
	return Song{Name: name, Artist: artist, Length: length}
}
