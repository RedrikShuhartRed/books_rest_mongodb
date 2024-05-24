package models

type Movie struct {
	Title    string   `json:"title"`
	Director string   `json:"director"`
	Year     int      `json:"year"`
	Genres   []string `json:"genres"`
	Rating   float64  `json:"rating"`
	Duration struct {
		Hours   int `json:"hours"`
		Minutes int `json:"minutes"`
	} `json:"duration"`
	Reviews []Review `json:"reviews"`
}

type Review struct {
	Name string `json:"name"`
	Text string `json:"text"`
}
