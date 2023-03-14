package main

import (
	"time"
)

type Strategy struct {
	Name    string    `json:"name"`
	Script  string    `json:"script"`
	Created time.Time `json:"created"`
}
