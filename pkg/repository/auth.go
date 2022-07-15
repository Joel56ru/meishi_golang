package repository

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"meishi_golang/senti"
)

//AuthPgsql имплементирует структуру для работы с Postgres
type AuthPgsql struct {
	db *sqlx.DB
}

//MyAuthPgsql Репозиторий для работы с базой
func MyAuthPgsql(db *sqlx.DB) *AuthPgsql {
	return &AuthPgsql{db: db}
}

//CreateUser Создание пользователя временное
func (r *AuthPgsql) CreateUser(user senti.UserRegister) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int64
	err = tx.QueryRow(`
		insert into staffs (fio,tip,tip_name) values ($1,'admin','admin') returning staffs_id
	`, user.FIO).Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	//Вставка связанных опций пользователя
	_, err = tx.Exec(`
		insert into acs_staffs (staff_id, login, password_hash) values ($1,$2,$3);
	`, id, user.Username, user.Password)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	_, err = r.db.NamedExec(`
	insert into staffs_logs (parent_id, child_id, tip, newval) values (:parent_id, :child_id, :tip, :newval);
	`, senti.LogUser{NewVal: "Создание", ParentId: id, Tip: "create"})
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	a2 := senti.UserJWT{StuffsID: id, Role: "admin", SecureIp: true, IP: "0.0.0.0", IsAdmin: true}
	mar2, _ := json.Marshal(&a2)
	_, err = tx.Exec(`
		insert into staffs_role (staff_id, attrs, tip) values ($1,$2,'main');
	`, &id, &mar2)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	//Весь список сотрудников
	var at []senti.CategoryUsersST
	_ = r.db.Select(&at, `
		select tip, person, personend from category_staff where tip = 'staff'
	`)
	var c senti.StaffsRole
	c.List = &at
	role, _ := json.Marshal(&c)
	_, _ = tx.Exec(`
		insert into staffs_role (staff_id, attrs, tip) values ($1,$2,'staff');
	`, &id, &role)
	//Весь список клиентов
	var at2 []senti.CategoryUsersST
	var c2 senti.StaffsRole
	_ = r.db.Select(&at2, `
		select tip, person, personend from category_staff where tip = 'client'
	`)
	c2.List = &at2
	role1, _ := json.Marshal(&c2)
	_, _ = tx.Exec(`
		insert into staffs_role (staff_id, attrs, tip) values ($1,$2,'client');
	`, &id, &role1)
	//Весь список соискателей
	var at3 []senti.CategoryUsersST
	var c3 senti.StaffsRole
	_ = r.db.Select(&at3, `
		select tip, person, personend from category_staff where tip = 'aspirant'
	`)
	c3.List = &at3
	role2, _ := json.Marshal(&c3)
	_, _ = tx.Exec(`
		insert into staffs_role (staff_id, attrs, tip) values ($1,$2,'aspirant');
	`, &id, &role2)

	return id, tx.Commit()
}
