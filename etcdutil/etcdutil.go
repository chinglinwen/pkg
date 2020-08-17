package etcdutil

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	// "go.etcd.io/etcd/clientv3"
	// "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/clientv3"
	"k8s.io/klog"
)

const (
	pkibase  = "/etc/kubernetes/pki/etcd/"
	cafile   = pkibase + "ca.crt"
	certfile = pkibase + "server.crt"
	keyfile  = pkibase + "server.key"
)

// filepath for tls files
type EtcdTls struct {
	CaFile   string
	CertFile string
	KeyFile  string
}

var defaultEtcdTls = EtcdTls{
	CaFile:   cafile,
	CertFile: certfile,
	KeyFile:  keyfile,
}

type Client struct {
	*clientv3.Client
	// usetls bool
}

type Config struct {
	EtcdTls *EtcdTls
}

type options func(c *Config)

func SetTls(etcdtls EtcdTls) options {
	return func(c *Config) {
		c.EtcdTls = &etcdtls
	}
}

// close after use defer cli.Close()
func New(endpoints []string, ops ...options) (*Client, error) {
	c := &Config{EtcdTls: &defaultEtcdTls}
	for _, op := range ops {
		op(c)
	}
	config := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	}
	if usetls(endpoints) {
		tls, err := GetTlsFromFiles(
			c.EtcdTls.CaFile,
			c.EtcdTls.CertFile,
			c.EtcdTls.KeyFile)
		if err != nil {
			return nil, err
		}
		config.TLS = tls
	}
	cli, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}
	klog.Info("create etcd client")
	return &Client{Client: cli}, nil
}

func (c *Client) Put(ctx context.Context, key, val string) error {
	resp, err := c.Client.Put(ctx, key, val)
	if err != nil || resp == nil {
		return err
	}
	// if resp.PrevKv != nil {
	// 	return resp.PrevKv.Value, err
	// }
	return err
}
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	resp, err := c.Client.Get(ctx, key)
	if err != nil {
		// if clientv3.IsConnCanceled(err) {
		// 	// gRPC client connection is closed
		// }
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "", nil
	}
	return string(resp.Kvs[0].Value), nil
}

func usetls(endpoints []string) bool {
	for _, v := range endpoints {
		if strings.HasPrefix(v, "https") {
			return true
		}
	}
	return false
}

// func (c *Client) MemberIDByIP(peerip string) (uint64, error) {
// 	members, err := c.memberList()
// 	if err != nil {
// 		return 0, err
// 	}
// 	for _, v := range members {
// 		peers := v.GetPeerURLs()
// 		if len(peers) == 0 {
// 			klog.Info("got empty peers for etcd member", v.GetName())
// 			continue
// 		}
// 		if matchPeers(peers, peerip, c.usetls) {
// 			return v.GetID(), nil
// 		}
// 	}
// 	return 0, ErrEtcdMemberNotFound
// }

// func matchPeers(peers []string, peerip string, usetls bool) bool {
// 	if len(peerip) == 0 {
// 		return false
// 	}
// 	peer := fmt.Sprintf("http://%v:2380", peerip)
// 	if usetls {
// 		peer = fmt.Sprintf("https://%v:2380", peerip)
// 	}
// 	for _, v := range peers {
// 		if len(v) == 0 {
// 			continue
// 		}
// 		klog.Infof("try match peer: %v, target: %v", v, peer)
// 		if v == peer {
// 			klog.Infof("matched expect peer: %v, got peer: %v", peer, v)
// 			return true
// 		}
// 	}
// 	return false
// }

// func (c *Client) memberList() ([]*etcdserverpb.Member, error) {
// 	klog.Info("try get etcd members")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	resp, err := c.Client.MemberList(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(resp.Members) == 0 {
// 		return nil, fmt.Errorf("member list empty")
// 	}
// 	klog.Info("get etcd members ok")
// 	return resp.Members, nil
// }

// func (c *Client) MemberRemove(peerip string) error {
// 	klog.Infof("try remove etcd member: %v", peerip)
// 	id, err := c.MemberIDByIP(peerip)
// 	if err != nil {
// 		if errors.Is(err, ErrEtcdMemberNotFound) {
// 			klog.Infof("etcd member: %v not found, consider removed", peerip)
// 			return nil
// 		}
// 		return err
// 	}
// 	klog.Infof("got etcd members id: %v, for: %v", id, peerip)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	r, err := c.Client.MemberRemove(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	klog.Infof("etcd member: %v removed", peerip)
// 	pretty("current etcd member list", r.Members)
// 	return nil
// }

func pretty(prefix string, a interface{}) {
	b, _ := json.MarshalIndent(a, "", "  ")
	fmt.Printf("%v: %s\n", prefix, string(b))
}
