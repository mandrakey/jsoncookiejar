# Go JsonCookieJar
This is a simple implementation of the `net/http/CookieJar` interface which automatically stores and loads known cookies
to and from a JSON file. The main purpose of this is to support stateless CLI applications which access session-based
HTTP endpoints on a "one connect per CLI command", i.e. end immediately after one or some commands have been processed.

The usage is as easy as this:

    import (
        "net/url"
        "net/http"

        "github.com/mandrakey/jsoncookiejar"
    )

    func main() {
        // Create a cookie jar. If the file exists, it is automatically loaded on creation
        jar, err := jsoncookiejar.New("/path/to/my/file.json"); if err != ni; {
            // handle error
        }

        u := url.Parse("https://mydomain.com")
        c := http.Cookie{Name: "some-cookie", Value: "some cookie value"}

        // Add cookies for u and write all cookies to disk immediately (follows the http/CookieJar interface)
        jar.SetCookies(u, []*http.Cookie{c}))

        // Add cookies only in memory
        jar.SetCookiesNoStore(u, []*http.Cookie{c})

        // Write all cookies to disk
        jar.Store()

        // (Re)Load cookies from disk:
        jar.Load()

        // Get loaded cookies from memory (follows the http/CookieJar interface):
        cookies := jar.Cookies(u)
    }
