package handler

import (
    // "net/http"
    "github.com/labstack/echo/v4"
    "net/http"
    "htmx-templ/view/landing"
    "htmx-templ/view/login"
    "htmx-templ/view/dashboard"
)

func LandingPageHandler(c echo.Context) error {
    return landing.Show().Render(c.Request().Context(), c.Response())
}

func DashboardPageHandler(c echo.Context) error {
    return dashboard.Show().Render(c.Request().Context(), c.Response())
}

func LoginPageHandler(c echo.Context) error {
    return login.Show().Render(c.Request().Context(), c.Response())
}

func MainAdmin(c echo.Context) error {
    return c.String(http.StatusOK, "ayayayay the secret admin page")
}
