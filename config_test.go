package etcdutil

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/pelletier/go-toml/v2"
	"sigs.k8s.io/yaml"
)

func testEtcdConfigYAML(t *testing.T) {
	t.Parallel()

	prefix := t.Name() + "/"
	cases := []struct {
		source   string
		expected Config
	}{
		{
			source: `
endpoints:
  - http://etcd1:2379
  - http://etcd2:2379
  - http://etcd3:2379
prefix: /test
timeout: 10s
`,
			expected: Config{
				Endpoints: []string{"http://etcd1:2379", "http://etcd2:2379", "http://etcd3:2379"},
				Prefix:    "/test",
				Timeout:   "10s",
			},
		},
		{
			source: `
username: testuser
password: testpass
`,
			expected: Config{
				Endpoints: DefaultEndpoints,
				Timeout:   DefaultTimeout,
				Prefix:    prefix,
				Username:  "testuser",
				Password:  "testpass",
			},
		},
		{
			source: `
tls-ca-file: ca.crt
tls-cert-file: client.crt
tls-key-file: client.key
`,
			expected: Config{
				Endpoints:   DefaultEndpoints,
				Timeout:     DefaultTimeout,
				Prefix:      prefix,
				TLSCAFile:   "ca.crt",
				TLSCertFile: "client.crt",
				TLSKeyFile:  "client.key",
			},
		},
	}

	for _, c := range cases {
		cfg := NewConfig(prefix)
		err := yaml.Unmarshal([]byte(c.source), cfg)
		if err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(*cfg, c.expected) {
			t.Errorf("%v != %v", *cfg, c.expected)
		}
	}
}

func testEtcdConfigJSON(t *testing.T) {
	t.Parallel()

	prefix := t.Name() + "/"
	cases := []struct {
		source   string
		expected Config
	}{
		{
			source: `
{
    "endpoints": [
        "http://etcd1:2379",
        "http://etcd2:2379",
        "http://etcd3:2379"
    ],
    "prefix": "/test",
    "timeout": "10s"
}
`,
			expected: Config{
				Endpoints: []string{"http://etcd1:2379", "http://etcd2:2379", "http://etcd3:2379"},
				Prefix:    "/test",
				Timeout:   "10s",
			},
		},
		{
			source: `
{
    "username": "testuser",
    "password": "testpass"
}
`,
			expected: Config{
				Endpoints: DefaultEndpoints,
				Timeout:   DefaultTimeout,
				Prefix:    prefix,
				Username:  "testuser",
				Password:  "testpass",
			},
		},
		{
			source: `
{
    "tls-ca-file": "ca.crt",
    "tls-cert-file": "client.crt",
    "tls-key-file": "client.key"
}
`,
			expected: Config{
				Endpoints:   DefaultEndpoints,
				Timeout:     DefaultTimeout,
				Prefix:      prefix,
				TLSCAFile:   "ca.crt",
				TLSCertFile: "client.crt",
				TLSKeyFile:  "client.key",
			},
		},
	}

	for _, c := range cases {
		cfg := NewConfig(prefix)
		err := json.Unmarshal([]byte(c.source), cfg)
		if err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(*cfg, c.expected) {
			t.Errorf("%v != %v", *cfg, c.expected)
		}
	}
}

func testEtcdConfigTOML(t *testing.T) {
	t.Parallel()

	prefix := t.Name() + "/"
	cases := []struct {
		source   string
		expected Config
	}{
		{
			source: `
endpoints = ["http://etcd1:2379", "http://etcd2:2379", "http://etcd3:2379" ]
prefix = "/test"
timeout = "10s"
`,
			expected: Config{
				Endpoints: []string{"http://etcd1:2379", "http://etcd2:2379", "http://etcd3:2379"},
				Prefix:    "/test",
				Timeout:   "10s",
			},
		},
		{
			source: `
username = "testuser"
password = "testpass"
`,
			expected: Config{
				Endpoints: DefaultEndpoints,
				Timeout:   DefaultTimeout,
				Prefix:    prefix,
				Username:  "testuser",
				Password:  "testpass",
			},
		},
		{
			source: `
tls-ca-file = "ca.crt"
tls-cert-file = "client.crt"
tls-key-file = "client.key"
`,
			expected: Config{
				Endpoints:   DefaultEndpoints,
				Timeout:     DefaultTimeout,
				Prefix:      prefix,
				TLSCAFile:   "ca.crt",
				TLSCertFile: "client.crt",
				TLSKeyFile:  "client.key",
			},
		},
	}

	for _, c := range cases {
		cfg := NewConfig(prefix)
		err := toml.Unmarshal([]byte(c.source), cfg)
		if err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(*cfg, c.expected) {
			t.Errorf("%v != %v", *cfg, c.expected)
		}
	}
}

func TestEtcdutilConfig(t *testing.T) {
	t.Run("etcdConfigYAML", testEtcdConfigYAML)
	t.Run("etcdConfigJSON", testEtcdConfigJSON)
	t.Run("etcdConfigTOML", testEtcdConfigTOML)
}
