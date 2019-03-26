package config

import (
	"net"
	"sync"
)

var WaitGroup sync.WaitGroup
var Connection net.Conn

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
