package user

type noOpUserCredentialsProvider int

// NoOpUserCredentialsProvider returns a user.CredentialsProvider that always returns empty strings
// for both username and password
func NoOpUserCredentialsProvider() CredentialsProvider {
	return new(noOpUserCredentialsProvider)
}

func (cp *noOpUserCredentialsProvider) Get() (Credentials, error) {
	return cp, nil
}

func (_ *noOpUserCredentialsProvider) Username() string {
	return ""
}

func (_ *noOpUserCredentialsProvider) Password() string {
	return ""
}
