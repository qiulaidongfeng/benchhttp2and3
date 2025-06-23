# benchhttp2and3
测试同一网址用http2和http3访问的速度

## 原理

先以每秒不超过10次的速度，发送200个http2请求，然后以每秒不超过10次的速度，发送200个http3请求，测量请求发送到响应读取完毕的平均用时。

## 使用方法:

> go install github.com/qiulaidongfeng/benchhttp2and3@latest
> 
> benchhttp2and3 -url 要测试的网址

输出类似

```
http2 10 次请求已发送
http2 20 次请求已发送
http2 30 次请求已发送
http2 40 次请求已发送
http2 50 次请求已发送
http2 60 次请求已发送
http2 70 次请求已发送
http2 80 次请求已发送
http2 90 次请求已发送
http2 100 次请求已发送
http2 110 次请求已发送
http2 120 次请求已发送
http2 130 次请求已发送
http2 140 次请求已发送
http2 150 次请求已发送
http2 160 次请求已发送
http2 170 次请求已发送
http2 180 次请求已发送
http2 190 次请求已发送
http2 200 次请求已发送
使用http2平均每个请求用时 330.588498ms
http3 10 次请求已发送
http3 20 次请求已发送
http3 30 次请求已发送
http3 40 次请求已发送
http3 50 次请求已发送
http3 60 次请求已发送
http3 70 次请求已发送
http3 80 次请求已发送
http3 90 次请求已发送
http3 100 次请求已发送
http3 110 次请求已发送
http3 120 次请求已发送
http3 130 次请求已发送
http3 140 次请求已发送
http3 150 次请求已发送
http3 160 次请求已发送
http3 170 次请求已发送
http3 180 次请求已发送
http3 190 次请求已发送
http3 200 次请求已发送
使用http3平均每个请求用时 302.131589ms
```