package repository

import "github.com/PyMarcus/rpc_chat/models"

type IRepository interface {
	Migrate() error
	InsertUser(userMessage models.Message) error
	GetAllRooms() ([]*models.Room, error)
	GetRoomById(id int64) (*models.Room, error)
	ListAllMessages(roomId int64) ([]*models.Message, error)
}
