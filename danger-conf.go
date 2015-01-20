package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const bytesPerMB int64 = 1000000

type DangerConfig struct {
	HttpPort int    `json:"http-port"`
	TcpPort  int    `json:"tcp-port"`
	DBFile   string `json:"db-file"`
	PidFile  string `json:"pid-file"`
	LogFile  string `json:"log-file"`
	WorkDir  string `json:"work-dir"`
}

func readConfig(confFile string) (DangerConfig, error) {

	dconf := DangerConfig{}
	content, err := ioutil.ReadFile(confFile)
	if err != nil {
		return dconf, fmt.Errorf("LoadConfig(): Error reading config:", err)
	}

	err = json.Unmarshal(content, &dconf)

	if err != nil {
		return dconf, fmt.Errorf("LoadConfig(): Error unmarshaling JSON", err)
	}

	return dconf, nil
}
