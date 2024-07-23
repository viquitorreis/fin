package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()
	e := echo.New()

	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.GET("/tickers", GetTickersList)

	e.Logger.Fatal(e.Start(":8080"))
}

func GetTickersList(c echo.Context) error {
	brapiList := os.Getenv("BRAPI_URL") + "/quote/list"

	client := &http.Client{}
	req, err := http.NewRequest("GET", brapiList, nil)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("BRAPI_TOKEN"))

	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	var response Response
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, response)
}
