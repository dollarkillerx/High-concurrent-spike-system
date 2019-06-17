/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 下午9:09
* */
package repositories

import (
	"High-concurrent-spike-system/datamodels"
	"High-concurrent-spike-system/datasource/databases/mysql_conn"
	"errors"
)

type UserRepositories interface {
	GetByName(string) (*datamodels.User,error)
	Insert(*datamodels.User) (int64,error)
	GetAllData()([]*datamodels.User,error)
	DelUser(name string) error
}

type UserManger struct {

}

func NewUserRepositories() UserRepositories {
	return &UserManger{}
}

func (u *UserManger)GetByName(name string) (*datamodels.User,error){
	sql := "SELECT * FROM WHERE `name` = ?"
	stmt, e := mysql_conn.MySQL.Prepare(sql)
	if e != nil {
		return nil,e
	}
	user := &datamodels.User{}
	e = stmt.QueryRow(name).Scan(&user.Id, &user.Name, &user.Age)
	return user,e
}

func (u *UserManger)Insert(user *datamodels.User) (int64,error) {
	sql := "INSERT INTO `user`(`name`,`age`) VALUE(?,?)"
	stmt, e := mysql_conn.MySQL.Prepare(sql)
	if e != nil {
		return 0,e
	}
	if result, e := stmt.Exec(user.Name, user.Age);e == nil {
		i, e := result.RowsAffected()
		if e == nil {
			if i>0 {
				return i,nil
			}
		}
	}
	return 0,errors.New("insert error")
}

func (u *UserManger)GetAllData()([]*datamodels.User,error) {
	sql :="SELECT * FROM `user`"
	stmt, e := mysql_conn.MySQL.Prepare(sql)
	if e != nil {
		return nil,e
	}
	rows, e := stmt.Query()
	if e != nil {
		return nil,e
	}
	var data []*datamodels.User
	for rows.Next() {
		user := &datamodels.User{}
		e := rows.Scan(&user.Id, &user.Name, &user.Age)
		if e != nil {
			return nil,e
		}
		data = append(data,user)
	}
	return data,nil
}

func (u *UserManger) DelUser(name string) error {
	sql := "DELETE FROM `user` WHERE `name` = ?"
	stmt, e := mysql_conn.MySQL.Prepare(sql)
	if e != nil {
		return e
	}
	if result, e := stmt.Exec(name);e == nil {
		i, e := result.RowsAffected()
		if e == nil {
			if i>0 {
				return nil
			}
		}
	}

	return errors.New("del error")
}