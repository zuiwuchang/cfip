package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/google/go-jsonnet"
	"github.com/zuiwuchang/cfip/cf"
	"github.com/zuiwuchang/cfip/configure"
)

var (
	Version, Date, Commit string
)

func main() {
	var (
		conf, url            string
		version, help, print bool
	)
	flag.BoolVar(&help, "help", false, "display help")
	flag.BoolVar(&version, "version", false, "display version")
	flag.BoolVar(&print, "print", false, "display configure json")
	flag.StringVar(&conf, "conf", "cfip.jsonnet", "configure filepath")
	flag.StringVar(&url, "url", "http://127.0.0.1:8080", "send the results as post json to this url")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	} else if version {
		fmt.Println("cfip-" + Version)
		fmt.Println(runtime.GOOS+`/`+runtime.GOARCH, Date, Commit)
		return
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	if print {
		vm := jsonnet.MakeVM()
		jsonStr, e := vm.EvaluateFile(conf)
		if e != nil {
			log.Fatalln(e)
			return
		}
		fmt.Println(jsonStr)
		return
	}

	c, e := configure.Load(conf)
	if e != nil {
		log.Fatalln(e)
	}
	if url != `` {
		c.Found.URL = url
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
