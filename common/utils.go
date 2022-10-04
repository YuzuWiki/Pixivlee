package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type doFunc func(string, *Query, *Params) (*http.Response, error)

// NewQuery return get params
func NewQuery(values map[string]interface{}) (*url.Values, error) {
	query := url.Values{}
	for k := range values {
		var v string

		switch reflect.TypeOf(values[k]).Kind() {
		case reflect.String:
			v = values[k].(string)
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
			v = fmt.Sprintf("%d", values[k])
		case reflect.Slice:
			_setV, err := json.Marshal(values[k])
			if err != nil {
				return nil, err
			}
			v = string(_setV)
		default:
			return nil, fmt.Errorf(fmt.Sprintf("Query error: unsupported type = %s", reflect.TypeOf(values[k]).String()))
		}

		query.Set(k, v)
	}
	return &query, nil
}

func strip(s string, sep string) string {
	if len(s) == 0 || len(sep) == 0 {
		return s
	}

	s = strings.TrimPrefix(s, sep)
	s = strings.TrimSuffix(s, sep)
	return s
}

// Path returns pixiv api
func Path(paths ...interface{}) string {
	sep := "/"
	elems := []string{"https://" + PixivHost}
	for i := range paths {
		switch reflect.TypeOf(paths[i]).Kind() {
		case reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int:
			elems = append(elems, fmt.Sprintf("%d", paths[i]))
		case reflect.String:
			elems = append(elems, strip(paths[i].(string), sep))
		default:
			log.Printf("Warning: Unsupported path = %v", paths[i])
		}
	}
	return strings.Join(elems, sep)
}

func request(fn doFunc, u string, query *Query, params *Params) ([]byte, error) {
	resp, err := fn(u, query, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Header return http.Header
func Header(fn doFunc, u string, query *Query, params *Params) (header http.Header, err error) {
	resp, err := fn(u, query, params)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return resp.Header, nil
}

// Body return http.body && error
func Body(fn doFunc, u string, query *Query, params *Params) ([]byte, error) {
	return request(fn, u, query, params)
}

// Json return common
func Json(fn doFunc, u string, query *Query, params *Params) ([]byte, error) {
	data, err := request(fn, u, query, params)
	if err != nil {
		return nil, err
	}

	body := &struct {
		Error   bool        `json:"error"`
		Message string      `json:"message"`
		Body    interface{} `json:"body"`
	}{}
	if err = json.Unmarshal(data, body); err != nil {
		return nil, err
	}

	if body.Error {
		return nil, fmt.Errorf(body.Message)
	}

	return json.Marshal(body.Body)
}

func EncodeURL(u string, data *Query) (string, error) {
	URL, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	if data != nil {
		URL.RawQuery = data.Encode()
	}

	return URL.String(), nil
}

func EncodeBody(params *Params) *bytes.Buffer {
	if body, err := json.Marshal(params); err == nil {
		return bytes.NewBuffer(body)
	}
	return nil
}
