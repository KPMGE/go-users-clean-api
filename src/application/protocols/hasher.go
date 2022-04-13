package protocols

type Hasher interface {
	Hash(plainText string) string
}
