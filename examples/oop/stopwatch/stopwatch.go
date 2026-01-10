package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start time.Time
	results []time.Duration
}

func (sw *Stopwatch) Start() {
	sw.start = time.Now()
}

func (sw *Stopwatch) SaveSplit() {
	sw.results = append(sw.results, time.Since(sw.start))
}

func (sw Stopwatch) GetResults() []time.Duration {
	return sw.results
}

func main() {
    sw := Stopwatch{}
   	sw.Start()

    time.Sleep(1 * time.Second)
    sw.SaveSplit()

    time.Sleep(500 * time.Millisecond)
    sw.SaveSplit()

    time.Sleep(300 * time.Millisecond)
    sw.SaveSplit()

    fmt.Println(sw.GetResults())
} 