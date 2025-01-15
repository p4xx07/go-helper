package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type IService interface {
	Lock(id string, refreshInterval time.Duration) (*clientv3.LeaseGrantResponse, error)
	LockAndRefresh(ctx context.Context, id string, refreshInterval time.Duration, done <-chan bool) error
	ContainsKey(id string) (bool, error)
}

type service struct {
	etcdClient *clientv3.Client
}

func NewService(cluster []string) (IService, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   cluster,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return nil, fmt.Errorf("error initializing the etcd client: %v", err)
	}

	return &service{
		etcdClient: c,
	}, nil
}

func (s *service) Lock(key string, refreshDuration time.Duration) (*clientv3.LeaseGrantResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), refreshDuration)
	defer c(cancel)
	lease, err := s.etcdClient.Lease.Grant(ctx, int64(refreshDuration.Seconds()))
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

func (s *service) LockAndRefresh(ctx context.Context, key string, duration time.Duration, done <-chan bool) error {
	lease, err := s.Lock(key, duration)
	if err != nil {
		return err
	}

	go s.refresh(ctx, lease, duration, done)

	return nil
}

func (s *service) refresh(ctx context.Context, lease *clientv3.LeaseGrantResponse, duration time.Duration, done <-chan bool) {
	ticker := time.NewTicker(duration)
	var (
		cancel context.CancelFunc
	)
	for {
		select {
		case <-done:
			c(cancel)
			return
		case <-ticker.C:
			ctx, cancel = context.WithTimeout(context.Background(), duration)
			_, err := s.etcdClient.KeepAliveOnce(ctx, lease.ID)
			if err != nil {
				c(cancel)
				return
			}
		}
	}
}

func (s *service) ContainsKey(key string) (bool, error) {
	resp, err := s.etcdClient.Get(context.Background(), key)
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
