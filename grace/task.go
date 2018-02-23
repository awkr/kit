package grace

import (
	"sync"
	"sync/atomic"
)

var (
	isStopped uint32
	wg        sync.WaitGroup
	quit      = make(chan bool)
)

func Add() bool {
	if atomic.LoadUint32(&isStopped) == 1 {
		return false
	}

	wg.Add(1)
	println("grace: add")

	return true
}

func Done() {
	wg.Done()
	println("grace: done")
}

func Stop() {
	println("grace: stop")
	atomic.StoreUint32(&isStopped, 1)
}

func Wait() {
	println("grace: wait")
	wg.Wait()
}

func Push() {
	println("grace: push")
	quit <- true
}

func Pop() {
	println("grace: pop")
	<-quit
}
