package user

import "testing"

func TestNoOpUserCredentialsProvider(t *testing.T) {
	cp := NoOpUserCredentialsProvider()
	c, err := cp.Get()

	if err != nil {
		t.Fatal("Unexpected error creating the NoOp credentials provider")
	}

	if c.Username() != "" {
		t.Error("Unpexpected username. Wanted an empty string but got %q", c.Username())
	}

	if c.Password() != "" {
		t.Error("Unpexpected password. Wanted an empty string but got %q", c.Password())
	}
}
