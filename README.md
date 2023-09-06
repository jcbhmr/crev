<meta name="go-import" content="jcbhmr.me/crev git https://github.com/jcbhmr/crev">
<meta name="go-source" content="jcbhmr.me/crev _ https://github.com/jcbhmr/crev/tree/main{/dir} https://github.com/jcbhmr/crev/blob/main{/dir}/{file}#L{line}">

# crev: Container Registry rEVerse proxy

🐳 Use your vanity domain and reverse proxy to an existing container registry

<p align=center>
  <img src="">
</p>

## Installation

```sh
go install jcbhmr.me/crev
```

## Usage

[devcontainers.community] is a great example of this project in the wild. Check
out [devcontainers-community/devcontainers-community.github.io] to see how it's
set up with Vercel!

```sh
docker run --rm -it devcontainers.community/cpp
oras manifest fetch devcontainers.community/features/deno
oras manifest fetch devcontainers.community/templates/dart
```

### Vercel

```jsonc
// vercel.json
{
  "rewrites": [
    { "source": "/v2(/.*)?", "destination": "/api/crev.go" },
    { "source": "/api/crev-token", "destination": "/api/crev.go" }
  ],
  "env": {
    "CREV_TOKEN_URL": "/api/crev-token",
    "CREV_REGISTRY_HOST": "ghcr.io",
    "CREV_REPO_PREFIX": "jcbhmr"
  }
}
```

```go
// api/crev.go
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
<!-- prettier-ignore-end -->
