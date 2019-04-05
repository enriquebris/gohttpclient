[![godoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/enriquebris/gohttpclient) [![Go Report Card](https://goreportcard.com/badge/github.com/enriquebris/gohttpclient)](https://goreportcard.com/report/github.com/enriquebris/gohttpclient) [![Build Status](https://travis-ci.org/enriquebris/gohttpclient.svg?branch=master)](https://travis-ci.org/enriquebris/gohttpclient)

# GOlang HTTP client

### What is this?

gohttpclient provides a HTTP Client interface with basic functions to make HTTP requests.

### Benefits of using gohttpclient.HTTPClient interface

This is all Dependency Inversion Principle (SOLID). Relying on abstractions (that would be interfaces in golang) would avoid any dependency over a particular implementation. That's why I encourage you to use gohttpclient.HTTPClient in case you need to make HTTP requests. This way you will be able to switch over different implementations as long as your project evolves.

Let's say that when you started your project the [golang native http package](https://golang.org/pkg/net/http/) was a good fit but after one year you decide to replace it by [fasthttp](https://github.com/valyala/fasthttp) or by any other golang HTTP package. Other than switch to a different gohttpclient.HTTPClient implementation, your code will not be modified at all. Even in the case that there were not an implementation for the package you think would fit better, you could write it ... and your original code would keep the same.  

### gohttpclient implementations

 - NativeHTTPClient - GOlang native HTTP 

### Installation

```bash
go get github.com/enriquebris/gohttpclient
```

### Examples

#### Simple curl
```go
package main

import (
	"fmt"

	"github.com/enriquebris/gohttpclient"
)

func main() {
	// golang native HTTP implementation
	httpClient := gohttpclient.NewNativeHTTPClient()

	body, err := simpleCURL(httpClient, "GET", "https://ifconfig.co/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ip: %v\n", body)
}

func simpleCURL(httpClient gohttpclient.HTTPClient, method string, url string) (string, error) {
	// reset any previous parameter
	httpClient.Reset()
	// set method (verb)
	httpClient.SetMethod(method)
	// set url
	httpClient.SetURL(url)

	// make the request
	if resp, err := httpClient.Do(); err != nil {
		return "", err
	} else {
		return resp.GetBody(), nil
	}
}
```

#### Bridge pattern
[gobetafaceapi](https://github.com/enriquebris/gobetafaceapi) - gohttpclient is used as the [Implementor](https://en.wikipedia.org/wiki/Bridge_pattern).

### TODO

- [x] GOlang native HTTP
- [ ] GOlang native HTTP testings 
- [ ] [fasthttp](https://github.com/valyala/fasthttp) implementation