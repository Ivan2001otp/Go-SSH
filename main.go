package main

import (
	"SSH/model"
	"fmt"
	"regexp"
	"strings"

	ssh "github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	sessions map[ssh.Session] *model.Room
	availableRooms []*model.Room
	enterCmd = regexp.MustCompile(`^/enter.*`)
	helpCmd = regexp.MustCompile(`^/help.*`)
	exitCmd = regexp.MustCompile(`^/exit.*`)
	listCmd = regexp.MustCompile(`^/list.*`)
)

func helpMsg() string{
	return `
		Hello and welcome to the chat server! Please 
		use one of the following commands :
			1. /list: To list available rooms
			2. /enter <room>: To enter a room
			3. /exit: To leave the server
			4. /help: To display this message
	`
}

func listRooms() string{
	var sb strings.Builder
	for _,r := range availableRooms{
		sb.WriteString(r.Name + "\n")
	}

	return sb.String();
}

func main(){
	fmt.Println("what's up 1");
}