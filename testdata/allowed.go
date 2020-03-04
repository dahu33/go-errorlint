package testdata

import (
	"database/sql"
	"fmt"
	"io"
	"os"
)

func ReadUntilEOF(r io.Reader) {
	var buf [4096]byte
	for {
		_, err := r.Read(buf[:])
		if err == io.EOF { // ok
			break
		}
	}
}
func ReadFileUntilEOF() {
	fd, _ := os.Open("file.txt")
	var buf [4096]byte
	for {
		_, err := fd.Read(buf[:])
		if err == io.EOF { // ok
			break
		}
	}
}

func MaybeGetSQLRow(db *sql.DB) {
	var i int
	row := db.QueryRow(`SELECT 1`)
	err := row.Scan(&i)

	if err == sql.ErrNoRows { // ok
		fmt.Println("no rows!")
	} else if err != nil {
		fmt.Println("error!")
	}
}
