build:
	docker build . -t hs/encrypt-tunnel:v1
server:
	docker run -p 8080:8080 \
		-e listenAddr=127.0.0.1:8080 \
		-e remoteAddr=192.168.216.140:8081 \
		-e role=A \
		-e secret=12345678911234567891123456789122 \
		-e tunnel=socks5 \
		-d hs/encrypt-tunnel:v1
client:
	docker run -p 8081:8081 \
		-e listenAddr=127.0.0.1:8081 \
		-e remoteAddr=192.168.200.60:10054 \
		-e role=B \
		-e secret=12345678911234567891123456789122 \
		-e tunnel=socks5 \
		-d hs/encrypt-tunnel:v1