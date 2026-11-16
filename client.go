package etcdutil

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/namespace"
	"google.golang.org/grpc/connectivity"
)

// NewClient creates etcd client.
func NewClient(c *Config) (*clientv3.Client, error) {
	timeout, err := time.ParseDuration(c.Timeout)
	if err != nil {
		return nil, err
	}

	cfg := clientv3.Config{
		Endpoints:            c.Endpoints,
		DialTimeout:          timeout,
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
	if err := blockUntilReady(timeout, client); err != nil {
		client.Close()
		return nil, err
	}
	if c.Prefix != "" {
		client.KV = namespace.NewKV(client.KV, c.Prefix)
		client.Watcher = namespace.NewWatcher(client.Watcher, c.Prefix)
		client.Lease = namespace.NewLease(client.Lease, c.Prefix)
	}

	return client, nil
}

// blockUntilReady waits until the gRPC connection of client becomes ready.
// Since etcd v3.7, clientv3.New does not wait for the connection to be
// established even if grpc.WithBlock() is specified.
// If timeout is not positive, it waits without a deadline as grpc.WithBlock() did.
// See https://github.com/kubernetes/kubernetes/pull/139599
func blockUntilReady(timeout time.Duration, client *clientv3.Client) error {
	ctx, cancel := context.Background(), func() {}
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	}
	defer cancel()

	conn := client.ActiveConnection()
	if conn == nil {
		return errors.New("no grpc connection")
	}

	for {
		currentState := conn.GetState()
		switch currentState {
		case connectivity.Ready:
			// ready, return
			return nil
		case connectivity.Shutdown:
			return errors.New("etcd grpc connection is shut down")
		case connectivity.Idle:
			// attempt connect
			conn.Connect()
		}
		// wait for state change from currentState until context times out
		if !conn.WaitForStateChange(ctx, currentState) {
			return fmt.Errorf("etcd grpc connection not ready: %w", ctx.Err())
		}
	}
}
