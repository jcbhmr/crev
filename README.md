# Your <ins>c</ins>ontainer <ins>r</ins>egistry r<ins>ev</ins>erse proxy

💻 Serverless container registry proxy for your domain \
💡 Inspired by [ahmetb/serverless-registry-proxy]

<p align=center>
  <img src="https://i.imgur.com/dMEK0O2.png">
</p>

🌼 Lets you use **your own domain** for your containers \
📦 Uses an existing container registry as the actual storage \
🔒 Intelligently proxies auth requests to the original auth service \
✅ Works with [docker.io] and [ghcr.io]

## Usage

The core of this project is a Go library that provides a `ServeHTTP()` function and can be hooked up to a serverless provider like Google Cloud Run or Vercel, or hosted as a standalone server. To get started, you can try it locally like this:

```sh
go mod init
go get mycrev.io/crev@v1
```

```go
// main.go
package main
import "mycrev.io/crev"
func main() {
  handler := crev.New("ghcr.io/octocat")
  err := http.ListenAndServe(":8080", handler)
  log.Fatalln(err)
}
```

If you're interested in a completely preconfigured solution, check out the [`examples/`] folder for Vercel, Docker, Google Cloud Run, and more templates. The gist is that you hook up any of those premade serverless things to your domain like `octocatcr.io` or `cr.octocat.me` and live happily ever after. You **can** host your container registry on your main domain. See the examples folder for more details.

[`examples/`]: https://github.com/mycrev/crev/tree/main/examples
[ahmetb/serverless-registry-proxy]: https://github.com/ahmetb/serverless-registry-proxy
