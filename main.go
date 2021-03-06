package main

import (
	"fmt"

	"github.com/go-some/crawler"
)

type Crawler interface {
	Run(crawler.DocsWriter)
}

func main() {
	// 각 사이트의 크롤러를 등록
	crawlers := []Crawler{
		&crawler.MarketWatch{},
		&crawler.Reuters{},
		&crawler.CNBC{},
		&crawler.WallST247{},
		&crawler.USAToday{},
		&crawler.SeekingAlpha{},
		/* 여기에 추가 해주세요*/
	}

	fmt.Println("Run Crawler")

	// mongoDB writer의 구현체를 얻음
	// crawler 패키지의 writer.go에 interface DocsWriter를 구현하는 구현체들이 모여 있음
	wtr := crawler.NewMongoDBWriter()

	// connection pool (collection handler) 생성
	err := wtr.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 크롤러의 실제 구현을 이용해 실행시키는 부분
	for _, crawler := range crawlers {
		crawler.Run(wtr)
	}

	fmt.Println("Fin Crawler")

	// connection pool 제거
	err = wtr.Destroy()
	if err != nil {
		fmt.Println(err)
		return
	}
}
