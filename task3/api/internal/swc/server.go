package swc

import (
	"context"
	"database/sql"
	pb "github.com/Starchavaya/Go/task3/api/proto"
	//"google.golang.org/grpc/channelz/service"
)

type server struct {
	d *sql.DB
}

func (s *server) ListMother(ctx context.Context, request *pb.ListMotherRequest) (*pb.ListMotherResponse, error) {
	rows, err := s.d.Query(
		"SELECT Id,Firstname,Lastname,Patronymic FROM Mother")
	if err != nil {
		panic(err)
	}
	mothers := []*pb.Mother{}
	for rows.Next() {
		mother := &pb.Mother{}
		err := rows.Scan(&mother.Id, &mother.Firstname, &mother.Lastname, &mother.Patronymic)
		if err != nil {
			panic(err)
		}
		//mother1 := m.AddChild(mother)
		mothers = append(mothers, mother)
	}
	return &pb.ListMotherResponse{Mothers:mothers},nil
}

func NewFamilyServer(db *sql.DB) *server {
	return &server{d:db}
}
