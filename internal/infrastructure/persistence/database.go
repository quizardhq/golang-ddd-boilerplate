package persistence

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/quizardhq/golang-ddd-boilerplate/configs"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/domain/entity"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database is the database type
type Database struct {
	Conn   *sql.DB
	GormDB *gorm.DB
}

// Init intializes a database connection based on provided config
func (DB *Database) Init(config configs.Config) {
	var err error
	connString := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Db.Host, config.Db.Port, config.Db.User, config.Db.Pwd, config.Db.Name)
	DB.Conn, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatalln("Error on create Database driver: ", err.Error())
	}
	DB.GormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: DB.Conn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error on create gorm driver: ", err.Error())
	}

	err = DB.Conn.Ping()
	if err != nil {
		log.Fatalln("Error on connect to Database: ", err.Error())
	}
}

func (db *Database) Migrate() {
	db.GormDB.Migrator().CreateTable(&entity.Person{})
}

// ProvideDatabase provides a database type instance
func ProvideDatabase(config configs.Config) *Database {
	db := &Database{}
	db.Init(config)
	db.Migrate()

	return db
}
