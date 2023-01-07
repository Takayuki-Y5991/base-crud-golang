package models

import (
	"base_crud/app/config"
	"base_crud/app/utils"
	"log"
)

type User struct {
	UserId       int
	UserName     string
	Password     string
	Email        string
	DepartmentId int
}

func (u *User) CreateUser() (err error) {
	config.ConnectionDatabase()
	defer config.CloseDatabase()
	cmd := `insert into users (
		user_name,
		password,
		email,
		department_id
		) values (?, ?, ?)`
	_, err = config.DB.Exec(cmd, u.UserName, utils.Encrypt(u.Password), u.DepartmentId)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUsers() (users []User, err error) {
	config.ConnectionDatabase()
	defer config.CloseDatabase()
	cmd := `select
				user_id,
				user_name,
				email,
				department_id 
			from 
				users`
	rows, err := config.DB.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.UserId,
			&user.UserName,
			&user.Email,
			&user.DepartmentId,
		)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	rows.Close()
	return users, err
}

func GetUser(userId int) (user User, err error) {
	config.ConnectionDatabase()
	defer config.CloseDatabase()
	user = User{}
	cmd := `select
				user_id,
				user_name,
				email,
				department_id 
			from 
				users 
			where user_id = ?`
	err = config.DB.QueryRow(cmd, userId).Scan(
		&user.UserId,
		&user.UserName,
		&user.Email,
		&user.DepartmentId,
	)
	return
}

func (u *User) UpdateUser() (err error) {
	config.ConnectionDatabase()
	defer config.CloseDatabase()
	cmd := `update 
				users 
			set 
				user_name = ?,
				email = ?,
				password = ?,
				department_id = ?
			where user_id = ?`
	_, err = config.DB.Exec(cmd, u.UserName, u.Email, u.Password, u.DepartmentId)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	config.ConnectionDatabase()
	defer config.CloseDatabase()
	cmd := `delete from users where user_id = ?`
	_, err = config.DB.Exec(cmd, u.UserId)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
