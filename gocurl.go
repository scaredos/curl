package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Print the help options or "banner" function
func banner() {
	fmt.Printf("Usage: %s <url> <options>\n", os.Args[0])
	fmt.Println("-h, --help\t display this menu")
	fmt.Println("-H, --headers\t \"<headers>\"")
	fmt.Println("-M, --method\t \"<method GET/POST/HEAD>\"")
	fmt.Println("-A, --user-agent \"<user agent>\"")
	fmt.Println("-S, --silent\t \"silence output\"")
	fmt.Println("-v, --verbose\t \"verbose output\"")
}

func main() {
	var headers []string
	var data []string
	var url string
	var verbose bool
	var silent bool
	defaultUa := "gcurl/1.1"
	ua := ""
	method := "GET"
	// parsing os arguments looking for headers
	if len(os.Args) <= 1 {
		banner()
		os.Exit(0)
	}
	for i, ch := range os.Args {
		if i == 0 {
			continue
		}
		if strings.Contains(ch, "http") {
			url = ch
		} else if strings.Contains(ch, "-H") || strings.Contains(ch, "--headers") {
			headers = append(headers, strings.Split(os.Args[i+1], ":")[0])
			tempItem := strings.Split(os.Args[i+1], ":")[1]
			data = append(data, strings.Split(tempItem, " ")[1])
		} else if strings.Contains(ch, "-M") || strings.Contains(ch, "--method") {
			if method != "GET" && method != "POST" && method != "HEAD" {
				banner()
				os.Exit(0)
			}
			method = os.Args[i+1]
		} else if strings.Contains(ch, "-v") || strings.Contains(ch, "--verbose") {
			verbose = true
		} else if strings.Contains(ch, "-A") || strings.Contains(ch, "--user-agent") {
			ua = os.Args[i+1]
		} else if strings.Contains(ch, "-s") || strings.Contains(ch, "--silent") {
			silent = true
		} else if strings.Contains(ch, "-h") || strings.Contains(ch, "--help") {
			banner()
			os.Exit(0)
		} else {
			continue
		}
	}
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	if ua == "" {
		req.Header.Set("User-Agent", defaultUa)
	} else {
		req.Header.Set("User-Agent", ua)
	}
	for i, _ := range headers {
		req.Header.Set(headers[i], data[i])
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if verbose {
		fmt.Printf("> %s %s %s\n", resp.Request.Method, resp.Request.RequestURI, resp.Request.Proto)
		for key, element := range resp.Request.Header {
			fmt.Printf("> %s: %s\n", key, element)
		}
		fmt.Printf("< %s\n", resp.Status)
		fmt.Printf("< %s\n", resp.Proto)
		for key, element := range resp.Header {
			fmt.Printf("< %s: %s\n", key, element)
		}
		if resp.TLS != nil {
			fmt.Printf("<> %s\n", resp.TLS.NegotiatedProtocol)
			fmt.Printf("<> %s\n", resp.TLS.CipherSuite)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		if silent {
			//
		} else {
			fmt.Printf(string(body))
		}
	} else {
		if silent {
			//
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf(string(body))
		}
	}
}
