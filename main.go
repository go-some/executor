package main

import (
  "fmt"
  "crawler"
)

type Crawler interface {
  Run()
}

func main() {
  crawlers := []Crawler{
    &crawler.USToday{},
    &crawler.MarketWatch{},
  }

  fmt.Println("Run Crawler")

  for _, crawler := range crawlers {
    crawler.Run()
  }

  fmt.Println("Fin Crawler")
}
