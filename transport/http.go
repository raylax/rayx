package transport

import (
	"context"
	"crypto/tls"
	"github.com/raylax/rayx/dsl"
	"github.com/raylax/rayx/expression"
	"github.com/samber/lo"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var httpMethods = []string{
	"OPTIONS",
	"GET",
	"HEAD",
	"POST",
	"PUT",
	"DELETE",
	"TRACE",
	"CONNECT",
}

type HttpResponse struct {
	raw *http.Response
}

type HttpTransport struct {
	env     *expression.Environment
	client  *http.Client
	url     string
	headers map[string]string
	cookies map[string]string
}

func NewHttpTransport(env *expression.Environment, url string, headers map[string]string, cookies map[string]string) *HttpTransport {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client.Jar = &httpCookieJar{
		jar: make(map[string][]*http.Cookie),
	}
	return &HttpTransport{
		env:     env,
		client:  client,
		url:     url,
		headers: headers,
		cookies: cookies,
	}
}

func (h *HttpTransport) Do(ctx context.Context, dsl *dsl.Request) (Response, error) {
	if dsl.Method == "" {
		dsl.Method = "GET"
	} else {
		dsl.Method = strings.ToUpper(dsl.Method)
		_, _, found := lo.FindIndexOf(httpMethods, func(item string) bool {
			return item == dsl.Method
		})
		if !found {
			dsl.Method = "GET"
		}
	}
	var body io.Reader = nil
	if dsl.Body != "" {
		bodyStr, err := h.env.GetString(dsl.Body)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(bodyStr)
	}
	pathStr, err := h.env.GetString(dsl.Path)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(ctx, dsl.Method, h.url+pathStr, body)

	for k, v := range h.headers {
		request.Header.Set(k, v)
	}

	for k, v := range h.cookies {
		cookie, err := request.Cookie(k)
		if err != nil || cookie == nil {
			request.AddCookie(&http.Cookie{
				Name:  k,
				Value: v,
			})
		} else {
			cookie.Value = v
		}
	}

	if err != nil {
		return nil, err
	}
	for key, value := range dsl.Headers {
		headerValue, err := h.env.GetString(value)
		if err != nil {
			return nil, err
		}
		request.Header.Set(key, headerValue)
	}

	if dsl.FollowRedirects == nil || *dsl.FollowRedirects {
		h.client.CheckRedirect = nil
	} else {
		h.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	response, err := h.client.Do(request)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		raw: response,
	}, nil
}

func (h *HttpTransport) Close() {

}

type httpCookieJar struct {
	jar map[string][]*http.Cookie
}

func (p *httpCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	p.jar[u.Host] = cookies
}

func (p *httpCookieJar) Cookies(u *url.URL) []*http.Cookie {
	return p.jar[u.Host]
}
