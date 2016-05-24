package user

type staticUserCredentialsProvider struct {
	username string
	password string
}

// NewStaticUserCredentialsProvider returns a user.CredentialsProvider that returns the username and password
// used in the arguments u and p, respectively
func NewStaticUserCredentialsProvider(u string, p string) CredentialsProvider {
	return &staticUserCredentialsProvider{username: u, password: p}
}

func (cp *staticUserCredentialsProvider) Get() (Credentials, error) {
	return cp, nil
}

func (cp *staticUserCredentialsProvider) Username() string {
	return cp.username
}

func (cp *staticUserCredentialsProvider) Password() string {
	return cp.password
}
