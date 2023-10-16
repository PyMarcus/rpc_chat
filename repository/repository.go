package repository

import (
	"database/sql"
	"log"

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

	pvMessages := `create table if not exists pv(
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						name text not null,
						message text not null,
						from text not null
					);`

	_, err := r.Conn.Exec(firstQuery)
	if err != nil {
		return err
	}

	_, err = r.Conn.Exec(secondQuery)
	if err != nil {
		return err
	}

	_, err = r.Conn.Exec(pvMessages)
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

func (r Repository) InsertPvMessage(pv *models.Pv) error {
	_, err := r.Conn.Exec("INSERT INTO pv (name, message, fromm) VALUES (?, ?, ?);", pv.Name, pv.Message, pv.From);
	if err != nil {
		return err
	}
	
	return nil
}

func (r Repository) GetPvMessage(name string) ([]*models.Pv, error){
	query := "SELECT fromm, message FROM pv where upper(name) = ?;"
	log.Println(query, name)
	rows, err := r.Conn.Query(query, name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pv []*models.Pv
	for rows.Next() {
		text := &models.Pv{}
		_ = rows.Scan(&text.From, &text.Message)
		pv = append(pv, text)
	}

	return pv, nil
}

func (r Repository) DeleteFromMessages(){
	query := "Delete from messages;"
	r.Conn.Exec(query)
}

// GetAllRooms returns room_id, room_name
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
		_ = rows.Scan(&room.Id, &room.Name)
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r Repository) GetRoomById(id int64) (*models.Room, error) {
	var room *models.Room
	query := "SELECT * FROM rooms WHERE id = ?;"
	row := r.Conn.QueryRow(query, id)

	err := row.Scan(&room.Id, &room.Name)

	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r Repository) ListAllMessages(roomId int64) ([]*models.Message, error) {
	query := "SELECT * FROM messages WHERE room_id = ? order by id desc;"
	rows, err := r.Conn.Query(query, roomId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		message := &models.Message{}
		_ = rows.Scan(&message.Id, &message.User, &message.Message, &message.RoomId)
		messages = append(messages, message)
	}

	return messages, nil
}

func (r Repository) ListAllUsers(roomId int64) ([]*models.UsersModel, error) {
	query := "SELECT user FROM messages WHERE room_id = ?;"
	log.Println(query, roomId)
	rows, err := r.Conn.Query(query, roomId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var messages []*models.UsersModel
	for rows.Next() {
		message := &models.UsersModel{}
		_ = rows.Scan(&message.Name)
		messages = append(messages, message)
	}

	return messages, nil
}