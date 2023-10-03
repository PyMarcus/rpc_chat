package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"

	"github.com/PyMarcus/rpc_chat/connection"
	"github.com/PyMarcus/rpc_chat/models"
	"github.com/PyMarcus/rpc_chat/repository"
)

type MyService struct{}

type Args struct {
	Id int64
}

type ReplyOk struct{}

var wg sync.WaitGroup

type ReplyListUsers struct {
	Data []*models.UsersModel
}

type ReplyListMessages struct {
	Data []*models.Message
}

func (s MyService) Enter(username string) error {
	return nil
}

func (s MyService) EnterRoom() {}

func (s MyService) ExitRoom() {}

func (s *MyService) SendMessage(message MessageArgs, reply *ReplyOk) error{
	log.Println("Calling send message")
	conn, _ := connection.ConnectionSQL()
	data := repository.NewRepository(conn)
	user := models.Message{Id:1, User: message.User,
		 Message: message.Message,
		  RoomId: message.RoomId}
	data.InsertUser(user)
	return nil
}

func (s MyService) ListMessage(id int64, reply *ReplyListMessages) error {
	conn, _ := connection.ConnectionSQL()
	data := repository.NewRepository(conn)
	message, _ := data.ListAllMessages(id)
	reply.Data = message
	return nil
}

func (s *MyService) ListUsers(id int64, reply *ReplyListUsers) error {
	conn, _ := connection.ConnectionSQL()
	data := repository.NewRepository(conn)
	users, _ := data.ListAllUsers(id)
	reply.Data = users
	return nil
}

// Run start server
func RunServer() {
	rpc.Register(new(MyService))
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":12345")

	if err != nil {
		panic(err)
	}
	log.Println("Listening on port: ", ":12345")

	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go http.Serve(listener, nil)
	wg.Wait()

}
