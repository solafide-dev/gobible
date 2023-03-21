package main

import (
	_ "embed"
)

type Bible struct {
	Version Version `json:"version"`
	Books   []Book  `json:"books"`
}

type Version struct {
	Name   string `json:"name"`
	Abbrev string `json:"abbrev"`
}

type Book struct {
	Name     string    `json:"name"`
	ID       string    `json:"id"`
	Chapters []Chapter `json:"chapters"`
}

type Chapter struct {
	Number int     `json:"number"`
	Name   string  `json:"name"`
	ID     string  `json:"id"`
	Verses []Verse `json:"verses"`
}

type Verse struct {
	Verse int    `json:"verse"`
	Text  string `json:"text"`
	OSIS  string `json:"osis"`
}
