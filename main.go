package main

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/bluza/trackyourself/internal/types"
	"github.com/bluza/trackyourself/internal/views"
	"github.com/labstack/echo"
)

func render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func main() {

	entries := make([]types.TrackEntry, 0)

	entries = append(
		entries,
		types.TrackEntry{
			Date:    time.Now().Add(-time.Hour * 100),
			ImgPath: "asdf",
		})

	entries = append(
		entries,
		types.TrackEntry{
			Date:    time.Now(),
			ImgPath: "asdf",
		})
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", func(ctx echo.Context) error {
		return render(ctx, http.StatusOK, views.Page(&entries))
	})
	e.GET("/modal", func(ctx echo.Context) error {
		return render(ctx, http.StatusOK, views.ImageModal())
	})
	e.Logger.Fatal(e.Start(":8080"))
}
