package main

import (
	"backend/model"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var cache = redis.NewClient(&redis.Options{
  Addr: "localhost:6379",
})

type Env struct {
  db *sql.DB
}

func main() {
  db, err := sql.Open(
    "mysql", "newuser:1111@tcp(172.22.208.1:3306)/countries",
    )

  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err.Error())
  }
  fmt.Println("Connected to database server")

  // model.TestCountries(db)
  // model.TestCurrency(db)

  r := gin.Default()

  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://localhost:3000"}
  r.Use(cors.New(config))

  r.GET("/", func(c *gin.Context) {
    data := model.GetAllCountries(db)
    c.JSON(http.StatusOK, data)
  })

  r.GET("/filter", func(c *gin.Context) {
    url := c.Request.URL.Query()
    min, err := strconv.Atoi(url.Get("min"))
    if err != nil { panic(err) }
    max, err := strconv.Atoi(url.Get("max"))
    if err != nil { panic(err) }
    data := model.GetCountriesByPopulation(db, min, max)
    c.JSON(http.StatusOK, data)
  })


  r.Run(":8080")
}
