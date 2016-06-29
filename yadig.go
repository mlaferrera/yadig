/*

   yadig.go

   yadig (pronounced "you dig") allows for DNS queries from the command line
   using Google's HTTPS DNS service.

   author: @mlaferrera
*/

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var version = "0.1"
var apiURL = "https://dns.google.com/resolve"
var ednsSubnet = "0.0.0.0/0"

type Response struct {
	Status int  `json:"Status"`
	TC     bool `json:"TC"`
	RD     bool `json:"RD"`
	RA     bool `json:"RA"`
	AD     bool `json:"AD"`
	CD     bool `json:"CD"`

	Question []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"Question"`

	Answer []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Answer"`

	Authority []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Authority"`

	Additional       []interface{} `json:"Additional"`
	EdnsClientSubnet string        `json:"edns_client_subnet"`
	Comment          string        `json:"Comment"`
}

func formattedPrint(r *chan []byte, resultWg *sync.WaitGroup) {
	rcodes := map[int]string{
		0: "NoError",
		1: "FormatError",
		2: "ServerFailure",
		3: "NameError",
		4: "NotImplemented",
		5: "Refused",
	}

	rrtypes := map[int]string{
		1:  "A",
		2:  "NS",
		5:  "CNAME",
		6:  "SOA",
		12: "PTR",
		15: "MX",
		16: "TXT",
		28: "AAAA",
	}

	for {

		response := &Response{}

		resp := <-*r

		err := json.Unmarshal(resp, &response)
		if err != nil {
			fmt.Println(err)
			resultWg.Done()
			continue
		}

		dnssec := response.AD
		rcode := response.Status
		query := response.Question[0].Name

		if rcode > 0 {
			fmt.Printf("ERROR: %v %v %v\n", rcodes[rcode], query, response.Comment)
			resultWg.Done()
			continue
		}

		for _, v := range response.Answer {
			fmt.Printf("Query: %v, ", v.Name)
			fmt.Printf("DNSSEC: %v, ", dnssec)

			var rrtype string

			if _, ok := rrtypes[v.Type]; ok {
				rrtype = rrtypes[v.Type]
			} else {
				rrtype = string(v.Type)
			}

			fmt.Printf("Type: %v, ", rrtype)
			fmt.Printf("TTL: %v, ", v.TTL)
			fmt.Printf("Response: %v", v.Data)
			fmt.Println()
		}

		for _, v := range response.Authority {
			fmt.Printf("Query: %v, ", query)
			fmt.Printf("DNSSEC: %v, ", dnssec)

			var rrtype string

			if _, ok := rrtypes[v.Type]; ok {
				rrtype = rrtypes[v.Type]
			} else {
				rrtype = string(v.Type)
			}

			fmt.Printf("Type: %v, ", rrtype)
			fmt.Printf("TTL: %v, ", v.TTL)
			fmt.Printf("Response: %v", v.Data)
			fmt.Println()
		}

		resultWg.Done()
	}
}

func resolveHost(q *chan map[string]string, r *chan []byte, queryWg *sync.WaitGroup, resultWg *sync.WaitGroup) {
	for {

		client := &http.Client{}

		req, _ := http.NewRequest("GET", apiURL, nil)

		queryMap := <-*q
		query := req.URL.Query()
		query.Add("name", queryMap["name"])
		query.Add("type", queryMap["type"])
		query.Add("edns_client_subnet", ednsSubnet)

		req.URL.RawQuery = query.Encode()

		resp, err := client.Do(req)
		defer resp.Body.Close()

		if err != nil {
			fmt.Println("Unable to resolve host: ", resp.Status)
			queryWg.Done()
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)

		*r <- body

		resultWg.Add(1)
		queryWg.Done()
	}
}

func main() {

	queryHost := flag.String("q", "", "Host to conduct a query on")
	queryType := flag.String("t", "A", "DNS record type")
	flag.Parse()

	queryMap := make(map[string]string)

	q := make(chan map[string]string)
	r := make(chan []byte)

	var queryWg sync.WaitGroup
	var resultWg sync.WaitGroup

	go resolveHost(&q, &r, &queryWg, &resultWg)
	go formattedPrint(&r, &resultWg)

	if len(*queryHost) > 0 {
		queryWg.Add(1)
		queryMap["type"] = *queryType
		queryMap["name"] = *queryHost
		q <- queryMap
	} else {
		buf := bufio.NewScanner(os.Stdin)
		for buf.Scan() {
			queryMap["type"] = *queryType
			queryMap["name"] = buf.Text()
			queryWg.Add(1)
			q <- queryMap
		}
	}

	queryWg.Wait()
	resultWg.Wait()
}
