package config

import (
	"crypto/tls"
	"net"
	"sync"
)

var WaitGroup sync.WaitGroup
var Connection net.Conn
var ClipboardContent string
var LockContent bool
var NoSync = false
var CLISuccess chan bool
var UserPerms bool
var AdminPerms bool
var TLSConfig = &tls.Config{
	InsecureSkipVerify: false,
}

//Config variables below
var (
	AdminHash     string
	UserHash      string
	ServerIP      string
	ServerUsesTLS bool
	Buffer        int
)

type cfg struct {
	AdminHash     string
	UserHash      string
	ServerIP      string
	ServerUsesTLS bool
	TLSInsecure   bool
	Buffer        int
}
