package etcdutil

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/namespace"
	"google.golang.org/grpc"
)

// NewClient creates etcd client.
func NewClient(c *Config) (*clientv3.Client, error) {
	timeout, err := time.ParseDuration(c.Timeout)
	if err != nil {
		return nil, err
	}

	cfg := clientv3.Config{
		Endpoints:   c.Endpoints,
		DialTimeout: timeout,
		//lint:ignore SA1019 etcd client requires to use grpc.WithBlocks()
		// see https://etcd.io/docs/v3.4/upgrades/upgrade_3_4/#require-grpcwithblock-for-client-dial
		DialOptions:          []grpc.DialOption{grpc.WithBlock()},
		DialKeepAliveTime:    10 * time.Second,
		DialKeepAliveTimeout: 2 * timeout,
		PermitWithoutStream:  true,
		Username:             c.Username,
		Password:             c.Password,
	}

	tlsCfg := &tls.Config{}
	if len(c.TLSCAFile) != 0 || len(c.TLSCA) != 0 {
		var rootCACert []byte
		if len(c.TLSCAFile) != 0 {
			rootCACert, err = os.ReadFile(c.TLSCAFile)
			if err != nil {
				return nil, err
			}
		} else {
			rootCACert = []byte(c.TLSCA)
		}
		rootCAs := x509.NewCertPool()
		ok := rootCAs.AppendCertsFromPEM(rootCACert)
		if !ok {
			return nil, errors.New("failed to parse PEM file")
		}
		tlsCfg.RootCAs = rootCAs
		cfg.TLS = tlsCfg
	}
	if (len(c.TLSCertFile) != 0 && len(c.TLSKeyFile) != 0) || (len(c.TLSCert) != 0 && len(c.TLSKey) != 0) {
		var cert tls.Certificate
		if len(c.TLSCertFile) != 0 && len(c.TLSKeyFile) != 0 {
			cert, err = tls.LoadX509KeyPair(c.TLSCertFile, c.TLSKeyFile)
			if err != nil {
				return nil, err
			}
		} else {
			cert, err = tls.X509KeyPair([]byte(c.TLSCert), []byte(c.TLSKey))
			if err != nil {
				return nil, err
			}
		}
		tlsCfg.Certificates = []tls.Certificate{cert}
		cfg.TLS = tlsCfg
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	if c.Prefix != "" {
		client.KV = namespace.NewKV(client.KV, c.Prefix)
		client.Watcher = namespace.NewWatcher(client.Watcher, c.Prefix)
		client.Lease = namespace.NewLease(client.Lease, c.Prefix)
	}

	return client, nil
}
