package gondm

// DB GONDM DB definition.
type DB struct {
	*Config
	Error error
}

// Create insert the value into database
func (db *DB) Create(value interface{}) *DB {
	return db
}

// First find first record that match given options, order by ID by default
func (db *DB) First(dest interface{}, options *Options) *DB {
	return db
}

// Update update attributes with callbacks
func (db *DB) Update(entity interface{}, column string, value interface{}) *DB {
	return db
}

// Updates update attributes with callbacks
func (db *DB) Updates(entity interface{}, values interface{}) *DB {
	return db
}

// Delete delete value match given options, if the value has ID, then will including the ID as condition
func (db *DB) Delete(value interface{}, options *Options) *DB {
	return db
}
