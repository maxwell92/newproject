package yce

import (
	"yce/config"
	"components/yce"
)

func main() {
	// config.Load()
	config.New().Load()
	// yce.Run()
	yce.Run()
}
