package server

type MessageArgs struct{
	User string 
	Message string 
	RoomId int64
}

type ok bool

type IServer interface {
	Enter(username string) string
	EnterRoom()
	ExitRoom()
	SendMessage(message MessageArgs, reply *ok) error
	ListMessage(id int64)
	ListUsers(id int64)
	GetPort() string
}
