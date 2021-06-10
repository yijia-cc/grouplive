package entity

import pb "github.com/yijia-cc/grouplive/proto/golang"

type Unit struct {
	Address   string	`json:"address,omitempty"`
	AptNumber string	`json:"apt_number,omitempty"`
}

func NewUnitFromProto(pbUnit *pb.Unit) Unit {
	return Unit{
		Address:   pbUnit.GetAddress(),
		AptNumber: pbUnit.GetAptNumber(),
	}
}
