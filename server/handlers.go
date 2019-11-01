package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *Server) handleExampleGet(w http.ResponseWriter, r *http.Request) {
	s.logger.WithFields(logrus.Fields{
		"path":   r.URL.Path,
		"method": r.Method,
	}).Info("simple get request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func (s *Server) handleExamplePost(w http.ResponseWriter, r *http.Request) {
	var ep examplePost
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&ep); err != nil {
		s.logger.WithFields(logrus.Fields{
			"path":   r.URL.Path,
			"method": r.Method,
		}).Errorf("error occured: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error occured: %v", err)))
	}

	s.logger.WithFields(logrus.Fields{
		"path":   r.URL.Path,
		"method": r.Method,
	}).Info(ep.Data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("data: %v", ep.Data)))
}
