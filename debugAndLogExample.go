package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func runAndWait() int {
	time.Sleep(time.Second * 1)
	return 10

}
func main() {
	for i := 0; i < 5; i++ {
		a := runAndWait()
		log.Info(" runAndWai finished....")
		log.Infof("result: %d", a)
	}
	fmt.Println("done")
}
