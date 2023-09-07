package crev

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var tokenPathname = os.Getenv("CREV_TOKEN_PATHNAME")
var registryHost = os.Getenv("CREV_REGISTRY_HOST")
var repoPrefix = os.Getenv("CREV_REPO_PREFIX")
var registryOAuth2URL string

func init() {
	u, err := findRegistryOAuth2URL(registryHost)
	if err != nil {
		panic(err)
	}

	registryOAuth2URL = u
}

var realmRegexp = regexp.MustCompile(`realm="(.*?)"`)

func findRegistryOAuth2URL(registryHost string) (string, error) {
	u := fmt.Sprintf("https://%s/v2/", registryHost)
	r, err := http.Get(u)
	if err != nil {
		return "", err
	}

	h := r.Header.Get("WWW-Authenticate")
	if h == "" {
		return "", fmt.Errorf("no WWW-Authenticate header at %s", u)
	}

	matches := realmRegexp.FindStringSubmatch(h)
	if len(matches) == 0 {
		return "", fmt.Errorf("no realm=... in %s", h)
	}

	return matches[1], nil
}

var host string
var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
	mux.HandleFunc("/v2/", v2.ServeHTTP)
	mux.HandleFunc(tokenPathname, token.ServeHTTP)
}

var v2 = httputil.ReverseProxy{
	Rewrite:        v2Rewrite,
	ModifyResponse: v2ModifyResponse,
}

func v2Rewrite(p *httputil.ProxyRequest) {
	p.Out.URL.Host = registryHost
	if repoPrefix != "" {
		if len(p.In.URL.Path) > 4 {
			p.Out.URL.Path = fmt.Sprintf("/v2/%s/%s", repoPrefix, p.In.URL.Path[4:])
		}
	}
}

var realmRegex = regexp.MustCompile(`realm="([^"]+)"`)

func v2ModifyResponse(res *http.Response) error {
	if res.Header.Get("WWW-Authenticate") != "" {
		h := res.Header.Get("WWW-Authenticate")
		h = realmRegex.ReplaceAllString(h, fmt.Sprintf(`realm="https://%s%s"`, host, tokenPathname))
		res.Header.Set("WWW-Authenticate", h)
	}
	return nil
}

var token = httputil.ReverseProxy{
	Rewrite: tokenRewrite,
}

func tokenRewrite(p *httputil.ProxyRequest) {
	scope := p.In.URL.Query().Get("scope")
	if scope == "" {
		panic(fmt.Errorf("?scope not set %s", p.In.URL.String()))
	}
	if repoPrefix != "" {
		scopeParts := strings.Split(scope, ":")
		if len(scopeParts) >= 2 {
			scopeParts[1] = repoPrefix + "/" + scopeParts[1]
		}
		scope = strings.Join(scopeParts, ":")
	}

	u, err := url.Parse(registryOAuth2URL)
	if err != nil {
		panic(err)
	}
	p.Out.URL = u

	q := p.In.URL.Query()
	q.Set("scope", scope)
	p.Out.URL.RawQuery = q.Encode()
}

func ServeHTTP(res http.ResponseWriter, req *http.Request) {
	host = req.Host
	mux.ServeHTTP(res, req)
}