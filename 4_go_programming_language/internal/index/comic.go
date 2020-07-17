package index

// Comic is the json representation of comic we get from xkcd
type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
}
