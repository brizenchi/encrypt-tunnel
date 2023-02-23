package tunnel

import (
	"encrypt-tunnel/crypto"
	"net"
)

type Tunnel interface {
	Forward(client, target crypto.CipherStream)
	Auth(client net.Conn) (err error)
	Connect(client net.Conn) (net.Conn, error)
}
