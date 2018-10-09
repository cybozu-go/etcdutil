package etcdutil

import (
	"flag"
	"io/ioutil"
	"reflect"
	"testing"
)

func testMakeConfig(c *Config) *Config {
	cfg := NewConfig("foo")
	if c == nil {
		return cfg
	}

	if len(c.Endpoints) > 0 {
		cfg.Endpoints = c.Endpoints
	}
	if len(c.Prefix) > 0 {
		cfg.Prefix = c.Prefix
	}
	if len(c.Timeout) > 0 {
		cfg.Timeout = c.Timeout
	}
	if len(c.Username) > 0 {
		cfg.Username = c.Username
	}
	if len(c.Password) > 0 {
		cfg.Password = c.Password
	}
	if len(c.TLSCAFile) > 0 {
		cfg.TLSCAFile = c.TLSCAFile
	}
	if len(c.TLSCertFile) > 0 {
		cfg.TLSCertFile = c.TLSCertFile
	}
	if len(c.TLSKeyFile) > 0 {
		cfg.TLSKeyFile = c.TLSKeyFile
	}
	return cfg
}

func TestFlags(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name     string
		Input    []string
		Expected *Config
		Err      bool
	}{
		{
			"valid endpoints 1",
			[]string{"-etcd-endpoints=http://12.34.56.78:2379"},
			&Config{
				Endpoints: []string{"http://12.34.56.78:2379"},
			},
			false,
		},
		{
			"valid endpoints 2",
			[]string{"-etcd-endpoints=http://12.34.56.78:2379,https://12.34.56.79:2379"},
			&Config{
				Endpoints: []string{"http://12.34.56.78:2379", "https://12.34.56.79:2379"},
			},
			false,
		},
		{
			"invalid endpoints 1",
			[]string{"-etcd-endpoints="},
			nil,
			true,
		},
		{
			"invalid endpoints 2",
			[]string{"-etcd-endpoints=,"},
			nil,
			true,
		},
		{
			"prefix",
			[]string{"-etcd-prefix=t/"},
			&Config{
				Prefix: "t/",
			},
			false,
		},
		{
			"timeout",
			[]string{"-etcd-timeout=1m"},
			&Config{
				Timeout: "1m",
			},
			false,
		},
		{
			"username",
			[]string{"-etcd-username=cybozu"},
			&Config{
				Username: "cybozu",
			},
			false,
		},
		{
			"password",
			[]string{"-etcd-password=passwd"},
			&Config{
				Password: "passwd",
			},
			false,
		},
		{
			"ca",
			[]string{"-etcd-tls-ca=/etc/ca.crt"},
			&Config{
				TLSCAFile: "/etc/ca.crt",
			},
			false,
		},
		{
			"cert",
			[]string{"-etcd-tls-cert=/etc/user.crt"},
			&Config{
				TLSCertFile: "/etc/user.crt",
			},
			false,
		},
		{
			"key",
			[]string{"-etcd-tls-key=/etc/user.key"},
			&Config{
				TLSKeyFile: "/etc/user.key",
			},
			false,
		},
	}

	for _, c := range testCases {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		fs.SetOutput(ioutil.Discard)
		cfg := NewConfig("foo")
		cfg.AddFlags(fs)

		err := fs.Parse(c.Input)
		if err != nil && !c.Err {
			t.Errorf("%s: unexpected error: %v", c.Name, err)
			continue
		}
		if err == nil && c.Err {
			t.Errorf("%s: expected an error", c.Name)
			continue
		}

		expected := testMakeConfig(c.Expected)
		if !reflect.DeepEqual(expected, cfg) {
			t.Errorf("%s: expected=%+v, actual=%+v", c.Name, expected, cfg)
		}
	}
}
