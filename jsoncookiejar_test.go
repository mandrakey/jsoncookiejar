package jsoncookiejar

import (
	"testing"
	"fmt"
	"net/url"
	"os"
	"net/http"
)

const (
	cookieJarFileName = "cookies.json"
	cookieDomain = "https://mydomain.com"
)

func TestSetCookie(t *testing.T) {
	jar, err := getJar(); if err != nil {
		t.Fatalf("Failed to open cookie jar: %s", err)
	}

	u, err := url.Parse("https://mydomain.com"); if err != nil {
		t.Fatalf("Failed to parse URL: %s", err)
	}
	c := &http.Cookie{
		Name: "mycookie",
		Value: "blabla",
		MaxAge: 0,
		Secure: false,
		HttpOnly: false,
	}

	err = jar.SetCookies(u, []*http.Cookie{c}); if err != nil {
		t.Fatalf("Failed to store cookies in the jar file")
	}
}

func TestReadCookies(t *testing.T) {
	jar, err := getJar(); if err != nil {
		t.Fatalf("Failed to open cookie jar: %s", err)
	}

	u, err := url.Parse(cookieDomain); if err != nil {
		t.Fatalf("Failed to parse URL: %s", err)
	}

	cookies := jar.Cookies(u)
	cookieCount := len(cookies)
	if cookieCount != 1 {
		t.Errorf("Expect 1 cookies, got %d", cookieCount)
	}
}

func TestUnsetCookies(t *testing.T) {
	jar, err := getJar(); if err != nil {
		t.Fatalf("Failed to open cookie jar: %s", err)
	}

	u, err := url.Parse(cookieDomain); if err != nil {
		t.Fatalf("Failed to parse URL: %s", err)
	}

	cookies := jar.Cookies(u)
	cookieCount := len(cookies)
	if cookieCount != 1 {
		t.Errorf("Expect 1 cookies, got %d", cookieCount)
	}

	jar.SetCookies(u, nil)
	cookies = jar.Cookies(u)
	cookieCount = len(cookies)
	if cookieCount != 0 {
		t.Errorf("Expect 0 cookies, got %d", cookieCount)
	}
}

func getJar() (*JsonCookieJar, error) {
	jar, err := New(cookieJarFile()); if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}
	return jar, nil
}

func cookieJarFile() string {
	return fmt.Sprintf("%s%s", os.TempDir(), cookieJarFileName)
}
