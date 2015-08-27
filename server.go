package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"strconv"
	"net/http"
	"github.com/labstack/echo"
)

type Orientation struct {
	alpha int
	beta int
}

func (o Orientation) ToPixels(maxAlpha int, maxBeta int) Movement {
	x := GetDisplayWidth()/2 + (GetDisplayWidth()/maxAlpha * o.alpha)
	y := GetDisplayHeight()/2 + (GetDisplayHeight()/maxBeta * o.beta)
	return Movement{x,y}
}

type Movement struct {
	x int
	y int
}

func main() {
	e := echo.New()

	e.Index("client/gyro.html")
	e.Get("/move/:alpha/:beta", Move)
	e.Static("/static/", "static/")

	log.Printf("Gyrojoy started at port 8000.")
	e.Run(":8000")
}

func LogToFile(x float64, y float64) {
	f, err := os.OpenFile("points.txt", os.O_APPEND|os.O_WRONLY,0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("%s,%f,%f\n", time.Now(), x, y))
}

func Move(c *echo.Context) error {
	alpha, _ := strconv.ParseInt(c.Param("alpha"), 10, 64)
	beta, _ := strconv.ParseInt(c.Param("beta"), 10, 64)

	MoveMouse(Orientation{int(alpha), int(beta)}.ToPixels(90, 90))

	return c.String(http.StatusOK, "Hello")
}
