package main

import (
	"flag"
	"fmt"
	"github.com/micahhausler/consul-uwsgi-healthcheck/types"
	"github.com/micahhausler/consul-uwsgi-healthcheck/uwsgi"
	"os"
)

var addrPtr = flag.String("address", "127.0.0.1", "The address or IP for uwsgi")
var portPtr = flag.Int("port", 8888, "The port for uwsgi")
var verbosePtr = flag.Bool("verbose", false, "Show verbose output")

func main() {
	Version := "0.0.1"
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println(Version)
		os.Exit(0)
	}

	flag.Parse()

	config := types.Config{
		Address: *addrPtr,
		Port:    *portPtr,
		Verbose: *verbosePtr,
	}

	ponged := uwsgi.Ping(config)
	if ponged == true {
		if *verbosePtr == true {
			fmt.Printf("%s:%d PONG\n", *addrPtr, *portPtr)
		}
		os.Exit(0)
	}
	if *verbosePtr == true {
		fmt.Printf("No PONG!\n")
	}
	os.Exit(2)
}
