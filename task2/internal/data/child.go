package data

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Child struct {
	ID         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Patronymic string `json:"patronymic"`
}

type childData struct {
	d *sql.DB
}

type IChildData interface {
	GetChildsByMotherIdDb(id int) []Child
	GetChildByMotherAndIdDb(idMother int, idChild int) Child
	CreateChildInDb(child Child, idMother int)
	UpdateChildFromDb(idMother int, idChild int, child Child)
	DeleteChildFromDb(idMother int, idChild int)
}

func NewChildData(d *sql.DB) *childData {
	return &childData{d: d}
}

func (c *childData) GetChildsByMotherIdDb(id int) []Child {
	return NewMotherData(c.d).GetMotherFromDb(id).Childs
}

func (c *childData) GetChildByMotherAndIdDb(idMother int, idChild int) Child {
	row := c.d.QueryRow("select ID,Firstname,Lastname,Patronymic from Child where idMother = $1 and id=$2", idMother, idChild)
	child := Child{}
	row.Scan(&child.ID, &child.Firstname, &child.Lastname, &child.Patronymic)
	return child
}

func (c *childData) CreateChildInDb(child Child, idMother int) {
	c.d.Exec("INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values($1,$2,$3,$4)", child.Firstname, child.Lastname, child.Patronymic, idMother)
}

func (c *childData) UpdateChildFromDb(idMother int, idChild int, child Child) {
	c.d.Exec("update Child set Firstname = $1, Lastname=$2, Patronymic=$3 where ID = $4 and idMother=$5", child.Firstname, child.Lastname, child.Patronymic, idChild, idMother)
}

func (c *childData) DeleteChildFromDb(idMother int, idChild int) {
	c.d.Exec("delete from Child where ID = $1 and idMother=$2", idChild, idMother)
}
