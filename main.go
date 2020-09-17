package main

import (
	"github.com/taschenbergerm/pokescraper/log"
	"github.com/taschenbergerm/pokescraper/pkg"
)

func main() {
	log.Info("Starting Main.go")
	pkg.Main()
	log.Info("End Main.go")
}
