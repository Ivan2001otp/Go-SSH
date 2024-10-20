package model

import (
	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type Room struct{
	Name string
	History []Message
	Users []User
}

func (r *Room)Enter (sess ssh.Session,term *terminal.Terminal){

	u := User{
		Session:sess,
		Terminal: term,
	}

	r.Users= append(r.Users,u);

	entryMsg := Message{
		From: r.Name,
		Message: "Welcome to the al-df room !",
	}

	send(u,entryMsg)
	for _,m := range r.History{
		send(u,m)
	}
}

func (r *Room) Leave(sess ssh.Session){
	r.Users = removeByUsername(r.Users,sess.User());
}

func (r *Room)SendMessage(from , message string){
	messageObj := Message{From: from,Message:  message}

	r.History = append(r.History, messageObj)
	for _,u := range r.Users{
		if (u.Session.User())!=from{
			send(u,messageObj);
		}
	}
}

func removeByUsername(s []User,n string)[]User{
	var index int

	for i,u := range s{
		if u.Session.User()==n{
			index=i
			break;
		}
	}

	//remove the user from list...
	return append(s[:index],s[index+1:]...)
}

func send(u User,m Message){
	raw := m.From + ">" + m.Message + "\n"
	u.Terminal.Write([]byte(raw));
}