package model

import "time"

type ScribeHeader struct {
	Time    time.Time
	User    string
	Chatter string
}

type ScribeEntry struct {
	Time time.Time
	Who  string
	What string
}

type ScribeFooter struct {
	Time time.Time
}
