package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

// Особеностью пакета LOGRUS задание своих полей в стандартной конфигурации время сообщение уровень time msg level
//для создания своих полей в пакете logrus есть фнкция withfield
// Настройки для логирования оторые не имеют отношния  функции main выносятся в фунцию init. Функция init в языке
//Golang запускается перед функцией main

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetLevel(log.InfoLevel)
}

func runAndWait2(a, b int) int {
	c := a + b
	log.WithFields(log.Fields{
		"a": a,
		"b": b,
	}).Infof("c = %d", c)
	time.Sleep(time.Millisecond * 1)
	return c

}
func main() {
	for i := 0; i < 100; i++ {
		a := runAndWait2(i, i*2)
		log.Info(" runAndWai finished....")
		log.Infof("result: %d", a)

		if i > 90 {

			log.WithFields(log.Fields{
				"i": i,
			}).Warn("reached 100...")
		}
		if i == 99 {
			log.Fatal("reached 100...")

		}
	}
	fmt.Println("done")
}
