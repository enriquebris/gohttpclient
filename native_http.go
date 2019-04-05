package gohttpclient

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

// ************************************************************************************************
// ** NativeHTTPTestServer
// ************************************************************************************************
// ** HTTPTestServer implementation using native GO *httptest.Server
// ************************************************************************************************

func NewNativeHTTPTestServer(testServer *httptest.Server) *NativeHTTPTestServer {
	ret := &NativeHTTPTestServer{
		testServer: testServer,
	}

	return ret
}

type NativeHTTPTestServer struct {
	testServer *httptest.Server
}

func (st *NativeHTTPTestServer) GetURL() string {
	return st.testServer.URL
}

func (st *NativeHTTPTestServer) Close() {
	st.testServer.Close()
}

// ************************************************************************************************
// ** NativeHTTPClient
// ************************************************************************************************
// ** HTTPClient implementation using GO native http
// ************************************************************************************************

func NewNativeHTTPClient() HTTPClient {
	ret := &NativeHTTPClient{}
	ret.initialize()

	return ret
}

type NativeHTTPClient struct {
	method, url, payload string
	headers              map[string]string
}

func (st *NativeHTTPClient) initialize() {
	st.Reset()
}

func (st *NativeHTTPClient) Reset() {
	st.headers = make(map[string]string)
	st.method = ""
	st.url = ""
	st.payload = ""
}

func (st *NativeHTTPClient) SetMethod(method string) {
	st.method = method
}

func (st *NativeHTTPClient) SetURL(url string) {
	st.url = url
}

func (st *NativeHTTPClient) SetPayload(payload string) {
	st.payload = payload
}

func (st *NativeHTTPClient) AddHeader(key string, value string) {
	st.headers[key] = value
}

func (st *NativeHTTPClient) Do() (HTTPResponse, error) {
	var (
		req *http.Request
		err error
	)
	if st.payload != "" {
		payload := strings.NewReader(st.payload)
		req, err = http.NewRequest(st.method, st.url, payload)
	} else {
		req, err = http.NewRequest(st.method, st.url, nil)
	}
	if err != nil {
		return nil, err
	}

	// add the headers
	for k, v := range st.headers {
		req.Header.Add(k, v)
	}

	// do the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return NewDefaultHTTPResponse(string(body), res.StatusCode, res.Status, res.Header), nil
}

func (st *NativeHTTPClient) NewTestServer(fn func(ResponseWriter, Request)) HTTPTestServer {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO ::: call the user function here !!!
		fn(NewDefaultResponseWriter(w), nil)
	}))

	return NewNativeHTTPTestServer(ts)
}
