package v2

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var splat_registryHost string
var splat_registryScheme string
var splat_repoPrefix string
var splat_tokenURL url.URL

func init() {
	splat_registryScheme = os.Getenv("CREV_REGISTRY_SCHEME")
	if splat_registryScheme == "" {
		splat_registryScheme = "https"
	}

	splat_registryHost = os.Getenv("CREV_REGISTRY_HOST")
	if splat_registryHost == "" {
		splat_registryHost = "registry-1.docker.io"
	}

	splat_repoPrefix = os.Getenv("CREV_REPO_PREFIX")

	baseHRef := os.Getenv("VERCEL_URL")

	baseURL, err := url.Parse(baseHRef)
	if err != nil {
		log.Fatalf("no base URL")
	}

	tokenRelative := os.Getenv("CREV_TOKEN_URL")
	u, err := baseURL.Parse(tokenRelative)
	if err != nil {
		panic(err)
	}

	splat_tokenURL = *u
}

var splat_rp = httputil.ReverseProxy{
	Rewrite:        splat_rewrite,
	ModifyResponse: splat_modifyResponse,
}

func splat_rewrite(p *httputil.ProxyRequest) {
	log.Printf("in method=%s url=%s headers=%s", p.In.URL.String(), p.In.Method, p.In.Header)
	defer func() {
		log.Printf("out method=%s url=%s headers=%s", p.Out.URL.String(), p.Out.Method, p.Out.Header)
	}()

	p.Out.Host = splat_registryHost

	p.Out.URL.Scheme = splat_registryScheme
	p.Out.URL.Host = splat_registryHost

	p.Out.URL.Path = "/v2/"
	p2 := strings.TrimPrefix(p.In.URL.Path, "/v2/")
	if p2 != "" {
		if splat_repoPrefix != "" {
			p.Out.URL.Path += splat_repoPrefix + "/" + p2
		} else {
			p.Out.URL.Path += p2
		}
	}
}

var splat_scopeRepoRegex = regexp.MustCompile(`scope="repository:(.*?):(.*?)"`)

func splat_modifyResponse(r *http.Response) error {
	log.Printf("status=%d url=%s type=%s length=%d", r.StatusCode, r.Request.URL.String(), r.Header.Get("Content-Type"), r.ContentLength)
	log.Printf(`in WWW-Authenticate: %s`, r.Header.Get("WWW-Authenticate"))
	defer func() {
		log.Printf(`out WWW-Authenticate: %s`, r.Header.Get("WWW-Authenticate"))
	}()

	if splat_repoPrefix == "" {
		return nil
	}

	h := r.Header.Get("WWW-Authenticate")
	if h == "" {
		return nil
	}

	h = splat_scopeRepoRegex.ReplaceAllString(h, fmt.Sprintf(`scope="repository:%s/$1:$2"`, splat_repoPrefix))

	r.Header.Set("WWW-Authenticate", h)

	return nil
}

func V2(w http.ResponseWriter, r *http.Request) {
	log.Printf("method=%s url=%s", r.Method, r.URL.String())

	splat_rp.ServeHTTP(w, r)
}
