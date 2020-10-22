package main

import (
	//"fmt"
	//"io/ioutil"

	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	req, _ := http.NewRequestWithContext(ctx, "GET", "http://timeapi:8082/api/v1/time", nil)

	req.Header.Add("x-request-id", r.Header.Get("x-request-id"))
	req.Header.Add("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	req.Header.Add("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	req.Header.Add("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	req.Header.Add("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	req.Header.Add("x-b3-flags", r.Header.Get("x-b3-flags"))
	req.Header.Add("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	tclient := &http.Client{}
	tresp, _ := tclient.Do(req)
	tbody, _ := ioutil.ReadAll(tresp.Body)
	tresp.Body.Close()

	dreq, err := http.NewRequestWithContext(ctx, "GET", "http://dayapi:8083/api/v1/day", nil)
	if err != nil {
		logger.Fatal()
	}

	dreq.Header.Add("x-request-id", r.Header.Get("x-request-id"))
	dreq.Header.Add("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	dreq.Header.Add("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	dreq.Header.Add("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	dreq.Header.Add("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	dreq.Header.Add("x-b3-flags", r.Header.Get("x-b3-flags"))
	dreq.Header.Add("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	dclient := &http.Client{}
	dresp, _ := dclient.Do(dreq)
	dbody, _ := ioutil.ReadAll(dresp.Body)
	dresp.Body.Close()

	response := "Day:" + string(dbody) + "\n" + "Date:" + string(tbody) + "\n"

	fmt.Fprintf(w, string(response))
	cancel()

}

func main() {
	http.HandleFunc("/gettime", handler)
	logger.Info("Listening on 8081...")
	logger.Fatal(http.ListenAndServe(":8081", nil))
}
