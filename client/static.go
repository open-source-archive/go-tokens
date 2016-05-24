package client

type staticClientCredentialsProvider struct {
	id     string
	secret string
}

// NewStaticClientCredentialsProvider returns a client.CredentialsProvider that returns the id and secret
// used in the arguments clientId and clientSecret, respectively
func NewStaticClientCredentialsProvider(clientID string, clientSecret string) CredentialsProvider {
	return &staticClientCredentialsProvider{id: clientID, secret: clientSecret}
}

func (cp *staticClientCredentialsProvider) Get() (Credentials, error) {
	return cp, nil
}

func (cp *staticClientCredentialsProvider) Id() string {
	return cp.id
}

func (cp *staticClientCredentialsProvider) Secret() string {
	return cp.secret
}
