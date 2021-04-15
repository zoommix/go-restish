package main

import (
	"fmt"

	"go-restish/internal/database"
	internalHttp "go-restish/internal/http"

	"github.com/valyala/fasthttp"
)

// Run - entrypoint of the api
func Run() error {
	fmt.Println("Running App")

	_, err := database.InitDatabase()

	if err != nil {
		return err
	}

	h := internalHttp.NewHandler()
	h.InitRouter()

	if err := fasthttp.ListenAndServe(":"+"8080", h.Router.Handler); err != nil {
		return err
	}

	return nil
}

func main() {
	err := Run()

	if err != nil {
		fmt.Println("Error running app")
		fmt.Println(err)
	}
}
