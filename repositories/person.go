package repositories

import (
	"api/types"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreatePersonRepository(db *sql.DB, person types.PostPerson) error {
	uuid := uuid.New()

	_, error := db.Query(`INSERT INTO people (id, name, nickname,birthday,stack)
	VALUES ($1, $2,$3,$4, $5)`, uuid, person.Name, person.Nickname, person.Birthday, pq.Array(person.Stack))

	if error != nil {
		return error
	}

	return nil
}

func GetPersonRepository(db *sql.DB, person types.GetPerson) error {
	return nil
}
