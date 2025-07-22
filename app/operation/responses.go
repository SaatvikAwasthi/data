package operation

type RawDataResponse struct {
	Data []RawData `json:"data"`
}
type RawData struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Source    string `json:"source"`
	Posts     []Post `json:"posts"`
}

type Post struct {
	UserId uint64 `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type AddRawDataResponse struct {
	Message string `json:"message"`
}
