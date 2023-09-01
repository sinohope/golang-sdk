package features

type Signer interface {
	Sign(path, timestamp, payload string) (string, error)
	PublicKey() string
	PrivateKey() string
}
