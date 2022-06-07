package model

import (
	"database/sql"
	"fmt"
)

type Country struct {
  Name string
  Flag string
  Population int
  Region string
  Currencies []Currency
}

func GetAllCountries(db *sql.DB) []Country {
  results, err := db.Query(
    "SELECT * FROM country",
    )
  if err != nil {
    panic(err.Error())
  }

  countries := []Country{}

  for results.Next() {
    var c Country
    err = results.Scan(
      &c.Name,
      &c.Flag,
      &c.Population,
      &c.Region,
      )
    if err != nil {
      panic(err.Error())
    }
    c.Currencies = GetCurrenciesFromCountry(db, c.Name)
    countries = append(countries, c)
  }
  return countries
}

func GetCountriesByRegion(db *sql.DB, region string) []Country {
  results, err := db.Query(
    "SELECT * FROM country WHERE region = ?",
    region,
    )
  if err != nil {
    panic(err.Error())
  }

  countries := []Country{}

  for results.Next() {
    var c Country
    err = results.Scan(
      &c.Name,
      &c.Flag,
      &c.Population,
      &c.Region,
      )
    if err != nil {
      panic(err.Error())
    }
    c.Currencies = GetCurrenciesFromCountry(db, c.Name)
    countries = append(countries, c)
  }
  return countries
}

func GetCountriesByPopulation(db *sql.DB, min int, max int) []Country {
  results, err := db.Query(
    "SELECT * FROM country WHERE population >= ? AND population <= ?",
    min, max,
    )
  if err != nil {
    panic(err.Error())
  }

  countries := []Country{}

  for results.Next() {
    var c Country
    err = results.Scan(
      &c.Name,
      &c.Flag,
      &c.Population,
      &c.Region,
      )
    if err != nil {
      panic(err.Error())
    }
    c.Currencies = GetCurrenciesFromCountry(db, c.Name)
    countries = append(countries, c)
  }
  return countries
}


func GetAllRegions(db *sql.DB) []string {
  results, err := db.Query(
    "SELECT DISTINCT region FROM country",
    )
  if err != nil {
    panic(err.Error())
  }

  regions := []string{}

  for results.Next() {
    var c string
    err = results.Scan(
      &c,
      )
    if err != nil {
      panic(err.Error())
    }
    regions = append(regions, c)
  }
  return regions
}

func TestCountries(db *sql.DB) {
  // a := GetAllCountries(db)
  // fmt.Println(a)
  b := GetCountriesByRegion(db, "Antarctic Ocean")
  fmt.Println(b)
  // c := GetCountriesByPopulation(db, 100, 200)
  // fmt.Println(c)
  // d := GetAllRegions(db)
  // fmt.Println(d)
}
