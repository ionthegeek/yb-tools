// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// The following only applies to changes made to this file as part of YugaByte development.
//
// Portions Copyright (c) YugaByte, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied.  See the License for the specific language governing permissions and limitations
// under the License.
//
// Request/Response in different package to test that RPC methods
// handle arguments with packages different from the service itself.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: yb/rpc/rtest_diff_package.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReqDiffPackagePB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqDiffPackagePB) Reset() {
	*x = ReqDiffPackagePB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_rpc_rtest_diff_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqDiffPackagePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqDiffPackagePB) ProtoMessage() {}

func (x *ReqDiffPackagePB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_rpc_rtest_diff_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqDiffPackagePB.ProtoReflect.Descriptor instead.
func (*ReqDiffPackagePB) Descriptor() ([]byte, []int) {
	return file_yb_rpc_rtest_diff_package_proto_rawDescGZIP(), []int{0}
}

type RespDiffPackagePB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RespDiffPackagePB) Reset() {
	*x = RespDiffPackagePB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_rpc_rtest_diff_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespDiffPackagePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespDiffPackagePB) ProtoMessage() {}

func (x *RespDiffPackagePB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_rpc_rtest_diff_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespDiffPackagePB.ProtoReflect.Descriptor instead.
func (*RespDiffPackagePB) Descriptor() ([]byte, []int) {
	return file_yb_rpc_rtest_diff_package_proto_rawDescGZIP(), []int{1}
}

var File_yb_rpc_rtest_diff_package_proto protoreflect.FileDescriptor

var file_yb_rpc_rtest_diff_package_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x79, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64,
	0x69, 0x66, 0x66, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x79, 0x62, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64,
	0x69, 0x66, 0x66, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x52,
	0x65, 0x71, 0x44, 0x69, 0x66, 0x66, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x42, 0x22,
	0x13, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x44, 0x69, 0x66, 0x66, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x50, 0x42,
}

var (
	file_yb_rpc_rtest_diff_package_proto_rawDescOnce sync.Once
	file_yb_rpc_rtest_diff_package_proto_rawDescData = file_yb_rpc_rtest_diff_package_proto_rawDesc
)

func file_yb_rpc_rtest_diff_package_proto_rawDescGZIP() []byte {
	file_yb_rpc_rtest_diff_package_proto_rawDescOnce.Do(func() {
		file_yb_rpc_rtest_diff_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_yb_rpc_rtest_diff_package_proto_rawDescData)
	})
	return file_yb_rpc_rtest_diff_package_proto_rawDescData
}

var file_yb_rpc_rtest_diff_package_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yb_rpc_rtest_diff_package_proto_goTypes = []interface{}{
	(*ReqDiffPackagePB)(nil),  // 0: yb.rpc_test_diff_package.ReqDiffPackagePB
	(*RespDiffPackagePB)(nil), // 1: yb.rpc_test_diff_package.RespDiffPackagePB
}
var file_yb_rpc_rtest_diff_package_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yb_rpc_rtest_diff_package_proto_init() }
func file_yb_rpc_rtest_diff_package_proto_init() {
	if File_yb_rpc_rtest_diff_package_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yb_rpc_rtest_diff_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqDiffPackagePB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yb_rpc_rtest_diff_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespDiffPackagePB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yb_rpc_rtest_diff_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yb_rpc_rtest_diff_package_proto_goTypes,
		DependencyIndexes: file_yb_rpc_rtest_diff_package_proto_depIdxs,
		MessageInfos:      file_yb_rpc_rtest_diff_package_proto_msgTypes,
	}.Build()
	File_yb_rpc_rtest_diff_package_proto = out.File
	file_yb_rpc_rtest_diff_package_proto_rawDesc = nil
	file_yb_rpc_rtest_diff_package_proto_goTypes = nil
	file_yb_rpc_rtest_diff_package_proto_depIdxs = nil
}