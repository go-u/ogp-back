// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1-devel
// 	protoc        (unknown)
// source: bookmark.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

type Bookmark struct {
	UserId    uint64               `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Fqdn      string               `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
}
