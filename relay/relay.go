package relay

import (
	"encrypt-tunnel/crypto"
	"encrypt-tunnel/tunnel"
	"net"
)

func Relay(client net.Conn, remoteAddr, role, secret, tun string) {
	remote, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		client.Close()
		return
	}

	var src, dst crypto.CipherStream
	if role == "A" {
		src = client
		dst, err = crypto.NewChacha20Stream([]byte(secret), remote)
	} else {
		src, err = crypto.NewChacha20Stream([]byte(secret), client)
		dst = remote
	}

	if err != nil {
		src.Close()
		dst.Close()
		return
	}
	var t tunnel.Tunnel
	if tun == "socks5" {
		t = &tunnel.Socks5Tunnel{}
	}
	t.Forward(src, dst)
}
