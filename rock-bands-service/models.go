package main

type Artist struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Born  string `json:"born"`
	Genre string `json:"genre"`
}

type Album struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Year     string `json:"year"`
	ArtistID string `json:"artist_id"`
}
