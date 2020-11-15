package main

import (
	"fmt"

	"example.com/pemReader/lib/reader"
)

// example
func main() {
	pem := []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCh3ZCmpZgYikZk3bOWVRjpQYkM
qL+wtLbBYPUQ0P21LqbQGSCWnyTPpsQ2IPpZ/UB5XzRwfWG4JM62hFoW1claSBcJ
UEaJLA6uhKCf4J7o6AAvUja9w5el9FXBfrG0as4Y4zASdNGBbGPY6HF1pja3vWEZ
M8m725fkaG1LIwtRYQIDAQAB
-----END PUBLIC KEY-----`)
	r := reader.New()
	pub, err := r.ReadPublicKey(pem)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%v", pub)
}
