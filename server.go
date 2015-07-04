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

type Acceleration struct {
	x float64
	y float64
}

var signal chan Acceleration

func main() {
	_ = "breakpoint"
	e := echo.New()

	e.Index("client/gyro.html")
	e.Get("/move/:x/:y", Move)
	e.Static("/static/", "static/")

	signal = make(chan Acceleration, 1)
	go Loop(signal)

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
	x, _ := strconv.ParseFloat(c.Param("x"), 64)
	y, _ := strconv.ParseFloat(c.Param("y"), 64)
	LogToFile(x, y)
	signal <- Acceleration{x, y}
	return c.String(http.StatusOK, "Hello")
}
