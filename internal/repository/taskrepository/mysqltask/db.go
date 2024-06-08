package mysqltask

import (
	"go-backend-clean-arch/internal/domain"
	"go-backend-clean-arch/internal/infrastructure/persistance/db/mysql"
	"log"
)

type DB struct {
	conn mysql.DB
}

func New(conn mysql.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (d *DB) Create(t domain.Task) (domain.Task, error) {
	panic("mysqltask -> create : IMPL ME")
}

func (d *DB) List() {
	log.Print("mysql-taskgateway -> Find - IMPL ME")
}
