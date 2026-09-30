package main

import (
	"fmt"

	"config-demo/env"
)

func main() {
	fmt.Println("APP_ENV =", env.AppEnv)
}
