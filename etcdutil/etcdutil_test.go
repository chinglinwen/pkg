package etcdutil

import (
	"context"
	"testing"
	"time"
)

func TestPutAndGet(t *testing.T) {
	// endpoints := []string{"10.249.178.44:2379"}
	endpoints := []string{"localhost:2379"}
	c, err := New(endpoints)
	if err != nil || c == nil {
		t.Error("err", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	k, v := "hello", "there"
	err = c.Put(ctx, k, v)
	if err != nil {
		t.Error("err", err)
		return
	}
	v1, err := c.Get(ctx, k)
	if err != nil {
		t.Error("err", err)
		return
	}
	if string(v1) != v {
		t.Error("returned value not equal", err)
		return
	}
	t.Log("got v1", string(v1))
}

// func TestMemberList(t *testing.T) {
// 	// endpoints := []string{"10.249.178.44:2379"}
// 	endpoints := []string{"localhost:2379"}
// 	c, err := New(endpoints)
// 	if err != nil {
// 		t.Error("err", err)
// 		return
// 	}
// 	m, err := c.memberList()
// 	if err != nil {
// 		t.Error("err", err)
// 		return
// 	}
// 	pretty("list", m)
// }
// func TestMemberIDByIP(t *testing.T) {
// 	// peer := "10.249.178.44"
// 	peer := "localhost"
// 	endpoints := []string{fmt.Sprintf("https://%v:2379", peer)}
// 	c, err := New(endpoints)
// 	if err != nil {
// 		t.Error("err", err)
// 		return
// 	}
// 	id, err := c.MemberIDByIP(peer)
// 	if err != nil {
// 		t.Error("err", err)
// 		return
// 	}
// 	t.Log("got id", id, "for peer", peer)
// }
// func TestMemberRemove(t *testing.T) {
// 	// peer := "10.249.178.44"
// 	peer := "http://localhost:2379"
// 	endpoints := []string{peer}
// 	c, err := New(endpoints)
// 	if err != nil {
// 		t.Error("err", err)
// 		return
// 	}
// 	err = c.MemberRemove("a")
// 	if err != nil {
// 		t.Error("err", err)
// 		return
// 	}
// }

// func TestMatchPeer(t *testing.T) {
// 	cases := []struct {
// 		usetls   bool
// 		endpoint string
// 		ip       string
// 		result   bool
// 	}{
// 		{false, "localhost", "localhost", true},
// 		{false, "localhost", "a", false},
// 		{true, "aaa", "aaa", true},
// 		{true, "aaa", "a", false},
// 		{true, "", "", false},
// 	}
// 	for _, v := range cases {
// 		peers := []string{fmt.Sprintf("http://%v:2380", v.endpoint)}
// 		if v.usetls {
// 			peers = []string{fmt.Sprintf("https://%v:2380", v.endpoint)}

// 		}
// 		if result := matchPeers(peers, v.ip, v.usetls); result != v.result {
// 			t.Errorf("tls: %v, peers: %v, ip: %v, expect: %v, got: %v\n", v.usetls, peers, v.ip, v.result, result)
// 			return
// 		}
// 	}
// }
