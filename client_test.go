package etcdutil

import (
	"context"
	"errors"
	"log"
	"os"
	"os/exec"
	"testing"
)

const (
	etcdClientURL = "http://localhost:12379"
	etcdPeerURL   = "http://localhost:12380"
	// nothing listens on this URL.
	unreachableURL = "http://localhost:12381"
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
	client, err := NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	client.Close()
}

func testEtcdClientUnreachable(t *testing.T) {
	cfg := NewConfig(t.Name() + "/")
	cfg.Endpoints = []string{unreachableURL}
	cfg.Timeout = "1s"

	client, err := NewClient(cfg)
	if client != nil {
		defer client.Close()
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline exceeded, got %v", err)
	}
}

func TestEtcdutilClient(t *testing.T) {
	t.Run("etcdClient", testEtcdClient)
	t.Run("etcdClientUnreachable", testEtcdClientUnreachable)
}
