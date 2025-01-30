package authentication

type Auth interface {
	GetToken(int, string) (string, error)
	GetPayload(string) (int, string, error)
}