package repository

import (
	"database/sql"
	"log"

	"github.com/tomygp97/weather-notifier/internal/domain"
)

type MySQLUserRepo struct {
	DB *sql.DB
}

func (r *MySQLUserRepo) Save(user *domain.User) error {
	if user.OptedOut == nil {
		user.OptedOut = new(bool)
		*user.OptedOut = false
	}

	result, err := r.DB.Exec("INSERT INTO users (name, email, opted_out) VALUES (?, ?, ?)", user.Name, user.Email, user.OptedOut)
	if err != nil {
		log.Println("Error al insertar usuario:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener LastInsertId:", err)
		return err
	}

	user.ID = int(id)
	return nil
}

func (r *MySQLUserRepo) FindByID(id int) (*domain.User, error) {
	row := r.DB.QueryRow("Select id, name, email, opted_out FROM users WHERE id = ?", id)
	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.OptedOut)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepo) FindAll() ([]domain.User, error) {
	rows, err := r.DB.Query("Select id, name, email, opted_out FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.OptedOut); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *MySQLUserRepo) Update(user *domain.User) error {
	_, err := r.DB.Exec("UPDATE users SET name = ?, email = ?, opted_out = ? WHERE id = ?", user.Name, user.Email, user.OptedOut, user.ID)
	return err
}

func (r *MySQLUserRepo) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
