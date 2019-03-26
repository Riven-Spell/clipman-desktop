package config

import "sync"

var WaitGroup sync.WaitGroup

//Config variables below
var (
	AdminHash string
	UserHash  string
	ServerIP  string
	Buffer    int
)

type cfg struct {
	AdminHash string
	UserHash  string
	ServerIP  string
	Buffer    int
}
