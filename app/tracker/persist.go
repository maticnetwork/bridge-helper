package tracker

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Tabler - To be implemented by all DB models, so that we can override
// default behaviour ( i.e. table name setting strategy ) of `gorm`
type Tabler interface {
	TableName() string
}

// RootChain - Tx performed on root chain during deposit/ withdraw, to be
// persisted in this table, along with respective status
type RootChain struct {
	TransactionHash string `gorm:"column:txhash;type:char(66);primaryKey"`
	Code            int    `gorm:"column:code;type:smallint;not null"`
	Message         string `gorm:"column:msg;type:varchar;not null"`
}

// TableName - Overriding default table name
func (RootChain) TableName() string {
	return "root_chain"
}

// ChildChain - Tx performed on child chain during deposit/ withdraw, to be
// persisted in this table, along with respective status
type ChildChain struct {
	TransactionHash string `gorm:"column:txhash;type:char(66);primaryKey"`
	Code            int    `gorm:"column:code;type:smallint;not null"`
	Message         string `gorm:"column:msg;type:varchar;not null"`
}

// TableName - Overriding default table name
func (ChildChain) TableName() string {
	return "child_chain"
}

// Connecting to postgres database
func connectToDB() *gorm.DB {
	dbPort, err := strconv.Atoi(get("DB_PORT"))
	if err != nil {
		log.Fatalln("[!] ", err)
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", get("DB_USER"), get("DB_PASSWORD"), get("DB_HOST"), dbPort, get("DB_NAME"))),
		&gorm.Config{})
	if err != nil {
		log.Fatalln("[!] ", err)
	}

	return db
}
