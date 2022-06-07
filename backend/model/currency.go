package model

import (
	"database/sql"
	"fmt"
)

type Currency struct {
  Name string
}

func GetCurrenciesFromCountry(db *sql.DB, country string) []Currency {
  
  results, err := db.Query(
    "SELECT currency FROM country_currency WHERE country = ?",
    country,
    )
  if err != nil {
    panic(err.Error())
  }

  currencies := []Currency{}

  for results.Next() {
    var c Currency
    err = results.Scan(
      &c.Name,
      )
    if err != nil {
      panic(err.Error())
    }
    currencies = append(currencies, c)
  }
  return currencies
}

func TestCurrency(db *sql.DB) {
  a := GetCurrenciesFromCountry(db, "Thailand")
  fmt.Println(a)
}
