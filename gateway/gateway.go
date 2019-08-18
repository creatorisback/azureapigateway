package gateway

import (
	"net/http"
)

// APIGateway takes two parameter one is handler function for that route and other is whether it is secured or not
func APIGateway(handler http.HandlerFunc, secure bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if route is secured or not.
		if secure {
			r2 := new(http.Request)
			// Copying request header
			*r2 = *r
			client := &http.Client{}
			req, err := http.NewRequest("GET", "https://authenticate1.azurewebsites.net/auth", nil)

			// Passing request header to auth api
			req.Header = r2.Header
			if err != nil {
				println("Auth service is not running")
			}
			resp, err := client.Do(req)

			// Check response status code and either forward or reject request
			if resp.StatusCode == 200 {
				handler.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}

		} else {
			// If route is unsecured execute handler function
			handler.ServeHTTP(w, r)
		}
	})
}
