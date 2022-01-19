//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"crypto/tls"
	"crypto/x509"
	"path/filepath"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

// Redis TLS file settings
const (
	TlsCertOutputDir    = "/tmp/edgex/secrets/ca"
	CaKeyFileName       = "ca.key"
	CaCertFileName      = "ca.crt"
	RedisKeyFileName    = "redis.key"
	RedisCsrFileName    = "redis.csr"
	RedisCertFileName   = "redis.crt"
	OpensslConfFileName = "openssl.conf"
	RsaKySize           = "4096"
)

// CreateRedisTlsConfig loads TLS certificates from specified path and creates Redis TLS config
func CreateRedisTlsConfig() (tlsConfig *tls.Config, err error) {
	redisKeyFilePath := filepath.Join(TlsCertOutputDir, RedisKeyFileName)
	redisCertFilePath := filepath.Join(TlsCertOutputDir, RedisCertFileName)
	cert, err := tls.LoadX509KeyPair(redisCertFilePath, redisKeyFilePath)
	if err != nil {
		return tlsConfig, errors.NewCommonEdgeX(errors.KindServerError, "fail to parse the Redis TLS key pair", err)
	}
	certificate, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return tlsConfig, errors.NewCommonEdgeX(errors.KindServerError, "fail to parse the certificate", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(certificate)

	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		MinVersion:   tls.VersionTLS12,
		// skip server side SSL verification, primarily for self-signed certs
		InsecureSkipVerify: true, // nolint:gosec
	}
	return tlsConfig, nil
}
