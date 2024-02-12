package model

import "time"

type Context struct {
	Time         time.Time
	User         string
	Chatter      string
}

type Entry struct {
	Time time.Time
	Who  string
	What string
}

type Outcome struct {
	Time time.Time
}
