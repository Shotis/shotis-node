package crypto

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func SingleCertificatePool(certFile string) (*x509.CertPool, error) {
	pemCert, err := ioutil.ReadFile(certFile)

	if err != nil {
		return nil, err
	}

	pemBlock, _ := pem.Decode(pemCert)
	cert, err := x509.ParseCertificate(pemBlock.Bytes)

	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(cert)
	return certPool, nil
}
