package main

func main() {

	router := InitializeRouter()

	router.Run(":8080")
}
