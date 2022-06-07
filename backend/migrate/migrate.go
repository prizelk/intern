package migrate

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Country struct {
  Name string `json:"name"`
  Flag string `json:"flag"`
  Population int`json:"population"`
  Region string `json:"region"`
  Currencies []Currency `json:"currencies"`
}

type Currency struct {
  Name string `json:"name"`
}

func populateData(db *sql.DB) {
  resp, err := http.Get("https://restcountries.com/v2/all")

  if err != nil {
    panic(err.Error())
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err.Error())
  }

  var result []Country
  if err := json.Unmarshal([]byte(body), &result); err != nil {
    fmt.Println("Cannot unmarshal JSON")
  }

  for _, e := range result {
    AddCountry(db, e)
    AddCurrencies(db, e.Currencies)
    AddCountryCurrency(db, e)
  }
}

func AddCountry(db *sql.DB, country Country) {
  insert, err := db.Query(
    "INSERT INTO country VALUES (?, ?, ?, ?)",
    country.Name, country.Flag, country.Population, country.Region,
    )

  if err != nil {
    panic(err.Error())
  }
  fmt.Println("Added new country")

  defer insert.Close()
}

func AddCurrencies(db *sql.DB, currencies []Currency) {
  for _, currency := range currencies {
    insert, err := db.Query(
      "INSERT IGNORE INTO currency VALUES (?)",
      currency.Name,
      )
    if err != nil {
      fmt.Println("dupe currency")
    }
    defer insert.Close()
  }
}

func AddCountryCurrency(db *sql.DB, country Country) {
  for _, currency := range country.Currencies {
    insert, err := db.Query(
      "INSERT INTO country_currency VALUES (?, ?)",
      country.Name, currency.Name, 
      )
    if err != nil {
      fmt.Println("currency error")
    }
    defer insert.Close()
  }
}
