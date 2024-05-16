package roots

import (
	"crypto/x509"
)

func Contained(c *x509.Certificate, roots []*x509.Certificate) bool {
	for _, cert := range roots {
		if c.Equal(cert) {
			return true
		}
	}
	return false
}
