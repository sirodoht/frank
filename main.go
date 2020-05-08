package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/sirodoht/frank/totp"
)

func main() {
	fmt.Println("Input secret:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		panic("scanner error")
	}

	secret := scanner.Text()

	t := time.Now()
	token, err := totp.GenerateCode(secret, t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("TOTP token: %s\n", token)
}
