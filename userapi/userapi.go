package userapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/user"
)

type Manager interface {
	getAllUser() ([]user.User, error)
	getByID(id int) (*user.User, error)
	createUser(u *user.User) error
	updateUser(u *user.User) error
	deleteUserByID(id int) error
}

type Handler struct {
	M Manager
}

func (h *Handler) getAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.M.getAllUser()
	if err != nil {
		http.Error(w, fmt.Sprintf("users: %s", err), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(users)
	if err != nil {
		http.Error(w, fmt.Sprintf("users: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", b)
}
