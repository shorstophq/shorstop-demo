// Package server shows quantum-vulnerable key generation for the ShorStop demo.
package server

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
)

// SigningKey makes an RSA key: broken by Shor's algorithm on a quantum computer.
func SigningKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

// LegacyKey is too small even today.
func LegacyKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}

// SessionKey makes an ECDSA P-256 key: also broken by Shor's algorithm.
func SessionKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

// Fingerprint uses MD5, which is already broken.
func Fingerprint(b []byte) [16]byte { return md5.Sum(b) }

// TLSConfig still allows TLS 1.0.
func TLSConfig() *tls.Config {
	return &tls.Config{MinVersion: tls.VersionTLS10}
}
