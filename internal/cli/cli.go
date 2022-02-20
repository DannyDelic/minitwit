package cli

import (
	"flag"
	"fmt"
	"minitwit/internal/server"
	"os"
)

func usage() {
	fmt.Print(`This program runs the minitwit backend server.
 
Usage:
 
minitwit [arguments]
 
Supported arguments:
 
`)
	flag.PrintDefaults()
	os.Exit(1)
}

func Parse() {
	flag.Usage = usage
	env := flag.String("env", "dev", `Sets runtime environment. Possible values are "dev" and "prod"`)
	flag.BoolVar(&server.IS_SIM, "s", false, "Turns on sim-mode for DevOps course")
	flag.Parse()
	fmt.Println(env)
	fmt.Print("Server is handling simulator requests: ")
	fmt.Println(server.IS_SIM)
}
