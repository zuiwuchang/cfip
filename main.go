package main

import (
	"flag"
	"fmt"
	"runtime"
)

var (
	Version, Date, Commit string
)

func main() {
	var (
		conf          string
		version, help bool
	)
	flag.BoolVar(&help, "help", false, "display help")
	flag.BoolVar(&version, "version", false, "display version")
	flag.StringVar(&conf, "conf", "cfip.jsonnet", "configure filepath")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	} else if version {
		fmt.Println("cfip-" + Version)
		fmt.Println(runtime.GOOS+`/`+runtime.GOARCH, Date, Commit)
		return
	}

	

}
