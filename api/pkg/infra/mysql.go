package infra

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Db struct {
	client *gorm.DB
}

func Init_mysql() (Db, error) {
	user := os.Getenv("MYSQL_USER")

	password := os.Getenv("MYSQL_PASSWORD")

	host := os.Getenv("MYSQL_HOST")

	port := os.Getenv("MYSQL_PORT")

	database := os.Getenv("MYSQL_DATABASE")

	client, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", user, password, host, port, database))

	return Db{client: client}, err
}

func (db *Db) Find(shell interface{}) *Db {
	db.client = db.client.Find(shell)
	return db
}

func (db *Db) Scan(shell interface{}) *Db {
	db.client = db.client.Scan(shell)
	return db
}
func (db *Db) Where(key string, param ...interface{}) *Db {
	db.client = db.client.Where(key, param...)
	return db
}
func (db *Db) Close() {
	db.client.Close()

}
func (db *Db) Create(shell interface{}) *Db {
	db.client = db.client.Create(shell)
	return db
}
func (db *Db) Delete(value interface{}, where ...interface{}) *Db {

	db.client = db.client.Delete(value, where)
	return db
}
func (db *Db) Having(key string, param ...interface{}) *Db {
	db.client = db.client.Having(key, param)
	return db
}

func (db *Db) Offset(key interface{}) *Db {
	db.client = db.client.Offset(key)
	return db
}

func (db *Db) Limit(key interface{}) *Db {
	db.client = db.client.Limit(key)
	return db
}
func (db *Db) From(raw string) *Db {
	db.client = db.client.Table(raw)
	return db
}

func (db *Db) Join(key string, param ...interface{}) *Db {
	db.client = db.client.Joins(key, param)
	return db
}

func (db *Db) Select(key interface{}, param ...interface{}) *Db {
	db.client = db.client.Select(key, param)
	return db
}
func (db *Db) Rows() *Db {
	db.client.Rows()
	return db
}
func (db *Db) Raw(sql string, param ...interface{}) *Db {
	db.client = db.client.Raw(sql, param)
	return db
}

func (db *Db) Preload(colum string, param ...interface{}) *Db {
	db.client = db.client.Preload(colum, param)
	return db
}

func (db *Db) Exec(sql string, values ...interface{}) *Db {
	db.client = db.client.Exec(sql, values)
	return db
}

func (db *Db) Update(attrs ...interface{}) *Db {
	db.client = db.client.Update(attrs)
	return db
}
func (db *Db) Updates(values interface{}) *Db {
	db.client = db.client.Updates(values)
	return db
}

func (db *Db) Error(values interface{}) error {
	err := db.client.Error
	return err
}
