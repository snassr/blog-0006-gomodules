package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/snassr/blog-0006-gomodules/pkg/lyrics"
	"github.com/snassr/blog-0006-gomodules/pkg/sounds"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			sounds.MakeSounds()
			lyrics.WriteLryics()
			time.Sleep(2 * time.Second)
		}
	}()

	<-c
}
