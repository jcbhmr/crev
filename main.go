package main

import (
	"os"
	"net/http"
)

var registryHost string
var registryRepoPrefix string

func init() {
	crev := os.Getenv("CREV")
	dummyURL, err := url.Parse("https://" + crev)
	if err != nil {
		panic(err)
	}

	registryHost = dummyURL.Host

	if dummyURL.Path != "/" {
		registryRepoPrefix = dummyURL.Path
	}
}

var host string
var mux http.ServeMux

func init() {
	mux = http.NewServeMux()
	mux.HandleFunc("/v2/", v2.ServeHTTP)
	mux.HandleFunc("/api/token", token.ServeHTTP)
	mux.HandleFunc("/api/token.go", token.ServeHTTP)
	mux.HandleFunc("/token", token.ServeHTTP)
}

var token = httputil.ReverseProxy{
	Rewrite: tokenRewrite,
}

func tokenRewrite(p *httputil.ProxyRequest) {
	q := p.In.URL.Query()
	var destURL string
	if realm := q.Get("realm"); realm != "" {
		if realmURL, err := url.Parse(realm); err != nil {
			destURL = realmURL.Query().Get("remote")
		}
	}
	if scope := q.Get("scope"); scope != "" {
		scopeParts := strings.Split(scope, ":")
		if len(scopeParts) >= 2 {
			scopeParts[1] = repoPrefix + "/" + scopeParts[1]
		}
		scope = strings.Join(scopeParts, ":")
		q.Set("scope", scope)
	}

	p.Out.URL.Scheme = "https"
	p.Out.URL.Host = registryOAuthURL.Host
	p.Out.URL.Path = registryOAuthURL.Path
	q := p.Out.URL.Query()
	q.Set("realm", "https://ghcr.io/token")
	q.Set("service", p.In.URL.Query().Get("service"))
	q.Set("scope", scope)
	p.Out.URL.RawQuery = q.Encode()
}

func Handler(res http.ResponseWriter, req *http.Request) {
	host = req.Host
	mux.ServeHTTP(res, req)
}
