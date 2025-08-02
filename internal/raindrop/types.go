package raindrop

import "time"

type Access struct {
	Level     int  `json:"level"`
	Draggable bool `json:"draggable"`
}
type Coloborators struct {
	Id string `json:"$id"`
}

type User struct {
	Id int `json:"$id"`
}

type Collection struct {
	Id            int          `json:"_id"`
	Access        Access       `json:"access"`
	Collaborators Coloborators `json:"collaborators"`
	Color         string       `json:"color"`
	Count         int          `json:"count"`
	Cover         []string     `json:"cover"`
	Created       time.Time    `json:"created"`
	Expanded      bool         `json:"expanded"`
	LastUpdate    time.Time    `json:"lastUpdate"`
	Public        bool         `json:"public"`
	Sort          int          `json:"sort"`
	Title         string       `json:"title"`
	User          User         `json:"user"`
	View          string       `json:"view"`
}

type Collections struct {
	Result bool         `json:"result"`
	Items  []Collection `json:"items"`
}
