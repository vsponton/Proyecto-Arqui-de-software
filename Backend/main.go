package main

import (
	"cursos-ucc/app"
	"cursos-ucc/db"
)

func main() {
	db.StartDbEngine()
	app.StartRouter()
}
