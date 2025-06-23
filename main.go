package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"golang.org/x/net/http2"
)

// 测试发送的请求总数
const n = 200

// 测试每秒发送的请求数
const per = 10

var url *string = flag.String("url", "", "要测试的网址")

func main() {
	flag.Parse()
	if *url == "" {
		fmt.Println("请提供要测试的网址")
		flag.PrintDefaults()
		return
	}

	benchHttp2()
	benchHttp3()
}

func benchHttp2() {
	client := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				ClientSessionCache: tls.NewLRUClientSessionCache(100),
				InsecureSkipVerify: true,
			},
		},
	}

	wg := &sync.WaitGroup{}
	i := 1
	count := 1
	var all atomic.Int64
	for range n {
		i++
		if i == per {
			time.Sleep(1 * time.Second)
			i = 0
			fmt.Println("http2", count*per, "次请求已发送")
			count++
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			now := time.Now()
			resp, err := client.Get(*url)
			if err != nil {
				panic(err)
			}
			_, err = io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			all.Add(int64(time.Since(now)))
		}()
	}
	wg.Wait()
	tmp := all.Load()
	fmt.Printf("使用http2平均每个请求用时 %s\n", time.Duration(tmp/n))
}

func benchHttp3() {
	tr := &http3.Transport{
		QUICConfig: &quic.Config{
			Allow0RTT: true,
		},
		TLSClientConfig: &tls.Config{
			ClientSessionCache: tls.NewLRUClientSessionCache(100),
			InsecureSkipVerify: true,
		},
	}

	wg := &sync.WaitGroup{}
	i := 1
	count := 1
	var all atomic.Int64
	for range n {
		i++
		if i == per {
			time.Sleep(1 * time.Second)
			i = 0
			fmt.Println("http3", count*per, "次请求已发送")
			count++
		}
		wg.Add(1)
		go func() {
			req, err := http.NewRequest(http3.MethodGet0RTT, *url, nil)
			if err != nil {
				panic(err)
			}
			defer wg.Done()
			now := time.Now()
			resp, err := tr.RoundTrip(req)
			if err != nil {
				panic(err)
			}
			_, err = io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			all.Add(int64(time.Since(now)))
		}()
	}
	wg.Wait()
	tmp := all.Load()
	fmt.Printf("使用http3平均每个请求用时 %s\n", time.Duration(tmp/n))
}
