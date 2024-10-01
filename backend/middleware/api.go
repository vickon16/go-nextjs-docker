package middleware

import "net/http"

func enableCORS(next http.Handler) http.Handler {}

func jsonContentType(next http.Handler) http.Handler {}
