package main

import (
  "fmt"
  "crawler"
  "encoding/csv"
  "os"
  "log"
)

type Crawler interface {
  Run(crawler.Writer)
}

func main() {
  // 각 사이트의 크롤러를 등록
  crawlers := []Crawler{
    &crawler.USToday{},
    &crawler.MarketWatch{},
  }

  fmt.Println("Run Crawler")

  // 실제 데이터를 저장소에 쓰는 부분 입니다.
  // 추후에 데이터베이스로 쓰기 연산을 수행하도록 수정.
  wtr := func (fName string, docs []crawler.News) {
    file, err := os.Create(fName)
    if err != nil {
      log.Fatalf("Cannot create file %q: %s\n", fName, err)
      return
    }
    defer file.Close()
    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write([]string{"Title", "Body", "Time", "Url", "Origin"})

    for _, doc := range docs {
      writer.Write([]string{
        doc.Title,
        doc.Body,
        doc.Time,
        doc.Url,
        doc.Origin,
      })
    }
  }

  // 크롤러의 실제 구현을 이용해 실행시키는 부분
  for _, crawler := range crawlers {
    crawler.Run(wtr)
  }

  fmt.Println("Fin Crawler")
}
