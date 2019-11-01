package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gorilla/mux"
	"github.com/kinluek/logserver/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logger, err := createLogger(os.Getenv("LOGSTASH"))
	if err != nil {
		log.Fatal("could not initialise logger")
	}

	logger.Info("starting server...")

	app := server.New(mux.NewRouter(), logger)
	app.InitRoutes()
	app.Run(":8080")
}

// createLogger hooks a logstash server into a logrus logger
// using the logstash address and returns it.
func createLogger(address string) (*logrus.Logger, error) {

	l := logrus.New()

	var err error
	// retry connection to logstash incase the server has not yet come up
	for retryCount := 0; retryCount < 20; retryCount++ {
		conn, err := net.Dial("tcp", address)
		if err == nil {

			hook, err := logrustash.NewHookWithConn(conn, "si_test_app")
			if err != nil {
				return nil, err
			}

			l.Hooks.Add(hook)
			return l, err
		}

		fmt.Println("unable to connect to logstash, retrying")
		time.Sleep(2 * time.Second)
	}

	return nil, err
}
