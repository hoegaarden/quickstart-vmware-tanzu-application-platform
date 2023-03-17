package pkg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

const (
	aboutOneYear = time.Hour * 24 * 365

	defaultCommonName = "VMware Tanzu Application Platform QuickStart CA"
)

type CA struct {
	CommonName string
}

func (ca CA) Generate() (string, string, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(aboutOneYear)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return "", "", err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: ca.CommonName,
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		IsCA: true,

		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	if template.Subject.CommonName == "" {
		template.Subject.CommonName = defaultCommonName
	}

	certRaw, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return "", "", err
	}

	certBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certRaw})
	if certBytes == nil {
		return "", "", fmt.Errorf("could not encode certificate")
	}

	keyRaw, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return "", "", err
	}

	keyBytes := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyRaw})
	if keyBytes == nil {
		return "", "", fmt.Errorf("could not encode private key")
	}

	return string(certBytes), string(keyBytes), nil
}
