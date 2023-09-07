<meta name="go-import" content="jcbhmr.me/crev git https://github.com/jcbhmr/crev">
<meta name="go-source" content="jcbhmr.me/crev _ https://github.com/jcbhmr/crev/tree/main{/dir} https://github.com/jcbhmr/crev/blob/main{/dir}/{file}#L{line}">

# crev: Container Registry rEVerse proxy

🐳 Use your vanity domain and reverse proxy to an existing container registry \
💡 Inspired by [ahmetb/serverless-registry-proxy]

<p align=center>
  <img src="https://i.imgur.com/mgC49pV.png">
</p>

**Try it yourself! 🚀**

```sh
# Use the Docker Registry HTTP API
curl -fsSL https://jcbhmr.me/v2/crev-demo/manifests/latest
# OR just use Docker directly
docker run --rm -it jcbhmr.me/crev-demo
```

## Installation

```sh
go install jcbhmr.me/crev
```

<details><summary>🚄 Don't have <code>go</code> installed?</summary>

Use [webinstall.dev] to express install the Go toolchain!

```sh
curl -sS https://webi.sh/golang | sh
```

</details>

## Usage

[devcontainers.community] is a great example of this project in the wild. Check
out [devcontainers-community/devcontainers-community.github.io] to see how it's
set up with Vercel!

```sh
docker run --rm -it devcontainers.community/cpp
# These are non-runnable metadata images
oras manifest fetch devcontainers.community/features/deno
oras manifest fetch devcontainers.community/templates/dart
```

### Vercel

```jsonc
// vercel.json
{
  "rewrites": [
    { "source": "/v2(/.*)?", "destination": "/api/crev.go" }
  ],
  "env": {
    "CREV_TOKEN_URL": "/api/crev",
    "CREV_REGISTRY_HOST": "ghcr.io",
    "CREV_REPO_PREFIX": "jcbhmr"
  }
}
```

```go
// api/crev.go
// handles /v2(/.*)?
// handles /api/crev(.go)?
package handler
import "jcbhmr.me/crev"
const ServeHTTP = crev.ServeHTTP
```

## Development

Why Go and not JavaScript or TypeScript? 'Cause that's what the original project
[ahmetb/serverless-registry-proxy] was written in. 🤷‍♂️ Seems to be working well
so far.

<!-- prettier-ignore-start -->
[ahmetb/serverless-registry-proxy]: https://github.com/ahmetb/serverless-registry-proxy
[devcontainers.community]: https://devcontainers.community/
[devcontainers-community/devcontainers-community.github.io]: https://github.com/devcontainers-community/devcontainers-community.github.io
[webinstall.dev]: https://webinstall.dev/
<!-- prettier-ignore-end -->
