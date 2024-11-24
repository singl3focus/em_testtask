package models

type SongInfo struct {
	Group  string `json:"group"`
	Title  string `json:"title"`
}

type Song struct {
	SongInfo
	Verses []Verse `json:"verses"`
}

type Verse struct {
	Number int `json:"number"`
	Text   string `json:"text"`
}
