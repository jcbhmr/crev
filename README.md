# crev: Container Registry rEVerse proxy

🐳 Use your vanity domain and reverse proxy to an existing container registry \
💡 Inspired by [ahmetb/serverless-registry-proxy]

<p align=center>
  <img src="https://i.imgur.com/mgC49pV.png">
</p>

<p align=center>
  <a href="https://fizzy-octopus.vercel.app">
    <img alt="Go to website" src="https://img.shields.io/static/v1?style=for-the-badge&message=%E2%86%97%EF%B8%8F+Go+to+demo&color=008FC7&label=">
  </a>
</p>

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
  "rewrites": [{ "source": "/v2(/.*)?", "destination": "/api/crev.go" }],
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
// handles /api/crev
package handler
import "jcbhmr.me/crev"
const ServeHTTP = crev.ServeHTTP
```

## How it works

When you do `docker pull ghcr.io/octocat/my-image`, the `docker` CLI client is
doing a bunch of HTTP requests to the `/v2/...` API routes of the server. What
we can do is just reverse-proxy those requests and edit them just a bit to make
`octocat.me/my-image` successfully proxy to `ghcr.io/octocat/my-image`.

The easy part is mapping things like this:

```sh
https://octocat.me/v2/my-image/manifests/latest >> https://ghcr.io/v2/octocat/my-image/manifests/latest
```

...but then you run into some issues with OAuth tokens and you need some more
complicated logic. When you request credentials for `my-image` you get
`repository:my-image:pull`. Then you try to access `my-image` and get redirected
to `octocat/my-image`, your credentials are wrong. So what can we do to make the
Docker CLI request the right credentials? We can change the return response
`WWW-Authenticate` headers so that they point to our _custom_ token endpoint
(like `/api/crev`) so that we can properly edit `repository:my-image:pull` to
`repository:octocat/my-image:pull` before forwarding it to the _actual_ token
endpoint.

<!-- prettier-ignore -->
Other interesting articles: \
📚 [Hack your own custom domains for Container Registry | Google Cloud Blog](https://cloud.google.com/blog/topics/developers-practitioners/hack-your-own-custom-domains-container-registry) \
📚 [Dodge the next Dockerpocalypse: how to own your own Docker Registry address | HTTP Toolkit](https://httptoolkit.com/blog/docker-image-registry-facade/)

## Development

Why Go and not JavaScript or TypeScript? 'Cause that's what the original project
[ahmetb/serverless-registry-proxy] was written in. 🤷‍♂️ Seems to be working well
so far.

The Vercel site is auto-deployed to [fizzy-octopus.vercel.app] on each push to
a branch in this repo under the `demo/` subfolder. Vercel also automagically
creates preview deployments from any Pull Request branches that are currently
active.

⚠️ The demo project uses the **published version** of the [jcbhmr.me/crev] Go
library instead of the current local copy.

<!-- prettier-ignore-start -->
[ahmetb/serverless-registry-proxy]: https://github.com/ahmetb/serverless-registry-proxy
[devcontainers.community]: https://devcontainers.community/
[devcontainers-community/devcontainers-community.github.io]: https://github.com/devcontainers-community/devcontainers-community.github.io
[webinstall.dev]: https://webinstall.dev/
[fizzy-octopus.vercel.app]: https://fizzy-octopus.vercel.app/
[jcbhmr.me/crev]: https://jcbhmr.me/crev/
<!-- prettier-ignore-end -->
