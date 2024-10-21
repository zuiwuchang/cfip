package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/zuiwuchang/cfip/cf"
	"github.com/zuiwuchang/cfip/configure"
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
	c, e := configure.Load(conf)
	if e != nil {
		log.Fatalln(e)
	}

	ctx, e := cf.New(c)
	if e != nil {
		log.Fatalln(e)
	}

	e = ctx.Serve()
	if e != nil {
		os.Exit(1)
	}
}
