package hendler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/SashaMelva/tourism/internal/storage/model"
)

func (s *Service) HendlerAuthAndReg(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if req.URL.Path == "/auth/" {
		if req.Method == http.MethodPost {
			s.authUser(w, req, ctx)
		}
	}
	if req.URL.Path == "/reg/" {
		if req.Method == http.MethodPost {
			s.regUser(w, req, ctx)
		}
	}
}

func (s *Service) authUser(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	var user model.User

	body, err := io.ReadAll(req.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		err = json.Unmarshal(body, &user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}

	id, err := s.app.RegisterUser(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("user id : %v", id)))
}

func (s *Service) regUser(w http.ResponseWriter, req *http.Request, ctx context.Context) {
	s.Logger.Info("handling get all events at %s\n", req.URL.Path)

	allUsers, err := s.app.GetAllUsers(ctx)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(allUsers)

	if err != nil {
		s.Logger.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
