package gohttpclient

// ************************************************************************************************
// ** HTTPClient
// ************************************************************************************************
// ** HTTP client interface
// ************************************************************************************************

type HTTPClient interface {
	SetMethod(string)
	SetURL(string)
	SetPayload(string)
	AddHeader(key string, value string)
	Do() (HTTPResponse, error)
	Reset()

	NewTestServer(func(ResponseWriter, Request)) HTTPTestServer
}

// ************************************************************************************************
// ** HTTPTestServer
// ************************************************************************************************
// ** HTTP Test server
// ************************************************************************************************

type HTTPTestServer interface {
	GetURL() string
	Close()
}
