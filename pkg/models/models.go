package models

import (
	"errors"
	"time"
)

var ErrorNoRecord = errors.New("Record: no matching entry found ")

type Snipped struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}