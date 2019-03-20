package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/snassr/blog-0006-gomodules/pkg/lyrics"
	"github.com/snassr/blog-0006-gomodules/pkg/sounds"
	listener "github.com/snassr/blog-0006-listener"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			sounds.MakeSounds()
			lyrics.WriteLryics()
			listener.Listener()
			time.Sleep(2 * time.Second)
		}
	}()

	<-c
}
