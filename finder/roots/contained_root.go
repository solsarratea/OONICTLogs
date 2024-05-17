package roots

import (
	"crypto/x509"
)

func Contained(c *x509.Certificate, roots []*x509.Certificate) bool {
	//FIXME: Obsolete
	for _, cert := range roots {
		if c.Equal(cert) {
			return true
		}
	}
	return false
}

// FindParent performs a linear search for resolving a certificate to a valid root
func FindParent(c *x509.Certificate, roots []*x509.Certificate) *x509.Certificate {
	for _, cert := range roots {
		ans := c.CheckSignatureFrom(cert)
		if ans == nil {
			return cert
		}
	}
	return nil
}
