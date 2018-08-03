package etcdutil

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

const (
	etcdClientURL = "http://localhost:12379"
	etcdPeerURL   = "http://localhost:12380"
)

func testMain(m *testing.M) int {
	circleci := os.Getenv("CIRCLECI") == "true"
	if circleci {
		code := m.Run()
		os.Exit(code)
	}

	etcdPath, err := ioutil.TempDir("", "etcdutil-test")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("etcd",
		"--data-dir", etcdPath,
		"--initial-cluster", "default="+etcdPeerURL,
		"--listen-peer-urls", etcdPeerURL,
		"--initial-advertise-peer-urls", etcdPeerURL,
		"--listen-client-urls", etcdClientURL,
		"--advertise-client-urls", etcdClientURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		cmd.Process.Kill()
		cmd.Wait()
		os.RemoveAll(etcdPath)
	}()

	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testEtcdConfig(t *testing.T) {
	t.Parallel()

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
				Prefix:    DefaultPrefix,
				Timeout:   DefaultTimeout,
				Username:  "testuser",
				Password:  "testpass",
			},
		},
		{
			source: `
tls-ca: ca.crt
tls-cert: client.crt
tls-key: client.key
`,
			expected: Config{
				Endpoints: DefaultEndpoints,
				Prefix:    DefaultPrefix,
				Timeout:   DefaultTimeout,
				TLSCA:     "ca.crt",
				TLSCert:   "client.crt",
				TLSKey:    "client.key",
			},
		},
	}

	for _, c := range cases {
		cfg := NewConfig()
		err := yaml.Unmarshal([]byte(c.source), cfg)
		if err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(*cfg, c.expected) {
			t.Errorf("%v != %v", *cfg, c.expected)
		}
	}
}

func testEtcdClient(t *testing.T) {
	t.Parallel()

	var clientURL string
	circleci := os.Getenv("CIRCLECI") == "true"
	if circleci {
		clientURL = "http://localhost:2379"
	} else {
		clientURL = etcdClientURL
	}

	cfg := NewConfig()
	cfg.Prefix = t.Name() + "/"
	cfg.Endpoints = []string{clientURL}
	_, err := cfg.Client()
	if err != nil {
		t.Fatal(err)
	}
}

func TestEtcdutil(t *testing.T) {
	t.Run("etcdConfig", testEtcdConfig)
	t.Run("etcdClient", testEtcdClient)
}
