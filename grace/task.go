package grace

import (
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	debugEnabled bool

	isStopped uint32
	wg        sync.WaitGroup
	quit      = make(chan bool)
)

func init() {
	debugEnabled = strings.ToLower(os.Getenv("GRACE_DEBUG")) == "true"
}

func Add() bool {
	if atomic.LoadUint32(&isStopped) == 1 {
		return false
	}

	wg.Add(1)

	debug("grace: add")

	return true
}

func Done() {
	wg.Done()

	debug("grace: done")

}

func Stop() {
	debug("grace: stop")

	atomic.StoreUint32(&isStopped, 1)
}

func Wait() {
	debug("grace: wait")

	wg.Wait()
}

func Push() {
	debug("grace: push")

	quit <- true
}

func Pop() {
	debug("grace: pop")

	<-quit
}

func debug(msg string) {
	if debugEnabled {
		println(msg)
	}
}
