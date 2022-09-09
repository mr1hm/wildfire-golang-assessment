package api

type NameResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type JokeResponse struct {
	Type  string  `json:"type"`
	Value JokeMap `json:"value"`
}
type JokeMap struct {
	ID         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

type Responses struct {
	Name     NameResponse
	JokeData JokeResponse
}
