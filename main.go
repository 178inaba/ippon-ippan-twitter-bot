package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
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

	filename := "/tmp/ss.png"
	if err := p.Screenshot(filename); err != nil {
		log.Fatal(err)
	}

	// Crop
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	dst := image.NewRGBA(image.Rect(0, 0, 980, 480))
	draw.Draw(dst, dst.Bounds(), img, image.Pt(310, 430), draw.Src)

	dstFile, err := os.Create("/tmp/dst.png")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	if err := png.Encode(dstFile, dst); err != nil {
		log.Fatal(err)
	}
}
