package url_helper

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_CleanUrl(t *testing.T) {
	expected := "http://localhost:8080/asdf/asdf//asdf////asdf//asdf////asdf"
	cleanedExpected := "http://localhost:8080/asdf/asdf/asdf/asdf/asdf/asdf"
	url, _ := url.Parse(expected)

	if url.String() != expected {
		panic("not using the right string ")
	}

	if CleanUrl(url).String() != cleanedExpected {
		fmt.Println(url.String())
		panic("different")
	}
}
