package main

func main() {
	service := MakeHandler()
	service.HttpServe()
}