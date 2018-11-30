package userapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"
	"strconv"
)

func StartServer(h *Handler) error {
	http.Handle("/users/", h.Main())
	return http.ListenAndServe(":8000", nil)
}

func (h *Handler) Main() http.Handler {
	return http.StripPrefix("/users/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "" {
			switch r.Method {
			case http.MethodGet:
				h.getAllUserHandler(w, r)
			case http.MethodPost:
				h.createUserHandler(w, r)
			default:
				http.NotFound(w, r)
			}
		} else {
			id, err := strconv.Atoi(r.URL.Path)
			if err != nil {
				http.NotFound(w, r)
				return
			}

			switch r.Method {
			case http.MethodGet:
				h.getByIDHandler(id, w, r)
			case http.MethodPost, http.MethodPut:
				h.updateUserHandler(id, w, r)
			case http.MethodDelete:
				h.deleteUserByIDHandler(id, w, r)
			}
		}
	}))
}

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

func (h *Handler) getAllUserHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) getByIDHandler(id int, w http.ResponseWriter, r *http.Request) {
	user, err := h.M.getByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("get user: %s", err), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("get user: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", b)
}

func (h *Handler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("read users: %s", err), http.StatusInternalServerError)
		return
	}

	var u user.User
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, fmt.Sprintf("get users: %s", err), http.StatusInternalServerError)
		return
	}

	err = h.M.createUser(&u)
	if err != nil {
		http.Error(w, fmt.Sprintf("create users: %s", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) updateUserHandler(id int, w http.ResponseWriter, r *http.Request) {
	user, err := h.M.getByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("update user: %s", err), http.StatusInternalServerError)
		return
	}
	err = h.M.updateUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("update user: %s", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) deleteUserByIDHandler(id int, w http.ResponseWriter, r *http.Request) {
	err := h.M.deleteUserByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("create user: %s", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
