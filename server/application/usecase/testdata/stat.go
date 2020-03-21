package testdata

import (
	"server/domain/model"
	pb "server/etc/protocol"
	"time"
)

// entity
var Stat1 = &model.Stat{
	Date:        time.Date(2020, 10, 10, 10, 10, 10, 10, location),
	FQDN:        "example.com",
	Host:        "asia-n1",
	Count:       30,
	Title:       "title",
	Description: "description",
	Image:       "https://example.com/image01.png",
	Type:        "website",
	Lang:        "ja",
}

// pb
var PbStat1 = &pb.Stat{
	Fqdn:        Stat1.FQDN,
	Count:       int32(Stat1.Count),
	Title:       Stat1.Title,
	Description: Stat1.Description,
	Image:       Stat1.Image,
	Lang:        Stat1.Lang,
}

// array
var Stats = []*model.Stat{Stat1}
var PbStats = []*pb.Stat{PbStat1}
