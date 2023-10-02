package repository

import (
	"database/sql"

	"github.com/PyMarcus/rpc_chat/models"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	Conn *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Conn: db}
}

func (r *Repository) Migrate() error {
	firstQuery := `CREATE TABLE if not exists rooms(id integer primary key autoincrement, name text unique not null);`
	secondQuery := `CREATE TABLE messages (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					user TEXT NOT NULL,
					message TEXT NOT NULL,
					room_id INTEGER NOT NULL
				);`

	_, err := r.Conn.Exec(firstQuery)
	if err != nil {
		return err
	}

	_, err = r.Conn.Exec(secondQuery)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) InsertUser(userMessage models.Message) error {
	query := "INSERT INTO messages (user, message, room_id) VALUES (?, ?, ?);"
	_, err := r.Conn.Exec(query, userMessage.User, userMessage.Message, userMessage.RoomId)

	if err != nil {
		return err
	}
	return nil
}

func (r Repository) InsertRoomsIntoDB(rooms []*models.Room) error {
	for _, room := range rooms {
		_, err := r.Conn.Exec("INSERT INTO rooms (name) VALUES (?);", room.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r Repository) GetAllRooms() ([]*models.Room, error) {
	query := "SELECT * FROM rooms;"
	rows, err := r.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rooms []*models.Room
	for rows.Next() {
		room := &models.Room{}
		_ = rows.Scan(&room.Id, room.Name)
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r Repository) GetRoomById(id int64) (*models.Room, error) {
	var room *models.Room
	query := "SELECT * FROM rooms WHERE id = ?;"
	row := r.Conn.QueryRow(query, id)

	err := row.Scan(&room.Id, &room.Name)

	if err != nil{
		return nil, err 
	}
	return room, nil 
}

func (r Repository) ListAllMessages(roomId int64) ([]*models.Message, error){
	query := "SELECT * FROM messages WHERE room_id = ?;"
	rows, err := r.Conn.Query(query, roomId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		message := &models.Message{}
		_ = rows.Scan(&message.Id, message.User, message.Message, message.RoomId)
		messages = append(messages, message)
	}

	return messages, nil
}
