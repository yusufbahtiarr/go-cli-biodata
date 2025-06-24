package models

type Profile struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Age       int    `db:"age"`
	Gender    string `db:"gender"`
	CreatedAt string `db:"created_at"`
}
