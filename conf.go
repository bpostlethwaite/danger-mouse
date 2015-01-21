package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func readConfigs(filename string, dirs []string) (DangerConfig, error) {
	for _, d := range dirs {
		confpath := filepath.Join(d, filename)
		if _, err := os.Stat(confpath); os.IsNotExist(err) {
			continue
		} else {
			fmt.Println("attempting to read from", confpath)
			conf, err := readConfig(confpath)
			return conf, err
		}
	}
	// didn't find a config file
	return DangerConfig{}, fmt.Errorf("readConfigs(); no danger.json configuration file found")

}
func readConfig(confFile string) (DangerConfig, error) {

	dconf := DangerConfig{}
	content, err := ioutil.ReadFile(confFile)
	if err != nil {
		return dconf, fmt.Errorf("readConfig(): Error reading config:", err)
	}

	err = json.Unmarshal(content, &dconf)

	if err != nil {
		return dconf, fmt.Errorf("readConfig(): Error unmarshaling JSON", err)
	}

	return dconf, nil
}

func applyConfigDefaults(conf DangerConfig) DangerConfig {

	if conf.TcpPort == 0 {
		conf.TcpPort = 3344
	}
	if conf.HttpPort == 0 {
		conf.HttpPort = 8080
	}
	if conf.WorkDir == "" {
		conf.WorkDir = "/var/lib/danger/"
	}
	if conf.PidFile == "" {
		conf.PidFile = "danger.log"
	}
	if conf.LogFile == "" {
		conf.LogFile = "danger.log"
	}
	if conf.DBFile == "" {
		conf.DBFile = "danger.db"
	}

	return conf
}
