package jsoncookiejar

import (
    "fmt"
    "os"
    "net/http"
    "net/url"
    "encoding/json"
)

type JsonCookieJar struct {
    filename string
    entries map[string][]*http.Cookie
}

func New(filename string) (*JsonCookieJar, error) {
    jar := &JsonCookieJar{
        filename: filename,
        entries: make(map[string][]*http.Cookie),
    }

    // Try to load from file and silently fail if the file exists
    _, err := os.Stat(jar.filename); if !os.IsNotExist(err) {
        err = jar.Load(); if err != nil {
            return nil, err
        }
    }
    return jar, nil
}

func (jar *JsonCookieJar) Load() error {
    f, err := os.OpenFile(jar.filename, os.O_RDONLY, 0644); if err != nil {
        return fmt.Errorf("failed to open file %s for reading: %s", jar.filename, err)
    }
    defer f.Close()

    decoder := json.NewDecoder(f)
    err = decoder.Decode(&jar.entries); if err != nil {
        return fmt.Errorf("failed to unmarshal cookie store: %s", err)
    }
    return nil
}

func (jar *JsonCookieJar) Store() error {
    f, err := os.OpenFile(jar.filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644); if err != nil {
        return fmt.Errorf("failed to open file for writing: %s", err)
    }
    defer f.Close()

    data, err := json.Marshal(jar.entries); if err != nil {
        return fmt.Errorf("failed to marshal data: %s", err)
    }

    _, err = f.Write(data); if err != nil {
        return fmt.Errorf("faile to write data: %s", err)
    }

    return nil
}

func (jar *JsonCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) error {
    jar.SetCookiesNoStore(u, cookies)
    return jar.Store()
}

func (jar *JsonCookieJar) SetCookiesNoStore(u *url.URL, cookies []*http.Cookie) {
    if cookies == nil {
        delete(jar.entries, u.String())
    } else {
        jar.entries[u.String()] = cookies
    }
}

func (jar *JsonCookieJar) Cookies(u *url.URL) []*http.Cookie {
    return jar.entries[u.String()]
}
