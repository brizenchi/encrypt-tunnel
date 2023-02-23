package crypto

type CipherStream interface {
	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
	Close() error
}
