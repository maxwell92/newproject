package main

import (
	"yce/config"
	"components/yce"
)

func main() {
	// config.Load()
	config.New().ReadOsEnv().Load()
	// yce.Run()
	yce.Run()
}
