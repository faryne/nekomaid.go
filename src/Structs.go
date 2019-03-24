package nekomaid

// 定義 json 結構
type Header struct {
	num         int
	start       int
	perpage     int
}
type Photo struct {
	Image           string
	Width           int
	Height          int
}
type Artwork struct {
	Id              string
	Author_id       string
	Site            string
	title           string
	Thumbnail       string
	R_18            bool
	Original_url    string
	Page_url        string
	Photos          []Photo
	Tags            []string
}
type SearchResult struct {
	Header      Header
	Artworks    []Artwork
}
