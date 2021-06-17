package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	go heartbeatLoop()
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9090", nil)
	panic(err)
}

var heartbeats = promauto.NewCounter(prometheus.CounterOpts{
	Name: "heartbeat_total",
	Help: "The total number of processed heartbeat events.",
})

func heartbeatLoop() {
	for {
		fmt.Printf("Ping! Current timestamp is: %d\n", time.Now().Unix())
		heartbeats.Inc()
		time.Sleep(time.Second * 1)
	}
}
