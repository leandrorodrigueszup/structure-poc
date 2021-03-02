package repository

import "gorm.io/gorm"

// Database interface
type Database interface {
	Find(dest interface{}, conds ...interface{}) (tx Database)
	Raw(sql string, values ...interface{}) (tx Database)
	Scan(dest interface{}) (tx Database)
	Save(value interface{}) (tx Database)
	Model(value interface{}) (tx Database)
	Where(query interface{}, args ...interface{}) (tx Database)
	First(dest interface{}, conds ...interface{}) (tx Database)
	Updates(values interface{}) (tx Database)
	Delete(value interface{}, conds ...interface{}) (tx Database)
	GetError() error
	GetRowsAffected() int64
	GetStatement() *gorm.Statement
}

// DBWrapper is a wrapper for gorm.DB
type DBWrapper struct {
	*gorm.DB
}

func (db *DBWrapper) Find(dest interface{}, conds ...interface{}) (tx Database) {
	t := db.DB.Find(dest, conds)
	return wrap(t)
}

func (db *DBWrapper) Raw(sql string, values ...interface{}) (tx Database) {
	t := db.DB.Raw(sql, values)
	return wrap(t)
}

func (db *DBWrapper) Scan(dest interface{}) (tx Database) {
	t := db.DB.Scan(dest)
	return wrap(t)
}

func (db *DBWrapper) Save(value interface{}) (tx Database) {
	t := db.DB.Save(value)
	return wrap(t)
}

func (db *DBWrapper) Model(value interface{}) (tx Database) {
	t := db.DB.Model(value)
	return wrap(t)
}

func (db *DBWrapper) Where(query interface{}, args ...interface{}) (tx Database) {
	t := db.DB.Where(query, args)
	return wrap(t)
}

func (db *DBWrapper) First(dest interface{}, conds ...interface{}) (tx Database) {
	t := db.DB.First(dest, conds)
	return wrap(t)
}

func (db *DBWrapper) Updates(values interface{}) (tx Database) {
	t := db.DB.Updates(values)
	return wrap(t)
}

func (db *DBWrapper) Delete(value interface{}, conds ...interface{}) (tx Database) {
	t := db.DB.Delete(value, conds)
	return wrap(t)
}

func wrap(db *gorm.DB) *DBWrapper {
	return &DBWrapper{db}
}

// GetError return the underling transaction error
func (db *DBWrapper) GetError() error {
	return db.Error
}

// GetRowsAffected returns the transaction's rows affected
func (db *DBWrapper) GetRowsAffected() int64 {
	return db.RowsAffected
}

// GetStatement return statement
func (db *DBWrapper) GetStatement() *gorm.Statement {
	return db.Statement
}
