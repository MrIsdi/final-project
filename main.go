package main

import "final-project/routers"

func main() {
	router := routers.StartApp()
	router.Run()
}
