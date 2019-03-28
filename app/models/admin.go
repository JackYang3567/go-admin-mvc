package models

import (
	"fmt"
	"time"
)

type Admin struct {
	Id        int
	Uuid      string
	UserName      string
	Email     string
	Password  string
	LastLoginTime int
	LastLoginIp string
	Type bool
	Status bool
	SessionId string
	GoogleSecret string
	CreatedAt time.Time
}



// Create a new session for an existing admin
func (admin *Admin) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid, email, admin_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, admin_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), admin.Email, admin.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.AdminId, &session.CreatedAt)
	return
}

// Get the session for an existing admin
func (admin *Admin) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, admin_id, created_at FROM sessions WHERE admin_id = $1", admin.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.AdminId, &session.CreatedAt)
	return
}

func (s *Session) SessionByUuid() (session Session, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, admin_id, created_at FROM sessions WHERE uuid = $1", s.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.AdminId, &session.CreatedAt)
	return
}

// Check if session is valid in the database
func (session *Session) CheckAdmin() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid,  email, admin_id, created_at FROM sessions WHERE uuid = $1", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.AdminId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

// Delete session from database
func (session *Session) DeleteAdminByUUID() (err error) {
	statement := "delete from sessions where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

// Get the admin from the session
func (session *Session) Admin() (admin Admin, err error) {
	admin = Admin{}
	err = Db.QueryRow("SELECT id, uuid, username, email, created_at FROM admins WHERE uuid = $1", session.Uuid).
		Scan(&admin.Id, &admin.Uuid, &admin.UserName, &admin.Email, &admin.CreatedAt)
	return
}

// Delete all sessions from database
func SessionDeleteAdminAll() (err error) {
	statement := "delete from sessions"
	_, err = Db.Exec(statement)
	return
}

// Create a new admin, save admin info into the database
func (admin *Admin) CreateAdmin() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.

	statement := "insert into admins (uuid, username, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"

	
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the Admin struct
	err = stmt.QueryRow(createUUID(), admin.UserName, admin.Email, Encrypt(admin.Password), time.Now()).Scan(&admin.Id, &admin.Uuid, &admin.CreatedAt)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// Delete admin from database
func (admin *Admin) Delete() (err error) {
	statement := "delete from admin where adminname<>'admin' && id = $1 "
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Id)
	return
}

// Update admin information in the database
func (admin *Admin) Update() (err error) {
	statement := "update admins set username = $2, email = $3 where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Id, admin.UserName, admin.Email)
	return
}

// Delete all admins from database
func AdminDeleteAll() (err error) {
	statement := "delete from admins where adminname<>'admin' "
	_, err = Db.Exec(statement)
	return
}

// Get all admins in the database and returns it
func Admins() (admins []Admin, err error) {
	rows, err := Db.Query("SELECT id, uuid, username, email, password, created_at FROM admins")
	if err != nil {
		return
	}
	for rows.Next() {
		admin := Admin{}
		if err = rows.Scan(&admin.Id, &admin.Uuid, &admin.UserName, &admin.Email, &admin.Password, &admin.CreatedAt); err != nil {
			return
		}
		admins = append(admins, admin)
	}
	rows.Close()
	return
}

// Get a single admin given the email
func AdminByEmail(email string) (admin Admin, err error) {
	admin = Admin{}
	err = Db.QueryRow("SELECT id, uuid, username, email, password, created_at FROM admins WHERE email = $1", email).
		Scan(&admin.Id, &admin.Uuid, &admin.UserName, &admin.Email, &admin.Password, &admin.CreatedAt)
	return
}

func UserByUsername(username string) (admin Admin, err error) {
	admin = Admin{}
	err = Db.QueryRow("SELECT id, uuid, username, email, password, created_at FROM admins WHERE username = $1", username).
		Scan(&admin.Id, &admin.Uuid, &admin.UserName, &admin.Email, &admin.Password, &admin.CreatedAt)
	return
}

func AdminByID(id int) (admin Admin, err error) {
	admin = Admin{}
	err = Db.QueryRow("SELECT id, uuid, username, email, password, created_at FROM admins WHERE id = $1", id).
		Scan(&admin.Id, &admin.Uuid, &admin.UserName, &admin.Email, &admin.Password, &admin.CreatedAt)
	return
}
// Get a single admin given the UUID
func AdminByUUID(uuid string) (admin Admin, err error) {
	admin = Admin{}
	err = Db.QueryRow("SELECT id, uuid, username, email, password, created_at FROM admins WHERE uuid = $1", uuid).
		Scan(&admin.Id, &admin.Uuid, &admin.UserName, &admin.Email, &admin.Password, &admin.CreatedAt)
	return
}
