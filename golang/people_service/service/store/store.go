package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Store struct {
	conn *pgx.Conn
}

type People struct {
	ID   int32
	Name string
}

// NewStore creates new database connection
func NewStore(connString string) *Store {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	// make migration

	return &Store{
		conn: conn,
	}
}

func (s *Store) ListPeople() ([]People, error) {

	people := make([]People, 0, 0)

	rows, err := s.conn.Query(context.Background(), "select * from people")
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("error while reding: %d", err)
		}

		people = append(people, People{
			ID:   values[0].(int32),
			Name: values[1].(string),
		})
	}

	return people, nil
}

func (s *Store) GetPeopleByID(id string) (People, error) {

	var name string
	var _id int32

	row := s.conn.QueryRow(context.Background(), "select * from people where id = "+id)

	err := row.Scan(&_id, &name)
	if err != nil {
		return People{}, err
	}
	return People{ID: _id, Name: name}, nil
}