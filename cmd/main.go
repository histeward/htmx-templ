package main

import (
    "htmx-templ/handler"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {

    // instantiate echo
    app := echo.New()

    g := app.Group("/admin")

    g.Use(middleware.Logger())

    g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
        // Be careful to use constant time comparison to prevent timing attacks
        if username == "asdf" && password == "asdf" {
            return true, nil
        }
        return false, nil
    }))

    g.GET("/main", handler.MainAdmin)
    g.GET("/dashboard", handler.DashboardPageHandler)

    // routes
    app.GET("/", handler.LandingPageHandler)
    app.GET("/login", handler.LoginPageHandler)

    // specify port
    // app.Logger.Fatal(app.Start(":3000"))
    app.Start(":3000")
}
