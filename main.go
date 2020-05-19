package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sirodoht/frank/totp"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  frank [secret key]")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  frank CI5LZGDSGSN5VTA")
	os.Exit(0)
}

func main() {
	if len(os.Args) != 2 {
		printHelp()
	}
	if os.Args[1] == "-h" {
		printHelp()
	}
	secret := os.Args[1]
	t := time.Now()
	token, err := totp.GenerateCode(secret, t)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
