package etcdutil

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var (
	endpoints = []string{
		"localhost:80",
		"localhost:443",
		"localhost:12359",
	}
)

func listen(t *testing.T) []net.Listener {
	listeners := make([]net.Listener, len(endpoints))
	for i, addr := range endpoints {
		l, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		listeners[i] = l
		go func() {
			c, err := l.Accept()
			if err != nil {
				return
			}
			c.Close()
		}()
	}
	return listeners
}

func TestReorder(t *testing.T) {
	if os.Getuid() != 0 {
		t.Skip("need to run as root")
	}

	listeners := listen(t)
	defer func() {
		for _, l := range listeners {
			l.Close()
		}
	}()

	timeout := time.Millisecond * 100

	t1 := []string{
		"https://localhost:23456",
		"http://localhost",
		"https://localhost",
	}
	reorderEndpoints(t1, timeout)
	if !cmp.Equal(t1, []string{
		"http://localhost",
		"https://localhost:23456",
		"https://localhost",
	}) {
		t.Error("t1 is not reordered correctly", t1)
	}

	t2 := []string{
		"https://localhost:23456",
		"https://localhost",
		"http://localhost",
	}
	reorderEndpoints(t2, timeout)
	if !cmp.Equal(t2, []string{
		"https://localhost",
		"https://localhost:23456",
		"http://localhost",
	}) {
		t.Error("t2 is not reordered correctly", t2)
	}

	t3 := []string{
		"https://localhost",
		"http://localhost",
		"https://localhost:23456",
	}
	reorderEndpoints(t3, timeout)
	if !cmp.Equal(t3, []string{
		"https://localhost",
		"http://localhost",
		"https://localhost:23456",
	}) {
		t.Error("t3 should not be reordered", t3)
	}

	t4 := []string{
		"http://localhost:9998",
		"https://localhost:12359",
		"http://localhost",
	}
	reorderEndpoints(t4, timeout)
	if !cmp.Equal(t4, []string{
		"https://localhost:12359",
		"http://localhost:9998",
		"http://localhost",
	}) {
		t.Error("t4 is not reordered correctly", t4)
	}

	t5 := []string{
		"localhost:13333",
		"localhost:12359",
	}
	reorderEndpoints(t5, timeout)
	if !cmp.Equal(t5, []string{
		"localhost:12359",
		"localhost:13333",
	}) {
		t.Error("t5 is not reordered correctly", t5)
	}

	t6 := []string{
		"localhost:13333",
		"localhost:13334",
	}
	reorderEndpoints(t6, timeout)
	if !cmp.Equal(t6, []string{
		"localhost:13333",
		"localhost:13334",
	}) {
		t.Error("t6 should not be reordered", t6)
	}
}
