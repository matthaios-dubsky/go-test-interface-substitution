package reader

import (
	"crypto/rsa"
	"encoding/pem"
	"testing"
)

// a basic table drive test
func TestPublicKeyReader(t *testing.T) {
	for _, test := range testTablePublicKeyReader {
		t.Run(test.name, func(t *testing.T) {
			r := New()
			actual, err := r.ReadPublicKey([]byte(test.pubPem))
			if err != test.expectedError {
				t.Errorf("ReadPublicPemKey should not returns an error: %v", err.Error())
				return
			}
			// the pub key return should be rsa type
			switch actual.(type) {
			case *rsa.PublicKey:
				return
			default:
				if err != nil {
					return
				}
				t.Error("ReadPublicPemKey should return a RSA type public key")
			}
		})
	}
}

// test based on interface subsitution
func TestPublicKeyDecoder(t *testing.T) {
	r := New()
	t.Run("replace a decoder stub that always returns an error", func(t *testing.T) {
		errDecoder := func([]byte) (*pem.Block, []byte) {
			return nil, nil
		}
		r.Decoder = errDecoder
		// This is the happy path pubPem
		_, err := r.ReadPublicKey([]byte(testTablePublicKeyReader[0].pubPem))
		if err != ErrPublicKeyDecode {
			t.Errorf("ReadPublicPemKey should returns an error: %v", err.Error())
			return
		}
	})
}
