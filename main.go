package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/connor-cahill/goDoList/cmd"
	"github.com/connor-cahill/goDoList/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()                  //	finds users home directory
	dbPath := filepath.Join(home, "tasks.db") //	creates a tasks.db file in home dir
	must(db.Init(dbPath))                     //	opens or creates the db and task bucket

	must(cmd.RootCmd.Execute()) //	executes our CLI package
}

//must is an error handler that simply prints error if there is one
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
