package lock

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestLock(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	go StartLock(dir + "\\dist\\.LOCK")
	cmd := exec.Command(dir + "\\dist\\CleaningProject.exe")
	cmd.Dir = dir + "\\dist"
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	b := <-C
	if b != OK {
		t.Fail()
	}
}
