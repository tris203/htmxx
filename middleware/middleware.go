package middleware

import (
	"context"
	"crypto/md5"
	"fmt"
	"htmxx/model"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type MiddlewareService struct {
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func GetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get user from session

		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		forward := r.Header.Get("X-Forwarded-For")

		if forward != "" {
			addrs := strings.Split(forward, ",")
			ip = net.ParseIP(addrs[0]).String()
		}

		userobj := model.User{
			Username: fmt.Sprintf("%x", md5.Sum([]byte(ip))),
		}

		//TODO: Return onj in the context

		// set user in context
		ctx := context.WithValue(r.Context(), "user", userobj.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
