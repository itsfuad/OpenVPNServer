package encryption

import (
    "crypto/tls"
    "crypto/x509"
	"os"
    "vpn-server/config"
)

var tlsConfig *tls.Config

func Initialize(cfg *config.Config) error {
    cert, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
    if err != nil {
        return err
    }

    caCert, err := os.ReadFile(cfg.CAFile)
    if err != nil {
        return err
    }

    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    tlsConfig = &tls.Config{
        Certificates: []tls.Certificate{cert},
        ClientCAs:    caCertPool,
        ClientAuth:   tls.RequireAndVerifyClientCert,
    }

    return nil
}

func GetTLSConfig() *tls.Config {
    return tlsConfig
}
