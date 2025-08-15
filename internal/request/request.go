package request

import (
	"errors"
	"io"
	"strings"
)

var methodMaps = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PUT":    true,
	"PATCH":  true,
	"DELETE": true,
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

func parseRequestLine(reqBytes []byte) (RequestLine, error) {
	var reqLine RequestLine
	var reqLineBytes []byte

	for i, c := range reqBytes {
		if i+1 < len(reqBytes) && c == '\r' && reqBytes[i+1] == '\n' {
			reqLineBytes = append(reqLineBytes, reqBytes[:i]...)
			break
		}
	}

	reqLineString := strings.Split(string(reqLineBytes), " ")

	if len(reqLineString) != 3 {
		return reqLine, errors.New("start line has invalid parts")
	}

	var method, reqTarget, httpVersion string
	methodPart, reqTargetPart, httpVersionPart := reqLineString[0], reqLineString[1], reqLineString[2]

	_, ok := methodMaps[methodPart]
	if !ok {
		return reqLine, errors.New("method is not supported")
	}
	method = methodPart

	if httpVersionPart != "HTTP/1.1" {
		return reqLine, errors.New("http version is invalid")
	}

	httpVersionParts := strings.Split(httpVersionPart, "/")
	httpVersion = httpVersionParts[1]

	reqTarget = reqTargetPart

	reqLine.HttpVersion = httpVersion
	reqLine.Method = method
	reqLine.RequestTarget = reqTarget

	return reqLine, nil
}

func RequestFromReader(r io.Reader) (*Request, error) {
	req := new(Request)
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	reqLine, err := parseRequestLine(data)
	if err != nil {
		return nil, err
	}

	req.RequestLine = reqLine

	return req, nil

}
