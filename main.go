package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo/v4"
)

var (
	db *sqlx.DB
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

type Country struct {
	Code        string `json:"code,omitempty"  db:"Code"`
	Name        string `json:"name,omitempty"  db:"Name"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}


func main() {
	fmt.Printf("args count: %d\n", len(os.Args))
    fmt.Printf("args : %#v\n", os.Args)
    for i, v := range os.Args {
        fmt.Printf("args[%d] -> %s\n", i, v)
    }

	_db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	// fmt.Println("Connected!")
	// var city City
	// var country Country
    // if err := db.Get(&city, "SELECT * FROM city WHERE Name=?", os.Args[1]); errors.Is(err, sql.ErrNoRows) {
    //     log.Printf("no such city Name = %s", os.Args[1])
    // } else if err != nil {
    //     log.Fatalf("DB Error: %s", err)
    // }
	// fmt.Printf("%sの人口は%d人です\n", os.Args[1], city.Population)
	
	// if err := db.Get(&country, "SELECT Name, Code, Population FROM country WHERE Code = ?", city.CountryCode); errors.Is(err, sql.ErrNoRows) {
    // } else if err != nil {
    //     log.Fatalf("DB Error: %s", err)
    // }
	// var parsent float64 = float64(city.Population)/float64(country.Population)*100
	// fmt.Printf("%sの人口は%d人です\n", country.Name, country.Population)
	// fmt.Printf("%sの人口は%sの人口の%.2fパーセントです\n", city.Name, country.Name, parsent)


	// cities := []City{}
	// db.Select(&cities, "SELECT * FROM city WHERE CountryCode='JPN'")
	// fmt.Println("日本の都市一覧")
	// for _, city := range cities {
	// 	fmt.Printf("都市名: %s, 人口: %d人\n", city.Name, city.Population)
	// }

	// db.Exec("INSERT INTO city (Name, CountryCode, District, Population) VALUES (?, 'JPN', 'Tokyo', 2147483647)", os.Args[2]); errors.Is(err, sql.ErrNoRows)

	// var city2 City
    // if err := db.Get(&city2, "SELECT * FROM city WHERE Name=?", os.Args[2]); errors.Is(err, sql.ErrNoRows) {
    //     log.Printf("no such city Name = %s", os.Args[2])
    // } else if err != nil {
    //     log.Fatalf("DB Error: %s", err)
    // }
	// fmt.Printf("%sの人口は%d人です\n", os.Args[2], city2.Population)

	db = _db

	e := echo.New()

	e.GET("/cities/:cityName", getCityInfoHandler)
	e.Start(":12400")

}

func getCityInfoHandler(c echo.Context) error {
	cityName := c.Param("cityName")
	fmt.Println(cityName)

	var city City
    if err := db.Get(&city, "SELECT * FROM city WHERE Name=?", cityName); errors.Is(err, sql.ErrNoRows) {
        log.Printf("No Such City Name=%s", cityName)
    } else if err != nil {
        log.Fatalf("DB Error: %s", err)
    }

	return c.JSON(http.StatusOK, city)
}