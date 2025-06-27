package files

import (
	"os"

	"github.com/fatih/color"

	"passGenForMySelf/output"
)

type Db struct {
	filename string
}

func NewJsonDB(name string) *Db {
	return &Db{
		filename: name,
	}
}

func (db *Db) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		output.PrintError(err)
		return nil, err
	}
	return data, nil
}

func (db *Db) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err)
		return
	}
	color.Green("Запись успешно сохранена")
}
