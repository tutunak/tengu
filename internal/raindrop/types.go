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

func (c Collections) Deadline() (time.Time, bool) {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Error(args ...any) {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Errorf(format string, args ...any) {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Fail() {
	//TODO implement me
	panic("implement me")
}

func (c Collections) FailNow() {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Failed() bool {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Fatal(args ...any) {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Fatalf(format string, args ...any) {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Helper() {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Log(args ...any) {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Logf(format string, args ...any) {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Name() string {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Parallel() {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Skip(args ...any) {
	//TODO implement me
	panic("implement me")
}

func (c Collections) SkipNow() {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Skipf(format string, args ...any) {
	//TODO implement me
	panic("implement me")
}

func (c Collections) Skipped() bool {
	//TODO implement me
	panic("implement me")
}
