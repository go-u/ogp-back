package testdata

import (
	"github.com/golang/protobuf/ptypes"
	"server/domain/model"
	pb "server/etc/protocol"
	"time"
)

// entity
var User1 = &model.User{
	ID:        1,
	UID:       "1",
	Name:      "user1",
	CreatedAt: time.Date(2020, 10, 10, 10, 10, 10, 10, location),
}

var UserNoName = &model.User{
	ID:        10,
	UID:       "10",
	Name:      "",
	CreatedAt: time.Date(2020, 10, 10, 10, 10, 10, 10, location),
}

// pb
var pbCreatedAtUser1, _ = ptypes.TimestampProto(User1.CreatedAt)
var PbUser1 = &pb.User{
	Id:        User1.ID,
	Uid:       User1.UID,
	Name:      User1.Name,
	CreatedAt: pbCreatedAtUser1,
}

var pbCreatedAtUserNoName, _ = ptypes.TimestampProto(UserNoName.CreatedAt)
var PbUserNoName = &pb.User{
	Id:        UserNoName.ID,
	Uid:       UserNoName.UID,
	Name:      UserNoName.Name,
	CreatedAt: pbCreatedAtUserNoName,
}
