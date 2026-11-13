# shorstop-demo

A deliberately vulnerable sample repository for
[ShorStop](https://github.com/shorstophq/shorstop), the post-quantum cryptography scanner.

The code here uses RSA, ECDSA, MD5, old TLS versions and an RSA certificate, so the
[workflow](.github/workflows/shorstop.yml) has something to find. The results are in this
repository's **Security → Code scanning** tab.

The whole setup is the workflow file:

```yaml
name: shorstop
on: [push, pull_request]
permissions:
  contents: read
  security-events: write
jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v7
      - uses: shorstophq/shorstop@v0.1.0
```

Nothing here is used anywhere; the certificate's private key was never saved.
