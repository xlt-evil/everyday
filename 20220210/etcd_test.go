package _0220210

/**
 * @Author: hxy
 * @Description:
 * @File:  etcd
 * @Date: 2022/02/10 16:55
 */
import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func Etcd() {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:    []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	// 监听
	go func() {
		rch := cli.Watch(context.Background(),"love")
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}()
	// 租约
	ctx,_ := context.WithTimeout(context.Background(),1 * time.Second)
	lease,err := cli.Create(ctx,5000)
	if err != nil {
		return
	}
	_,err = cli.Put(context.Background(),"love","爱",clientv3.WithLease(clientv3.LeaseID(lease.ID)))
	if err != nil {
		return
	}

	// 设置值
	ctx,_ = context.WithTimeout(context.Background(),1 * time.Second)
	_,err = cli.Put(ctx,"hxy","黄希云")
	if err != nil {
		return
	}
	// 获取值
	resp,err := cli.Get(ctx,"hxy")
	if err != nil {
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}

	// 删除
	_,err = cli.Delete(ctx,"hxy123")
	if err != nil {
		return
	}

	t := time.NewTimer(10 * time.Second)
	select {
	case <-t.C:
		
	}
	return
}


func Test_etcd(t *testing.T) {
	Etcd()
}


func demo() {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:    []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	lease,err := cli.Create(context.TODO(),10)
	if err != nil {
		return
	}
	// 撤销租约
	cli.Revoke(nil,clientv3.LeaseID(lease.ID))
}