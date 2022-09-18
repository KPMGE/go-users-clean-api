package protocols

type TokenProvider interface {
	Generate(data any) (string, error)
}
