package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/pocketbase",
			Handler: func(c echo.Context) error {

				return c.File("pocketbase.umd.js")
			},
		})

		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/scan",
			Handler: func(c echo.Context) error {
				b, err := ioutil.ReadFile("./index.html")
				if err != nil {
					fmt.Println("error", err)
					return err
				}
				return c.HTML(200, string(b))
			},
			// Middlewares: []echo.MiddlewareFunc{
			// 	apis.RequireAdminOrUserAuth(),
			// },
		})

		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/login",
			Handler: func(c echo.Context) error {
				b, err := ioutil.ReadFile("./login.html")
				if err != nil {
					fmt.Println("error", err)
					return err
				}
				return c.HTML(200, string(b))
			},
			// Middlewares: []echo.MiddlewareFunc{
			// 	apis.RequireAdminOrUserAuth(),
			// },
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
