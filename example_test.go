package main

import (
	"fmt"
	"github.com/zalando/go-tokens/tokens"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func ExampleTokens() {
	os.Setenv("CREDENTIALS_DIR", "tokens/testdata")
	handler := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		buf, _ := ioutil.ReadAll(req.Body)
		s := string(buf)
		if strings.Contains(s, "scope=foo.read") {
			fmt.Fprint(w, `{"access_token":"test1","token_type":"Bearer",`+
				`"expires_in":4,"scope":"foo.read","realm":"/services"}`)
		} else {
			fmt.Fprint(w, `{"access_token":"test2","token_type":"Bearer",`+
				`"expires_in":8,"scope":"foo.write","realm":"/services"}`)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	url := fmt.Sprintf("http://%s", server.Listener.Addr())
	tokens, err := tokens.Manage(url, []tokens.ManagementRequest{
		tokens.NewPasswordRequest("test1", "foo.read"),
		tokens.NewPasswordRequest("test2", "foo.write"),
	})
	if err != nil {
		log.Fatal(err)
	}

	at, err := tokens.Get("test1")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(at)
	}

	at, err = tokens.Get("test2")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(at)
	}
	// Output:
	// test1 expires in 4 second(s)
	// test2 expires in 8 second(s)
}
