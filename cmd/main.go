package main

func main() {
	app, err := CreateApp()
	if err != nil {
		panic(err)
	}
	app.Start()
}
