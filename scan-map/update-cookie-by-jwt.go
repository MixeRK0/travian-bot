package main

import "net/http"

func TryToUpdateCookieAfterRequest(r *http.Response) {
	if r == nil {
		return
	}

	if r.Header.Get("Set-Cookie") != "" {
		jwtCookie = r.Header.Get("Set-Cookie")
		cookie = commonCookie + jwtCookie

		println("Cookie updated")
	}
}
