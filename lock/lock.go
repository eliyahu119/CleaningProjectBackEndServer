package lock

import (
	"log"
	"os"
	"sync"
	"time"
)

var C chan string
var lock sync.Mutex

const (
	ERR = "something got wrong!"
	OK  = "OK!"
)

func init() {
	C = make(chan string)
}

//StartLock is a sync function for checking if lock exist (shouldnt be on a goroutine)
func StartLock(url string) {
	lock.Lock()
	defer lock.Unlock()
	isSafe := false
	go func() {
		go func() {
			timer1 := time.NewTimer(30 * time.Second)
			<-timer1.C
			if isSafe == false {
				C <- ERR
				if fileExists(url) {
					err := os.Remove(url)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}()
		go func() {
			for true {
				if fileExists(url) {
					break
				}
			}
			for true {
				if !fileExists(url) {
					isSafe = true
					C <- OK
					break
				}
			}
		}()
	}()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
