// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1-devel
// 	protoc        (unknown)
// source: tweet.proto

package pb

type Tweet struct {
	Url  string `protobuf:"bytes,1,opt,name=url,proto3" json:"url"`
	Html string `protobuf:"bytes,2,opt,name=html,proto3" json:"html"`
}
