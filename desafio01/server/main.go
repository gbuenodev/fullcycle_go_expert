package main

import (
	"fmt"
	"net/http"
	"server/internal/app"
	"server/internal/routes"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	defer app.DBConn.Close()

	r := routes.Routes(app)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("App is running on port: 8080")

	server.ListenAndServe()
}
