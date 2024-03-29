package etcdutil

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

const (
	etcdClientURL = "http://localhost:12379"
	etcdPeerURL   = "http://localhost:12380"
)

func testMain(m *testing.M) int {
	etcdPath, err := os.MkdirTemp("", "etcdutil-test")
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

func testEtcdClient(t *testing.T) {
	t.Parallel()

	cfg := NewConfig(t.Name() + "/")
	cfg.Endpoints = []string{etcdClientURL}
	_, err := NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEtcdutilClient(t *testing.T) {
	t.Run("etcdClient", testEtcdClient)
}
