package service

import (
	"crypto/md5"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type UserService struct {
}

func (s *UserService) GetCurrentUser(request *http.Request) string {
	// get user from request
	ip, _, _ := net.SplitHostPort(request.RemoteAddr)
	forward := request.Header.Get("X-Forwarded-For")

	if forward != "" {
		addrs := strings.Split(forward, ",")
		ip = net.ParseIP(addrs[len(addrs)-1]).String()
	}

	user := fmt.Sprintf("%x", md5.Sum([]byte(ip)))
	return user

}
