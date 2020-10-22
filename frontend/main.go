package main

import (
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

	req, err := http.NewRequest("GET", "http://timeservice:8081/gettime", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	req.Header.Add("x-request-id", r.Header.Get("x-request-id"))
	req.Header.Add("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	req.Header.Add("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	req.Header.Add("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	req.Header.Add("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	req.Header.Add("x-b3-flags", r.Header.Get("x-b3-flags"))
	req.Header.Add("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	fmt.Fprintf(w, "frontend x-request-id:"+r.Header.Get("x-request-id")+"\n")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	for k, v := range r.Header {
		for _, value := range v {
			fmt.Fprintf(w, k+":"+value+"\n")
		}
	}
	fmt.Fprintf(w, "\nReturned body:"+string(body)+"\n")

}

func main() {
	http.HandleFunc("/", handler)
	logger.Info("Listening on 8080...")
	logger.Fatal(http.ListenAndServe(":8080", nil))
}
