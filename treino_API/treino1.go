package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Name  string
	Preco float64
}

var cars []Car

func criarcarros() {
	cars = append(cars, Car{Name: "Ferrari", Preco: 10000})
	cars = append(cars, Car{Name: "Porsche", Preco: 5000})
	cars = append(cars, Car{Name: "Gt", Preco: 8000})
}
func main() {
	criarcarros()
	e := echo.New()
	e.GET("/cars", getcars)
	e.POST("cars", createcar)
	e.Logger.Fatal(e.Start(":8080"))
}

func getcars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createcar(c echo.Context) error {
	car := new(Car)
	if e := c.Bind(car); e != nil {
		return e
	}
	cars = append(cars, *car)
	savecar(*car)
	return c.JSON(200, cars)
}

func savecar(c Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("Insira dentro o (name, price) do carro/ valor($1, $2")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(c.Name, c.Preco)
	if err != nil {
		return err
	}
	return nil
}
