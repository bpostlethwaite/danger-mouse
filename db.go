package main

import "os"

type DB struct {
	path string
	file *os.File
}

func (db *DB) create() (err error) {
	file, err := os.Create(db.path)
	if err != nil {
		return
	}
	db.file = file
	return
}

func (db *DB) resize(mb int64) error {
	return os.Truncate(db.path, mb*bytesPerMB)
}

func (db *DB) size() (int64, error) {
	if db.file == nil {
		db.create()
	}
	stat, err := db.file.Stat()
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}
