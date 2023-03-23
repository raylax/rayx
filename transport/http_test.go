package transport

import (
	"context"
	"github.com/raylax/rayx/dsl"
	"github.com/raylax/rayx/expression"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	const url = "https://httpbin.org"
	tests := []struct {
		name    string
		request *dsl.Request
		assert  func(req *dsl.Request, resp *http.Response) bool
	}{
		{
			name: "default method",
			request: &dsl.Request{
				Path: "/get",
			},
			assert: func(req *dsl.Request, resp *http.Response) bool {
				return resp.StatusCode == http.StatusOK
			},
		},
		{
			name: "error method",
			request: &dsl.Request{
				Path:   "/get",
				Method: "get1",
			},
			assert: func(req *dsl.Request, resp *http.Response) bool {
				return resp.StatusCode == http.StatusOK
			},
		},
		{
			name: "get",
			request: &dsl.Request{
				Path:   "/get",
				Method: "get",
			},
			assert: func(req *dsl.Request, resp *http.Response) bool {
				return resp.StatusCode == http.StatusOK
			},
		},
		{
			name: "post body",
			request: &dsl.Request{
				Path:   "/post",
				Method: "POST",
				Body:   `{"a": "rayx-test"}`,
				Headers: dsl.Header{
					"Content-Type": "application/json",
				},
			},
			assert: func(req *dsl.Request, resp *http.Response) bool {
				bytes, _ := ioutil.ReadAll(resp.Body)
				return resp.StatusCode == http.StatusOK && strings.Contains(string(bytes), "rayx-test")
			},
		},
		{
			name: "redirect default",
			request: &dsl.Request{
				Path:   "/redirect-to?url=%2Fget",
				Method: "GEt",
			},
			assert: func(req *dsl.Request, resp *http.Response) bool {
				return resp.StatusCode == http.StatusOK
			},
		},
		{
			name: "redirect",
			request: &dsl.Request{
				Path:            "/redirect-to?url=%2Fget",
				Method:          "GEt",
				FollowRedirects: ref(true),
			},
			assert: func(req *dsl.Request, resp *http.Response) bool {
				return resp.StatusCode == http.StatusOK
			},
		},
		{
			name: "not redirect",
			request: &dsl.Request{
				Path:            "/redirect-to?url=%2Fget",
				Method:          "GEt",
				FollowRedirects: ref(false),
			},
			assert: func(req *dsl.Request, resp *http.Response) bool {
				return resp.StatusCode == http.StatusFound && resp.Header["Location"][0] == "/get"
			},
		},
	}

	env := expression.NewEnvironment()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := NewHttpTransport(env, url, nil, nil)
			defer transport.Close()
			response, err := transport.Do(context.Background(), tt.request)
			if err != nil {
				t.Error(err)
			}
			ok := tt.assert(tt.request, response.(*HttpResponse).raw)
			if !ok {
				t.Error(ok)
			}
		})
	}
}

func TestSetCookie(t *testing.T) {
	const url = "https://httpbin.org"
	env := expression.NewEnvironment()
	transport := NewHttpTransport(env, url, nil, nil)
	defer transport.Close()

	_, err := transport.Do(context.Background(), &dsl.Request{
		Path: "/cookies/set/rayx/rayx111",
	})
	if err != nil {
		t.Error(err)
	}

	resp, err := transport.Do(context.Background(), &dsl.Request{
		Path: "/cookies",
	})
	if err != nil {
		t.Error(err)
	}
	body := resp.(*HttpResponse).raw.Body
	defer body.Close()
	data, _ := ioutil.ReadAll(body)
	if !strings.Contains(string(data), "rayx111") {
		t.Error("set cookie error")
	}
}

func TestGlobalCookies(t *testing.T) {
	const url = "https://httpbin.org"
	env := expression.NewEnvironment()
	transport := NewHttpTransport(env, url, nil, map[string]string{
		"rayx": "rayx111",
	})
	defer transport.Close()

	resp, err := transport.Do(context.Background(), &dsl.Request{
		Path: "/cookies",
	})
	if err != nil {
		t.Error(err)
	}
	body := resp.(*HttpResponse).raw.Body
	defer body.Close()
	data, _ := ioutil.ReadAll(body)
	if !strings.Contains(string(data), "rayx111") {
		t.Error("set cookie error")
	}
}

func TestGlobalHeader(t *testing.T) {
	const url = "https://httpbin.org"
	env := expression.NewEnvironment()
	transport := NewHttpTransport(env, url, map[string]string{
		"x-rayx": "rayx111",
	}, nil)
	defer transport.Close()

	resp, err := transport.Do(context.Background(), &dsl.Request{
		Path: "/headers",
	})
	if err != nil {
		t.Error(err)
	}
	body := resp.(*HttpResponse).raw.Body
	defer body.Close()
	data, _ := ioutil.ReadAll(body)
	if !strings.Contains(string(data), "rayx111") {
		t.Error("set cookie error")
	}
}
