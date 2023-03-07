//
// Copyright (C) 2022-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

// TLS file settings
const (
	// General settings
	BaseOutputDir       = "/tmp/edgex/secrets"
	CaKeyFileName       = "ca.key"
	CaCertFileName      = "ca.crt"
	OpensslConfFileName = "openssl.conf"
	RsaKySize           = "4096"

	// Redis specific settings
	RedisTlsCertOutputDir = BaseOutputDir + "/redis"
	RedisTlsSecretName    = "redis-tls"
	RedisKeyFileName      = "redis.key"
	RedisCsrFileName      = "redis.csr"
	RedisCertFileName     = "redis.crt"

	// MQTT specific settings
	MqttClientKeyFileName  = "mqtt.key"
	MqttClientCertFileName = "mqtt.crt"
	EnvMessageBusMqttTls   = "EDGEXPERT_MESSAGEBUS_MQTT_TLS"
)

// CreateRedisTlsConfigFromPEM loads TLS certificates from PEM encoded data and creates Redis TLS config
func CreateRedisTlsConfigFromPEM(certPEMBlock, keyPEMBlock []byte) (*tls.Config, errors.EdgeX) {
	var tlsConfig *tls.Config
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return tlsConfig, errors.NewCommonEdgeX(errors.KindServerError, "fail to parse the Redis TLS key pair", err)
	}
	certificate, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return tlsConfig, errors.NewCommonEdgeX(errors.KindServerError, "fail to parse the Redis TLS certificate", err)
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
