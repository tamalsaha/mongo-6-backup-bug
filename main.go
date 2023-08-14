package main

import (
	"fmt"
	shell "gomodules.xyz/go-sh"
	"log"
	"os/exec"
)

/*
mongosh config --host localhost:27017 --quiet --eval 'JSON.stringify(rs.isMaster())' --username=root --password=hslDW*abfpZZ9Ofv --authenticationDatabase admin
*/

func main2() {
	cmd := exec.Command(
		"/opt/homebrew/bin/mongosh",
		"config",
		"--host", "localhost:27017",
		"--quiet",
		"--eval", "'JSON.stringify(rs.isMaster())'",
		"--username=root",
		"--password=hslDW*abfpZZ9Ofv",
		"--authenticationDatabase admin",
	)
	log.Printf("Running command and waiting for it to finish...")
	out, _ := cmd.Output()
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Printf("The date is %s\n", out)
}

func main3() {
	sh := shell.NewSession()
	sh.SetDir("/tmp")
	sh.ShowCMD = true
	err := sh.Command(
		"/opt/homebrew/bin/mongosh",
		"config",
		"--host", "localhost:27017",
		"--quiet",
		"--eval", "'JSON.stringify(rs.isMaster())'",
		"--username=root",
		"--password=hslDW*abfpZZ9Ofv",
		"--authenticationDatabase admin",
	).Run()
	if err != nil {
		fmt.Println("Err:", err.Error())
	}
}

// /opt/homebrew/bin/mongosh config --host localhost:27017 --quiet --eval 'JSON.stringify(rs.isMaster())' --username root --password hslDW*abfpZZ9Ofv --authenticationDatabase admin
func main() {
	sh := shell.NewSession()
	sh.SetDir("/tmp")
	sh.ShowCMD = true
	err := sh.Command(
		"sh",
		"-c",
		`mongosh config --host localhost:27017 --quiet --eval 'JSON.stringify(rs.isMaster())' --username root --password='hslDW*abfpZZ9Ofv' --authenticationDatabase admin`,
	).Run()
	if err != nil {
		fmt.Println("Err:", err.Error())
	}
}

func main5() {
	sh := shell.NewSession()
	sh.SetDir("/tmp")
	sh.ShowCMD = true
	err := sh.Command(
		"/opt/homebrew/Cellar/mongosh/1.10.4/libexec/bin/mongosh",
		"config",
		"--host", "localhost:27017",
		"--quiet",
		"--eval", "'JSON.stringify(rs.isMaster())'",
		"--username=root",
		"--password=hslDW*abfpZZ9Ofv",
		"--authenticationDatabase admin",
	).Run()
	if err != nil {
		fmt.Println("Err:", err.Error())
	}
}
