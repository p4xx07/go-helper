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

func Test_CleanUrl2(t *testing.T) {
	// Create a test URL with unnecessary slashes and a colon after the scheme
	inputUrl, _ := url.Parse("https://www.example.com/://path//to//file")

	// Call the function to clean the URL
	cleanedUrl := CleanUrl(inputUrl)

	// Check that the scheme has been cleaned properly
	if cleanedUrl.Scheme != "https" {
		t.Errorf("Scheme not cleaned properly, expected 'https', got '%s'", cleanedUrl.Scheme)
	}

	// Check that the host has been cleaned properly
	if cleanedUrl.Host != "www.example.com" {
		t.Errorf("Host not cleaned properly, expected 'www.example.com', got '%s'", cleanedUrl.Host)
	}
}
