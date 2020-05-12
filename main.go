package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sirodoht/frank/totp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("frank — CLI TOTP generator")
		fmt.Println("")
		fmt.Println("Usage: frank [secret key]")
		fmt.Println("")
		fmt.Println("Example:")
		fmt.Println("    frank C2DGYR26844297SAHI2SZW")
		os.Exit(0)
	}
	secret := os.Args[1]
	t := time.Now()
	token, err := totp.GenerateCode(secret, t)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
