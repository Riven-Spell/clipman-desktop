package config

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

func SaveCFG() {
	config := cfg{
		AdminHash: AdminHash,
		UserHash:  UserHash,
		Bind:      Bind,
		ServerIP:  ServerIP,
		Buffer:    Buffer,
	}

	data, _ := json.Marshal(config)
	if f, err := os.OpenFile(CfgLocation, os.O_CREATE|os.O_RDWR, 0666); err == nil {
		if _, err := f.Write(data); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}

func LoadCFG() bool {
	if fi, err := os.Stat(CfgLocation); err == nil {
		buffer := make([]byte, fi.Size())
		if f, err := os.Open(CfgLocation); err == nil {
			if _, err := f.Read(buffer); err == nil {
				var config cfg
				if err := json.Unmarshal(buffer, &config); err == nil {
					AdminHash = config.AdminHash
					UserHash = config.UserHash
					Buffer = config.Buffer
					Bind = config.Bind
					ServerIP = config.ServerIP
					return true
				}
			}
		}
	}

	fmt.Println("Config doesn't seem to exist, generating a config.")

	ServerIP = "127.0.0.1"
	Bind = ":7606"

	hash := sha256.New()
	hash.Write([]byte("password"))
	AdminHash = hex.EncodeToString(hash.Sum(nil))
	UserHash = hex.EncodeToString(hash.Sum(nil))

	SaveCFG()
	return false
}
