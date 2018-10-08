package etcdutil

import (
	"flag"
	"reflect"
	"testing"
)

func TestFlags(t *testing.T) {
	t.Parallel()

	c := NewConfig("foo/")
	fs := flag.NewFlagSet("test", flag.ExitOnError)

	c.AddFlags(fs)

	err := fs.Parse([]string{"-etcd-endpoints=http://12.34.56.78:2379,https://12.34.56.79:2379"})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(c.Endpoints, []string{"http://12.34.56.78:2379", "https://12.34.56.79:2379"}) {
		t.Error("etcd-endpoints was not set:", c.Endpoints)
	}
}
