package yce

import (
	"components/yce/resource/deployment"
)

func AddControllers() {
	deployment.New().AddControllers()
}
