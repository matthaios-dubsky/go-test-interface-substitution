package reader

var testTablePublicKeyReader = []struct {
	name          string
	pubPem        string
	expectedError error
}{
	{
		"Happy Path",
		`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCh3ZCmpZgYikZk3bOWVRjpQYkM
qL+wtLbBYPUQ0P21LqbQGSCWnyTPpsQ2IPpZ/UB5XzRwfWG4JM62hFoW1claSBcJ
UEaJLA6uhKCf4J7o6AAvUja9w5el9FXBfrG0as4Y4zASdNGBbGPY6HF1pja3vWEZ
M8m725fkaG1LIwtRYQIDAQAB
-----END PUBLIC KEY-----		
`,
		nil,
	},
	{
		"Bad Path",
		"some non sense",
		ErrPublicKeyDecode,
	},
}
