package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type IService interface {
	Lock(id string) (*clientv3.LeaseGrantResponse, error)
	LockAndRefresh(id string, done <-chan bool) (bool, error)
	ContainsKey(id string) (bool, error)
}

type service struct {
	etcdClient *clientv3.Client
	options    Options
}

func NewService(cluster []string, options Options) (IService, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   cluster,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return nil, fmt.Errorf("error initializing the etcd client: %v", err)
	}

	return &service{
		etcdClient: c,
		options:    options,
	}, nil
}

func (s *service) Lock(key string) (*clientv3.LeaseGrantResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.options.RefreshInterval)*time.Second)
	defer c(cancel)
	lease, err := s.etcdClient.Lease.Grant(ctx, s.options.Ttl)
	if err != nil {
		return nil, fmt.Errorf("issue while granting a lease: %v", err)
	}

	resp, err := s.etcdClient.Txn(ctx).
		If(clientv3.Compare(clientv3.CreateRevision(key), "=", 0)).
		Then(clientv3.OpPut(key, key, clientv3.WithLease(lease.ID))).
		Commit()

	if err != nil {
		return nil, fmt.Errorf("failed to set key: %v", err)
	}
	if !resp.Succeeded {
		return nil, fmt.Errorf("key already exists, exiting")
	}

	return lease, nil
}

func (s *service) LockAndRefresh(key string, done <-chan bool) (bool, error) {
	lease, err := s.Lock(key)
	if err != nil {
		return false, err
	}

	go s.refresh(lease, done)

	return true, nil
}

func (s *service) refresh(lease *clientv3.LeaseGrantResponse, done <-chan bool) {
	ticker := time.NewTicker(time.Duration(s.options.RefreshInterval) * time.Second)
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	for {
		select {
		case <-done:
			c(cancel)
			return
		case <-ticker.C:
			ctx, cancel = context.WithTimeout(context.Background(), time.Duration(s.options.RefreshInterval)*time.Second)
			_, err := s.etcdClient.KeepAliveOnce(ctx, lease.ID)
			if err != nil {
				c(cancel)
				return
			}
		}
	}
}
func (s *service) ContainsKey(key string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.options.RefreshInterval)*time.Second)
	defer c(cancel)
	resp, err := s.etcdClient.Get(ctx, key)
	if err != nil {
		return false, fmt.Errorf("issue while getting key %v", err)
	}

	return len(resp.Kvs) != 0, nil
}

func c(cancelFunc context.CancelFunc) {
	if cancelFunc != nil {
		cancelFunc()
	}
}
