package user

import "database/sql"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

type Manager struct {
	DB *sql.DB
}

func (m *Manager) getAllUser() ([]User, error) {
	users := []User{}
	r, err := m.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for r.Next() {
		var u User
		err := r.Scan(&u.ID, &u.FirstName, &u.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (m *Manager) getByID(id int) (*User, error) {
	var u User
	r, err := m.DB.Query("SELECT * FROM users WHERE ID = $1", id)
	if err != nil {
		return nil, err
	}
	err = r.Scan(&u.ID, &u.FirstName, &u.LastName)
	if err != nil {
		return nil, err
	}
	return &u, nil

}

func (m *Manager) createUser(u *User) error {
	_, err := m.DB.Exec("INSERT INTO users (FirstName, LastName) VALUES($1,$2)", u.FirstName, u.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) updateUser(u *User) error {
	_, err := m.DB.Exec("UPDATE users SET FirstName = $1 , LastName = $2 WHERE ID = $3", u.FirstName, u.LastName, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) deleteUserByID(id int) error {
	_, err := m.DB.Exec("DELETE FROM users WHERE ID = $1", id)
	if err != nil {
		return err
	}
	return nil
}
