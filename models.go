package main

import (
	"time"
)

type Strategy struct {
	Name    string    `bson:"name"`
	Script  string    `bson:"script"`
	Created time.Time `bson:"created"`
}
