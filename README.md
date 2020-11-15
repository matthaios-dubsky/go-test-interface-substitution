## A simple example on how to use interface substitution for testing in Go

This repo shows one way to refactor a function with dependency for testing without using mocks. This is a simplified example based on the excellent article *When Writing Unit Tests, Don't Use Mocks* (see reference).

The example function, *ReadPublicPemKey* reads a byte slice containing the public pem key and returns a pem struct. Ref: [Package x509](https://golang.org/pkg/encoding/pem/)

The objective here is to have better instruementation on x509 pem module without using mocks or sppy.

```
func ReadPublicPemKey(publicPem []byte) (interface{}, error) {
  block, _ := pem.Decode(publicPem)
  if block == nil {
    return nil, ErrPublicKeyDecode
  }

  pub, err := x509.ParsePKIXPublicKey(block.Bytes)
  if err != nil {
    errMsg := "failed to parse PEM/DER encoded public key: " + err.Error()
    return nil, errors.New(errMsg)
  }

  return pub, nil
}
```

The refactor version is shown in the *lib/reader* folder.

### Reference:
- [When Writing Unit Tests, Don’t Use Mocks](https://sendgrid.com/blog/when-writing-unit-tests-dont-use-mocks/)