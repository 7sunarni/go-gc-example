package main

import (
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

type HeapObj struct {
	Value string
}

var liveHeap string

// go build -o golimit golimit.go; GODEBUG=gctrace=1 GOMEMLIMIT=4096MiB ./golimit
func main() {

	var builder strings.Builder
	go http.ListenAndServe(":8080", nil)
	// 10Mib liveHeap
	for i := 0; i < 1024*1024*10; i++ {
		builder.WriteByte(byte(66))
	}
	liveHeap = builder.String()

	heapValue := ""
	for {
		r := newO()
		heapValue = heapValue + r.Value
		time.Sleep(10 * time.Millisecond)
	}
}

func newO() *HeapObj {
	o := new(HeapObj)
	o.Value = "datadata"
	return o
}
