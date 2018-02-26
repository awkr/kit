package grace

import (
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/solarwalker/kit/log"
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

	log.DebugFlag("grace: add", debugEnabled)

	return true
}

func Done() {
	wg.Done()

	log.DebugFlag("grace: done", debugEnabled)

}

func Stop() {
	log.DebugFlag("grace: stop", debugEnabled)

	atomic.StoreUint32(&isStopped, 1)
}

func Wait() {
	log.DebugFlag("grace: wait", debugEnabled)

	wg.Wait()
}

func Push() {
	log.DebugFlag("grace: push", debugEnabled)

	quit <- true
}

func Pop() {
	log.DebugFlag("grace: pop", debugEnabled)

	<-quit
}
