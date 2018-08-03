package etcdutil

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/namespace"
)

const (
	// DefaultPrefix is default etcd prefix key
	DefaultPrefix = ""
	// DefaultTimeout is default etcd connection timeout.
	DefaultTimeout = "2s"
)

var (
	// DefaultEndpoints is default etcd servers.
	DefaultEndpoints = []string{"http://localhost:2379"}
)

// Config represents configuration parameters to access etcd.
type Config struct {
	// Endpoints are etcd servers.
	Endpoints []string `yaml:"endpoints"`
	// Prefix is etcd prefix key.
	Prefix string `yaml:"prefix"`
	// Timeout is dial timeout of the etcd client connection.
	Timeout string `yaml:"timeout"`
	// Username is username for loging in to the etcd.
	Username string `yaml:"username"`
	// Password is password for loging in to the etcd.
	Password string `yaml:"password"`
	// TLSCA is root CA path.
	TLSCA string `yaml:"tls-ca"`
	// TLSCert is TLS client certificate path.
	TLSCert string `yaml:"tls-cert"`
	// TLSKey is TLS client private key path.
	TLSKey string `yaml:"tls-key"`
}

// NewConfig creates Config with default values.
func NewConfig() *Config {
	return &Config{
		Endpoints: DefaultEndpoints,
		Prefix:    DefaultPrefix,
		Timeout:   DefaultTimeout,
	}
}

// Client creates etcd client.
func (c *Config) Client() (*clientv3.Client, error) {
	timeout, err := time.ParseDuration(c.Timeout)
	if err != nil {
		return nil, err
	}

	cfg := clientv3.Config{
		Endpoints:   c.Endpoints,
		DialTimeout: timeout,
		Username:    c.Username,
		Password:    c.Password,
	}

	tlsCfg := &tls.Config{}
	if len(c.TLSCA) != 0 {
		rootCACert, err := ioutil.ReadFile(c.TLSCA)
		if err != nil {
			return nil, err
		}
		rootCAs := x509.NewCertPool()
		ok := rootCAs.AppendCertsFromPEM(rootCACert)
		if !ok {
			return nil, errors.New("Failed to parse PEM file")
		}
		tlsCfg.RootCAs = rootCAs
		cfg.TLS = tlsCfg
	}
	if len(c.TLSCert) != 0 && len(c.TLSKey) != 0 {
		cert, err := tls.LoadX509KeyPair(c.TLSCert, c.TLSKey)
		if err != nil {
			return nil, err
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
