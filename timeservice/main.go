package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"time"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("x-request-id", r.Header.Get("x-request-id"))
	w.Header().Set("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	w.Header().Set("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	w.Header().Set("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	w.Header().Set("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	w.Header().Set("x-b3-flags", r.Header.Get("x-b3-flags"))
	w.Header().Set("x-ot-span-context", r.Header.Get("x-ot-span-context"))


	fmt.Fprintf(w,"timeservice x-request-id:" +  r.Header.Get("x-request-id") + "\n")

	now := time.Now()
        w.Write([]byte(now.String()))

	
	/*
	req, err = http.NewRequest("GET", "http://service_a_envoy:8791/", nil)
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

	client = &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Fprintf(w, string(body))
	*/
}

func main() {
	http.HandleFunc("/gettime", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

