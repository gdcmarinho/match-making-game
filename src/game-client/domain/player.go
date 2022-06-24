package domain

type player struct {
    ID     		string  `json:"id"`
    Nickname  	string  `json:"nickname"`
    Tag 		string  `json:"tag"`
    Rank  		int 	`json:"rank"`
}