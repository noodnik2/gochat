package model

import "time"

type Context struct {
	Time         time.Time
	Participants []string
}

type Entry struct {
	Time time.Time
	Who  string
	What string
}

type Outcome struct {
	Time time.Time
}
