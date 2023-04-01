package main

import "challenge-2/routers"

func main() {
	var PORT = ":8081"

	routers.StartServer().Run(PORT)
}
