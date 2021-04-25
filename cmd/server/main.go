package main

import (
	"fmt"
	"os"

	"go-restish/internal/database"
	internalHttp "go-restish/internal/http"
	"go-restish/internal/services/user"

	"github.com/valyala/fasthttp"
)

// Run - entrypoint of the api
func Run() error {
	fmt.Println("Running App")

	os.Setenv("TZ", "UTC") // set server timezone to UTC

	db, err := database.InitDatabase()

	if err != nil {
		return err
	}

	userService := user.NewService(db)

	h := internalHttp.NewHandler(userService)
	h.InitRouter()

	if err := fasthttp.ListenAndServe(":8080", h.Router.Handler); err != nil {
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
