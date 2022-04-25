package main

import (
	"craw/craw/dzy"
	"craw/craw/sink"
	"craw/craw/wallpaper"
	"fmt"
	"github.com/robfig/cron"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ender := make(chan os.Signal, 1)
	signal.Notify(ender, os.Interrupt, os.Kill, syscall.SIGTERM)
	startTime := time.Now()

	go sink.Sinking()
	c := cron.New()
	c.AddFunc("0 0 2 * * ?", func() {
		fmt.Println("begin work")
		wallpaper.CrawWallPaperInOneColly()
	})
	//wallpaper.CrawWallPaperInOneColly()
	dzy.Craw()
	c.Start()

	fmt.Printf("[main] started in %3.2f seconds.", time.Now().Sub(startTime).Seconds())
	<-ender
}
