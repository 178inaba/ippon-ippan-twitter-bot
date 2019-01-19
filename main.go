package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
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

	dst := image.NewRGBA(image.Rect(0, 0, 980, 540))
	draw.Draw(dst, dst.Bounds(), img, image.Pt(310, 430), draw.Src)

	b := bytes.NewBuffer(nil)
	if err := png.Encode(b, dst); err != nil {
		log.Fatal(err)
	}

	// Post twitter
	api := anaconda.NewTwitterApiWithCredentials(
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET"),
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"))

	m, err := api.UploadMedia(base64.StdEncoding.EncodeToString(b.Bytes()))
	if err != nil {
		log.Fatal(err)
	}

	v := url.Values{}
	v.Set("media_ids", m.MediaIDString)
	t, err := api.PostTweet("#IPPON", v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.Id)
}
