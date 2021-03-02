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

// Find find records that match given conditions
func (db *DBWrapper) Find(dest interface{}, conds ...interface{}) (tx Database) {
	t := db.DB.Find(dest, conds...)
	return wrap(t)
}

// Raw executes a raw SQL statement
func (db *DBWrapper) Raw(sql string, values ...interface{}) (tx Database) {
	t := db.DB.Raw(sql, values...)
	return wrap(t)
}

// Scan scan value to a struct
func (db *DBWrapper) Scan(dest interface{}) (tx Database) {
	t := db.DB.Scan(dest)
	return wrap(t)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (db *DBWrapper) Save(value interface{}) (tx Database) {
	t := db.DB.Save(value)
	return wrap(t)
}

// Model specify the model you would like to run db operations
//    // update all users's name to `hello`
//    db.Model(&User{}).Update("name", "hello")
//    // if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
//    db.Model(&user).Update("name", "hello")
func (db *DBWrapper) Model(value interface{}) (tx Database) {
	t := db.DB.Model(value)
	return wrap(t)
}

// Where add conditions
func (db *DBWrapper) Where(query interface{}, args ...interface{}) (tx Database) {
	t := db.DB.Where(query, args...)
	return wrap(t)
}

// First find first record that match given conditions, order by primary key
func (db *DBWrapper) First(dest interface{}, conds ...interface{}) (tx Database) {
	t := db.DB.First(dest, conds...)
	return wrap(t)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (db *DBWrapper) Updates(values interface{}) (tx Database) {
	t := db.DB.Updates(values)
	return wrap(t)
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (db *DBWrapper) Delete(value interface{}, conds ...interface{}) (tx Database) {
	t := db.DB.Delete(value, conds...)
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
