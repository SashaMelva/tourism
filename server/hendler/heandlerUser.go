package hendler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (s *Service) HendlerUser(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if req.URL.Path == "/usesr/" {
		switch req.Method {
		case http.MethodGet:
			s.getAllUsersHandler(w, req, ctx)
		default:
			s.Logger.Error(fmt.Sprintf("expect method GET, DELETE or POST at /usesr/, got %v", req.Method))
			return
		}
	}
}

func (s *Service) getAllUsersHandler(w http.ResponseWriter, req *http.Request, ctx context.Context) {
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
