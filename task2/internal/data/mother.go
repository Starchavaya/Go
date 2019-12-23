package data

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Mother struct {
	ID         int     `json:"id"`
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
	Patronymic string  `json:"patronymic"`
	Childs     []Child `json:"childs"`
}

type motherData struct {
	d *sql.DB
}

type IMotherData interface {
	GetMothersFromDb() []Mother
	GetMotherFromDb(id int) Mother
	UpdateMotherFromDb(id int, mother Mother)
	DeleteMotherFromDb(id int)
	AddChild(mother Mother) Mother
	CreateMotherInDb(mother Mother)
}

func NewMotherData(d *sql.DB) *motherData {
	return &motherData{d: d}
}

func (m *motherData) CreateMotherInDb(mother Mother) {
	var id int
	m.d.QueryRow("INSERT INTO Mother(Firstname,Lastname,Patronymic) values($1,$2,$3)RETURNING ID", mother.Firstname, mother.Lastname, mother.Patronymic).Scan(&id)
	for _, child := range mother.Childs {
		m.d.Exec("INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values($1,$2,$3,$4) ", child.Firstname, child.Lastname, child.Patronymic, id)
	}
}

func (m *motherData) GetMothersFromDb() []Mother {
	rows, err := m.d.Query(
		"SELECT * FROM Mother")
	if err != nil {
		panic(err)
	}
	mothers := []Mother{}
	for rows.Next() {
		mother := Mother{}
		err := rows.Scan(&mother.ID, &mother.Firstname, &mother.Lastname, &mother.Patronymic)
		if err != nil {
			panic(err)
		}
		mother1 := m.AddChild(mother)
		mothers = append(mothers, mother1)
	}
	return mothers
}

func (m *motherData) GetMotherFromDb(id int) Mother {
	row := m.d.QueryRow("select * from Mother where id = $1", id)
	mother := Mother{}
	row.Scan(&mother.ID, &mother.Firstname, &mother.Lastname, &mother.Patronymic)
	mother = m.AddChild(mother)
	return mother
}

func (m *motherData) UpdateMotherFromDb(id int, mother Mother) {
	m.d.Exec("update Mother set Firstname = $1, Lastname=$2, Patronymic=$3 where ID = $4", mother.Firstname, mother.Lastname, mother.Patronymic, id)
}

func (m *motherData) DeleteMotherFromDb(id int) {
	m.d.Exec("delete from Mother where ID = $1", id)
}

func (m *motherData) AddChild(mother Mother) Mother {
	rows, err := m.d.Query(
		"SELECT ID,Firstname,Lastname,Patronymic FROM Child where idMother = $1", mother.ID)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		child := Child{}
		err := rows.Scan(&child.ID, &child.Firstname, &child.Lastname, &child.Patronymic)
		if err != nil {
			panic(err)
		}
		mother.Childs = append(mother.Childs, child)
	}
	return mother
}
