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

// Code generated by protoc-gen-ybrpc. DO NOT EDIT.

package tserver

import (
	"github.com/go-logr/logr"
	"github.com/yugabyte/yb-tools/protoc-gen-ybrpc/pkg/message"
)

// service: yb.tserver.TabletServerAdminService
// service: TabletServerAdminService
type TabletServerAdminService interface {
	CreateTablet(request *CreateTabletRequestPB) (*CreateTabletResponsePB, error)
	DeleteTablet(request *DeleteTabletRequestPB) (*DeleteTabletResponsePB, error)
	AlterSchema(request *ChangeMetadataRequestPB) (*ChangeMetadataResponsePB, error)
	GetSafeTime(request *GetSafeTimeRequestPB) (*GetSafeTimeResponsePB, error)
	BackfillIndex(request *BackfillIndexRequestPB) (*BackfillIndexResponsePB, error)
	BackfillDone(request *ChangeMetadataRequestPB) (*ChangeMetadataResponsePB, error)
	CopartitionTable(request *CopartitionTableRequestPB) (*CopartitionTableResponsePB, error)
	FlushTablets(request *FlushTabletsRequestPB) (*FlushTabletsResponsePB, error)
	CountIntents(request *CountIntentsRequestPB) (*CountIntentsResponsePB, error)
	AddTableToTablet(request *AddTableToTabletRequestPB) (*AddTableToTabletResponsePB, error)
	RemoveTableFromTablet(request *RemoveTableFromTabletRequestPB) (*RemoveTableFromTabletResponsePB, error)
	SplitTablet(request *SplitTabletRequestPB) (*SplitTabletResponsePB, error)
}

type TabletServerAdminServiceImpl struct {
	Log       logr.Logger
	Messenger message.Messenger
}

// Create a new, empty tablet with the specified parameters. Only used for
// brand-new tablets, not for "moves".

func (s *TabletServerAdminServiceImpl) CreateTablet(request *CreateTabletRequestPB) (*CreateTabletResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "CreateTablet", "request", request)
	response := &CreateTabletResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "CreateTablet", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "CreateTablet", "response", response)

	return response, nil
}

// Delete a tablet replica.

func (s *TabletServerAdminServiceImpl) DeleteTablet(request *DeleteTabletRequestPB) (*DeleteTabletResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "DeleteTablet", "request", request)
	response := &DeleteTabletResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "DeleteTablet", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "DeleteTablet", "response", response)

	return response, nil
}

// Alter a tablet's schema.

func (s *TabletServerAdminServiceImpl) AlterSchema(request *ChangeMetadataRequestPB) (*ChangeMetadataResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "AlterSchema", "request", request)
	response := &ChangeMetadataResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "AlterSchema", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "AlterSchema", "response", response)

	return response, nil
}

// GetSafeTime API to get the current safe time.

func (s *TabletServerAdminServiceImpl) GetSafeTime(request *GetSafeTimeRequestPB) (*GetSafeTimeResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "GetSafeTime", "request", request)
	response := &GetSafeTimeResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "GetSafeTime", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "GetSafeTime", "response", response)

	return response, nil
}

// Backfill the index for the specified index tables. Addressed to the indexed
// table.

func (s *TabletServerAdminServiceImpl) BackfillIndex(request *BackfillIndexRequestPB) (*BackfillIndexResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "BackfillIndex", "request", request)
	response := &BackfillIndexResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "BackfillIndex", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "BackfillIndex", "response", response)

	return response, nil
}

// Marks an index table as having completed backfilling.

func (s *TabletServerAdminServiceImpl) BackfillDone(request *ChangeMetadataRequestPB) (*ChangeMetadataResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "BackfillDone", "request", request)
	response := &ChangeMetadataResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "BackfillDone", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "BackfillDone", "response", response)

	return response, nil
}

// Create a co-partitioned table in an existing tablet

func (s *TabletServerAdminServiceImpl) CopartitionTable(request *CopartitionTableRequestPB) (*CopartitionTableResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "CopartitionTable", "request", request)
	response := &CopartitionTableResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "CopartitionTable", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "CopartitionTable", "response", response)

	return response, nil
}

func (s *TabletServerAdminServiceImpl) FlushTablets(request *FlushTabletsRequestPB) (*FlushTabletsResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "FlushTablets", "request", request)
	response := &FlushTabletsResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "FlushTablets", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "FlushTablets", "response", response)

	return response, nil
}

func (s *TabletServerAdminServiceImpl) CountIntents(request *CountIntentsRequestPB) (*CountIntentsResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "CountIntents", "request", request)
	response := &CountIntentsResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "CountIntents", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "CountIntents", "response", response)

	return response, nil
}

func (s *TabletServerAdminServiceImpl) AddTableToTablet(request *AddTableToTabletRequestPB) (*AddTableToTabletResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "AddTableToTablet", "request", request)
	response := &AddTableToTabletResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "AddTableToTablet", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "AddTableToTablet", "response", response)

	return response, nil
}

func (s *TabletServerAdminServiceImpl) RemoveTableFromTablet(request *RemoveTableFromTabletRequestPB) (*RemoveTableFromTabletResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "RemoveTableFromTablet", "request", request)
	response := &RemoveTableFromTabletResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "RemoveTableFromTablet", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "RemoveTableFromTablet", "response", response)

	return response, nil
}

func (s *TabletServerAdminServiceImpl) SplitTablet(request *SplitTabletRequestPB) (*SplitTabletResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.tserver.TabletServerAdminService", "method", "SplitTablet", "request", request)
	response := &SplitTabletResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerAdminService", "SplitTablet", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerAdminService", "method", "SplitTablet", "response", response)

	return response, nil
}
