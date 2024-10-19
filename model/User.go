package model

import (
	 "github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type User struct{
	Session ssh.Session
	Terminal *terminal.Terminal;
}