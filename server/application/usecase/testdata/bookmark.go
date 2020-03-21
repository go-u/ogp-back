package testdata

import (
	"github.com/golang/protobuf/ptypes"
	"server/domain/model"
	pb "server/etc/protocol"
	"time"
)

// entity
var Bookmark1 = &model.Bookmark{
	UserID:    1,
	FQDN:      "example.com",
	CreatedAt: time.Date(2020, 10, 10, 10, 10, 10, 10, location),
}
var Bookmark2 = &model.Bookmark{
	UserID:    2,
	FQDN:      "example.com",
	CreatedAt: time.Date(2020, 10, 10, 10, 10, 10, 10, location),
}

// pb
var pbCreatedAtBookmark1, _ = ptypes.TimestampProto(Bookmark1.CreatedAt)
var PbBookmark1 = &pb.Bookmark{
	UserId:    Bookmark1.UserID,
	Fqdn:      Bookmark1.FQDN,
	CreatedAt: pbCreatedAtBookmark1,
}
var pbCreatedAtBookmark2, _ = ptypes.TimestampProto(Bookmark2.CreatedAt)
var PbBookmark2 = &pb.Bookmark{
	UserId:    Bookmark2.UserID,
	Fqdn:      Bookmark2.FQDN,
	CreatedAt: pbCreatedAtBookmark2,
}

// array
var Bookmarks = []*model.Bookmark{Bookmark1}
var PbBookmarks = []*pb.Bookmark{PbBookmark1}
