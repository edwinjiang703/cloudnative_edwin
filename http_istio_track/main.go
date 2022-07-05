package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	// "runtime/pprof"
	"strings"

	"github.com/cncamp/golang/examples/module1/edwin/http/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/golang/glog"
)

func main() {

	// flag.Parse()
	// var cpuprofile = flag.String("cpuprofile", "/tmp/cpuproinfo", "write cpu profile to fule")
	// f, errcpu := os.Create(*cpuprofile)
	// if errcpu != nil {
	// 	log.Fatal(errcpu)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	// http.HandleFunc("/", rootHandler)
	// http.HandleFunc("/healthz/", healthzHandler)
	metrics.Register()
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz/", healthzHandler)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	mux.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("IP is " + r.RemoteAddr + " HTTP Code health is " + http.StatusText(200))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "1.0")
	for idx, value := range r.Header {
		w.Header().Add(idx, strings.Join(value, "-"))
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	ip := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Println(ip)
	log.Println("IP is " + ip + " HTTP Code root is " + http.StatusText(200))

	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(500)

	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))

	req, err := http.NewRequest("GET", "http://service1", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}
	lowerCaseHeader := make(http.Header)
	for key, value := range r.Header {
		lowerCaseHeader[strings.ToLower(key)] = value
	}
	glog.Info("headers:", lowerCaseHeader)
	req.Header = lowerCaseHeader
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Info("HTTP get failed with error: ", "error", err)
	} else {
		glog.Info("HTTP get succeeded")
	}
	if resp != nil {
		resp.Write(w)
	}
	glog.V(4).Infof("Respond in %d ms", randInt)

}
