package main

import (
	"fmt"
	"os"

	"github.com/ynotnauk/go/pkg/dotenv"
)

func main() {
	dotenv.Load()
	fmt.Println(os.Getenv("TWITCH_CLIENT_ID"))
}
