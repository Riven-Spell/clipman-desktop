package config

import (
	"net"
	"sync"
)

var WaitGroup sync.WaitGroup
var Connection net.Conn
var ClipboardContent string
var LockContent bool
var CLISuccess chan bool
var UserPerms bool
var AdminPerms bool

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
