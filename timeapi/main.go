package main

import (
	//"fmt"
	//"io/ioutil"

	"log"
	"net/http"
	"os"
	"time"

	logger "github.com/sirupsen/logrus"
)

func init() {
	logger.SetFormatter(&logger.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logger.InfoLevel)
}

func handler(w http.ResponseWriter, r *http.Request) {

	logger.WithFields(logger.Fields{
		"Client IP":      r.RemoteAddr,
		"Tracer Req ID:": r.Header.Get("x-request-id"),
	}).Info("Debug")

	w.Header().Set("x-request-id", r.Header.Get("x-request-id"))
	w.Header().Set("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	w.Header().Set("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	w.Header().Set("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	w.Header().Set("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	w.Header().Set("x-b3-flags", r.Header.Get("x-b3-flags"))
	w.Header().Set("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	t := time.Now()
	//d := time.Now().Weekday()
	w.Write([]byte(t.Format("2006-01-02 15:04:05")))
}

func main() {
	http.HandleFunc("/api/v1/time", handler)
	logger.Info("Listening on 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
