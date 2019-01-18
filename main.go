package main

import (
	"log"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	log.SetFlags(log.Lshortfile)

	wd := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{"no-sandbox", "force-device-scale-factor=2", "headless"}),
	)
	if err := wd.Start(); err != nil {
		log.Fatal(err)
	}
	defer wd.Stop()

	p, err := wd.NewPage()
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Navigate("https://s1.fujitv.co.jp/safe/ippon/judgement/index.html"); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	if err := p.Screenshot("image.png"); err != nil {
		log.Fatal(err)
	}
}
