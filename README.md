# gofollower

[![GoReport](https://goreportcard.com/badge/github.com/jksolbakken/gofollower)](https://goreportcard.com/report/github.com/jksolbakken/gofollower)

Native version of [linkfollower](https://www.npmjs.com/package/linkfollower), no JS runtime required.

### Verify the integrity of the release

The `checksums.txt` is signed using [Sigstore cosign](https://github.com/sigstore/cosign). To verify it, run:

```bash
cosign verify-blob \\
  --bundle checksums.txt.sigstore.json \\
  --certificate-identity="https://github.com/jksolbakken/gofollower/.github/workflows/main.yaml@refs/heads/main" \\
  --certificate-oidc-issuer="https://token.actions.githubusercontent.com" \\
  checksums.txt
```

Then verify the checksum of your binary of choice.

