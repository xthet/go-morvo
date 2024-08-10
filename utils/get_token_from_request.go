package utils

import "net/http"

func GetTokenFromRequest(r *http.Request) string {
	token_auth := r.Header.Get("Authorization")
	token_query := r.PathValue("token")

	if token_auth != "" {
		return token_auth
	}

	if token_query != "" {
		return token_query
	}

	return ""
}
