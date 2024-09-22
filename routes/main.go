package routes

import db "hwc/db"

var (
	Db = db.Db
)

type Error struct {
	Error string `json:"error"`
}
