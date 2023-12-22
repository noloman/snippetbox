# snippetbox

In order to run, a self-signed certificate must be generated:

1. `mkdir tls`
2. `cd tls`
3. `go run /opt/homebrew/Cellar/go/1.21.5/libexec/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost`
4. run the app with `go run ./cmd/web/`
