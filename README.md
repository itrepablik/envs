![envs package](https://user-images.githubusercontent.com/58651329/177908485-7a16a57d-96ad-431d-8d47-519a7c8949dc.png)

# Installation
```go
go get -u github.com/itrepablik/envs
```

# envs
The `envs` stands for `environment variables`. It's a simple tool to manage and load your environment variables in one place. It's recommended to load it as early as possible like in your `main` function.

# Usage
This is how you can use the `envs` package in your next Go project.
```go
package main

import (
	"fmt"
	"github.com/itrepablik/envs"
	"os"
)

func main() {
	// Load environment variables from .env files
	err := envs.LoadEnvFromFile("server.env", "server1.env")
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println("MY_ENV_VAR_1: ", os.Getenv("MY_ENV_VAR_1"))
    	fmt.Println("MY_ENV_VAR_5: ", os.Getenv("MY_ENV_VAR_5"))
}
```

# Subscribe to Maharlikans Code Youtube Channel:
Please consider subscribing to my Youtube Channel to recognize my work on any of my tutorial series. Thank you so much for your support!
https://www.youtube.com/c/MaharlikansCode?sub_confirmation=1

# License
Code is distributed under MIT license, feel free to use it in your proprietary projects as well.
