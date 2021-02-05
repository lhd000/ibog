package model

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"
	"os"
	"sync"

	"iblog/config"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type dbconn struct {
	DB  *gorm.DB
	SDB *sql.DB
}

var dbinst *dbconn
var once sync.Once
var DB *gorm.DB
var LOCK sync.Mutex

func GetDBInstance() *dbconn {

	d := initDb()
	dbinst = &dbconn{DB: d, SDB: d.DB()}

	return dbinst
}

func init() {
	log.Println("init the db start")

	once.Do(func() {
		GetDBInstance()

		log.Println("init db get instance end ")
		DB = dbinst.DB
	})
}

func initDb() *gorm.DB {

	log.Println("init the sql db")

	// Openning file
	if _, err := os.Stat(config.BLOG_DATA); os.IsNotExist(err) {
		// path/to/whatever does not exist
	}
	db, err := gorm.Open("sqlite3", config.BLOG_DATA)

	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}

	// Creating the table

	if !db.HasTable(&Admin{}) {
		db.CreateTable(&Admin{})

		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Admin{})
	}

	if !db.HasTable(&Category{}) {

		db.CreateTable(&Category{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Category{})
	}

	//CreateDB([]interface{}{Article{}}, db)

	if !db.HasTable(&Article{}) {

		db.CreateTable(&Article{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Article{})
	}

	return db
}

func CreateDB(tables []interface{}, db *gorm.DB) {
	for _, v := range tables {
		if !db.HasTable(&v) {
			db.CreateTable(&v)
			db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&v)
		}
	}
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
