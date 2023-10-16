package models

type Room struct {
	Id 	 int64   `db:"id" json:"id"`
	Name string  `db:"name" json:"name"`
}

type Pv struct{
	Name string  `db:"name" json:"name"`
	Message string  `db:"message" json:"message"`
	From string `db:"from" json:"from"`
}

type Message struct{
	Id		int64	`db:"id" json:"id"`
	User 	string  `db:"user" json:"user"`
	Message string  `db:"message" json:"message"`
	RoomId  int64   `db:"room_id" json:"room_id"`
}

type UsersModel struct{
	Name string `db:"name" json:"name"`
}