package smtplogin

import (
	"net/smtp"
	"errors"
)

type loginAuth struct {
	identity, username, password string
	host                         string
}

// LoginAuth returns an Auth that implements the LOGIN authentication
// mechanism.
// The returned Auth uses the given username and password to authenticate
// on TLS connections to host and act as identity. Usually identity will be
// left blank to act as username.
func LoginAuth(identity, username, password, host string) smtp.Auth {
	return &loginAuth{identity, username, password, host}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if !server.TLS {
		advertised := false
		for _, mechanism := range server.Auth {
			if mechanism == "LOGIN" {
				advertised = true
				break
			}
		}
		if !advertised {
			return "", nil, errors.New("unencrypted connection")
		}
	}
	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
	}
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	switch string(fromServer[:]) {
		case "Username:":
			resp := []byte(a.username + "\x00")
			return resp, nil

		case "Password:":
			resp := []byte(a.password + "\x00")
			return resp, nil
		default:
			return nil, errors.New(string(fromServer[:]))
	}

}