package main

import (
	"components/yce"
	"yce/config"
)

func main() {
	// config.Load()
	config.New().ReadOsEnv().Load()
	// yce.Run()
	yce.Run()
}
