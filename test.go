// go test file
package api

import (
	"time"
)

type Task struct {
	TaskID      int       `json:"task_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Deadline    time.Time `json:"deadline"`
	DateAdded   time.Time `json:"date_added"`
}

// type this function
// func heeheeFunc() {
//    spaceIndent := "heehee"
// }

// then do :set noexpandtab

// insert a line below the spaceIndent

// backspace

// press tab

// type any letter

// then an error will show up
