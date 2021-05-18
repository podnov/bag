package main

import "testing"

func Test_extracturlscheme_http(t *testing.T) {
	givenurl := "http://cryptobag.podnov.com"

	actual := extractUrlScheme(givenurl)

	expected := "http"

	if actual != expected {
		t.Errorf("got %s; want %s", actual, expected)
	}
}

func Test_extracturlscheme_https(t *testing.T) {
	givenurl := "https://cryptobag.podnov.com"

	actual := extractUrlScheme(givenurl)

	expected := "https"

	if actual != expected {
		t.Errorf("got %s; want %s", actual, expected)
	}
}
