package user

import (
	"encoding/json"
	"io/ioutil"
)

type jsonFileUserCredentialsProvider struct {
	path string
}

// NewJSONFileUserCredentialsProvider returns a user.CredentialsProvider that reads both username and
// password from a JSON file stored in the specified filesystem path.
// The contents of such file should follow the following specifications:
//		{"application_username":"foo","application_password":"bar"}
func NewJSONFileUserCredentialsProvider(path string) CredentialsProvider {
	return &jsonFileUserCredentialsProvider{path}
}

func (cp *jsonFileUserCredentialsProvider) Get() (Credentials, error) {
	buf, err := ioutil.ReadFile(cp.path)
	if err != nil {
		return nil, err
	}

	var credentials jsonFileUserCredentials
	err = json.Unmarshal(buf, &credentials)
	if err != nil {
		return nil, err
	}
	return credentials, nil
}

type jsonFileUserCredentials struct {
	User string `json:"application_username"`
	Pass string `json:"application_password"`
}

func (uc jsonFileUserCredentials) Username() string {
	return uc.User
}

func (uc jsonFileUserCredentials) Password() string {
	return uc.Pass
}
