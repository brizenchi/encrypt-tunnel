启动: 
go run main -l :8080 -f 192.168.200.58:8080
-l 监听端口
-f 转发端口


go get github.com/cnlh/benchmark

benchmark -c 10000 -n 1000000 \
-proxy socks5://127.0.0.1:1080 \
http://127.0.0.1:8080/ping