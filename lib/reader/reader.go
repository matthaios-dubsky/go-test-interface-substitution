package reader

// A simple example on how to use interface subsitution for testing in Go
// Ref:
// https://sendgrid.com/blog/when-writing-unit-tests-dont-use-mocks/

// ReadPublicPemKey reads a byte slice containing the public pem key and returns a pem struct
// Ref: https://golang.org/pkg/encoding/pem/
//
// Function to refactor. The objective is have better instruementation on x509 and pem module
// e.g. mocking, spy, stub etc.
//
//func ReadPublicPemKey(publicPem []byte) (interface{}, error) {
//	block, _ := pem.Decode(publicPem)
//	if block == nil {
//		return nil, ErrPublicKeyDecode
//	}
//
//	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
//	if err != nil {
//		errMsg := "failed to parse PEM/DER encoded public key: " + err.Error()
//		return nil, errors.New(errMsg)
//	}
//
//	return pub, nil
//}

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	// ErrPublicKeyDecode is the expected error return if pem.Decode failed on the public key pem
	ErrPublicKeyDecode = errors.New("failed to parse PEM block containing the public key")
	// ErrPrivateKeyDecode is the expected error return if pem.Decode failed on the private key pem
	ErrPrivateKeyDecode = errors.New("failed to parse PEM block containing the private key")
)

//
// let's define an interface
//

// PemReader interface provides methods for reading public and private pme keys
type PemReader interface {
	ReadPublicKey() (interface{}, error)
	ReadPrivateKey() (interface{}, error)
}

//
// Some short hand function types
//

// Decoder decodes pem in a byte slice
type Decoder func([]byte) (*pem.Block, []byte)

// Parser parse a byte slice into proper x509 struct
type Parser func([]byte) (interface{}, error)

//
// The struct that implements the PemReader interface
//

// Reader implements the PemReader interface
type Reader struct {
	Decoder          Decoder
	PublicKeyParser  Parser
	PrivateKeyParser Parser
}

// ReadPublicKey takes a byte slice and returns public key struct depending on the algoritm
func (r Reader) ReadPublicKey(publicPem []byte) (interface{}, error) {
	block, _ := r.Decoder(publicPem)
	if block == nil {
		return nil, ErrPublicKeyDecode
	}

	publicKey, err := r.PublicKeyParser(block.Bytes)
	if err != nil {
		errMsg := "failed to parse PEM/DER encoded public key: " + err.Error()
		fmt.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	return publicKey, nil
}

// ReadPrivateKey takes a byte slice and returns private key struct depending on the algoritm
func (r Reader) ReadPrivateKey(privatePem []byte) (interface{}, error) {
	block, _ := r.Decoder(privatePem)
	if block == nil {
		return nil, ErrPrivateKeyDecode
	}

	privateKey, err := r.PrivateKeyParser(block.Bytes)
	if err != nil {
		errMsg := "failed to parse PEM/DER encoded private key: " + err.Error()
		fmt.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	return privateKey, nil
}

// New returns a struct that supports the PemReader interface with default internal implemenation
func New() *Reader {
	return &Reader{
		Decoder:          pem.Decode,
		PublicKeyParser:  x509.ParsePKIXPublicKey,
		PrivateKeyParser: x509.ParsePKCS8PrivateKey,
	}
}
