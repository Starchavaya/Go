package swc

import (
	"context"
	"database/sql"
	pb "github.com/Starchavaya/Go/task3/api/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"log"

	//"google.golang.org/grpc/channelz/service"
)

type server struct {
	d *sql.DB
}

func (s *server) DeleteMother(ctx context.Context, request *pb.MotherRequest) (*empty.Empty, error) {
	m := request.GetMother()
	s.d.Exec("delete from Mother where ID = $1", &m.Id)
	return &empty.Empty{}, nil
}

func (s *server) DeleteChild(ctx context.Context, request *pb.MotherAndChildRequest) (*empty.Empty, error) {
	m := request.GetMother()
	c := request.GetChild()
	s.d.Exec("delete from Child where ID = $1 and idMother=$2", &c.Id, &m.Id)
	return &empty.Empty{}, nil
}

func (s *server) UpdateChild(ctx context.Context, request *pb.MotherAndChildRequest) (*empty.Empty, error) {
	m := request.GetMother()
	c := request.GetChild()
	s.d.Exec("update Child set Firstname = $1, Lastname=$2, Patronymic=$3 where ID = $4 and idMother=$5", &c.Firstname, &c.Lastname, &c.Patronymic, &c.Id, &m.Id)
	return &empty.Empty{}, nil
}

func (s *server) UpdateMother(ctx context.Context, request *pb.MotherRequest) (*empty.Empty, error) {
	m := request.GetMother()
	s.d.Exec("update Mother set Firstname = $1, Lastname=$2, Patronymic=$3 where ID = $4", &m.Firstname, &m.Lastname, &m.Patronymic, &m.Id)
	return &empty.Empty{}, nil
}

func (s *server) CreateChild(ctx context.Context, request *pb.MotherAndChildRequest) (*empty.Empty, error) {
	m := request.GetMother()
	c := request.GetChild()
	s.d.Exec("INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values($1,$2,$3,$4)", &c.Firstname, &c.Lastname, &c.Patronymic, &m.Id)
	return &empty.Empty{}, nil
}

func (s *server) CreateMother(ctx context.Context, request *pb.MotherRequest) (*empty.Empty, error) {
	m := request.GetMother()
	if err := s.d.QueryRow(
		"INSERT INTO Mother(Firstname,Lastname,Patronymic) values($1,$2,$3)RETURNING ID", &m.Firstname, &m.Lastname, &m.Patronymic).Scan(&m.Id); err != nil {
		return nil, err
	}
	for _, child := range m.Childs {
		s.d.Exec("INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values($1,$2,$3,$4) ", &child.Firstname, &child.Lastname, &child.Patronymic, &m.Id)
	}
	return &empty.Empty{}, nil
}

func (s *server) GetChildByMother(ctx context.Context, request *pb.MotherAndChildRequest) (*pb.ChildResponse, error) {
	m := request.GetMother()
	c := request.GetChild()
	row := s.d.QueryRow("select ID,Firstname,Lastname,Patronymic from Child where idMother = $1 and id=$2", &m.Id, &c.Id)
	child := &pb.Child{}
	row.Scan(&child.Id, &child.Firstname, &child.Lastname, &child.Patronymic)
	return &pb.ChildResponse{Child: child}, nil
}

func (s *server) GetMothersChilds(ctx context.Context, request *pb.MotherRequest) (*pb.ListChildsResponse, error) {
	m := request.GetMother()
	row := s.d.QueryRow("select * from Mother where id = $1", m.Id)
	mother := &pb.Mother{}
	row.Scan(&mother.Id, &mother.Firstname, &mother.Lastname, &mother.Patronymic)
	s.AddChild(mother)
	return &pb.ListChildsResponse{Childs: mother.Childs}, nil
}

func (s *server) GetMother(ctx context.Context, request *pb.MotherRequest) (*pb.MotherResponse, error) {
	m := request.GetMother()
	row := s.d.QueryRow("select * from Mother where id = $1", m.Id)
	mother := &pb.Mother{}
	row.Scan(&mother.Id, &mother.Firstname, &mother.Lastname, &mother.Patronymic)
	s.AddChild(mother)
	return &pb.MotherResponse{Mother: mother}, nil
}

func (s *server) ListMother(ctx context.Context, request *pb.ListMotherRequest) (*pb.ListMotherResponse, error) {
	rows, err := s.d.Query(
		"SELECT Id,Firstname,Lastname,Patronymic FROM Mother")
	if err != nil {
		log.Println(err)
	}
	mothers := []*pb.Mother{}
	for rows.Next() {
		mother := &pb.Mother{}
		err := rows.Scan(&mother.Id, &mother.Firstname, &mother.Lastname, &mother.Patronymic)
		if err != nil {
			log.Println(err)
		}
		s.AddChild(mother)
		mothers = append(mothers, mother)
	}
	return &pb.ListMotherResponse{Mothers: mothers}, nil
}

func NewFamilyServer(db *sql.DB) *server {
	return &server{d: db}
}

func (s *server) AddChild(m *pb.Mother) *pb.Mother {
	rows, err := s.d.Query(
		"SELECT ID,Firstname,Lastname,Patronymic FROM Child where idMother = $1", m.Id)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		child := &pb.Child{}
		err := rows.Scan(&child.Id, &child.Firstname, &child.Lastname, &child.Patronymic)
		if err != nil {
			log.Println(err)
		}
		m.Childs = append(m.Childs, child)
	}
	return m
}
