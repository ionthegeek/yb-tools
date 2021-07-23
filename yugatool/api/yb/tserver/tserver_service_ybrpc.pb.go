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

// service: yb.tserver.TabletServerService
// service: TabletServerService
type TabletServerService interface {
	Write(request *WriteRequestPB) (*WriteResponsePB, error)
	Read(request *ReadRequestPB) (*ReadResponsePB, error)
	NoOp(request *NoOpRequestPB) (*NoOpResponsePB, error)
	ListTablets(request *ListTabletsRequestPB) (*ListTabletsResponsePB, error)
	GetLogLocation(request *GetLogLocationRequestPB) (*GetLogLocationResponsePB, error)
	Checksum(request *ChecksumRequestPB) (*ChecksumResponsePB, error)
	ListTabletsForTabletServer(request *ListTabletsForTabletServerRequestPB) (*ListTabletsForTabletServerResponsePB, error)
	ImportData(request *ImportDataRequestPB) (*ImportDataResponsePB, error)
	UpdateTransaction(request *UpdateTransactionRequestPB) (*UpdateTransactionResponsePB, error)
	GetTransactionStatus(request *GetTransactionStatusRequestPB) (*GetTransactionStatusResponsePB, error)
	GetTransactionStatusAtParticipant(request *GetTransactionStatusAtParticipantRequestPB) (*GetTransactionStatusAtParticipantResponsePB, error)
	AbortTransaction(request *AbortTransactionRequestPB) (*AbortTransactionResponsePB, error)
	Truncate(request *TruncateRequestPB) (*TruncateResponsePB, error)
	GetTabletStatus(request *GetTabletStatusRequestPB) (*GetTabletStatusResponsePB, error)
	GetMasterAddresses(request *GetMasterAddressesRequestPB) (*GetMasterAddressesResponsePB, error)
	Publish(request *PublishRequestPB) (*PublishResponsePB, error)
	IsTabletServerReady(request *IsTabletServerReadyRequestPB) (*IsTabletServerReadyResponsePB, error)
	TakeTransaction(request *TakeTransactionRequestPB) (*TakeTransactionResponsePB, error)
}

type TabletServerServiceImpl struct {
	Log       logr.Logger
	Messenger message.Messenger
}

func (s *TabletServerServiceImpl) Write(request *WriteRequestPB) (*WriteResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "Write", "message", request)
	response := &WriteResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "Write", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "Write", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) Read(request *ReadRequestPB) (*ReadResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "Read", "message", request)
	response := &ReadResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "Read", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "Read", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) NoOp(request *NoOpRequestPB) (*NoOpResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "NoOp", "message", request)
	response := &NoOpResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "NoOp", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "NoOp", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) ListTablets(request *ListTabletsRequestPB) (*ListTabletsResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "ListTablets", "message", request)
	response := &ListTabletsResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "ListTablets", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "ListTablets", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) GetLogLocation(request *GetLogLocationRequestPB) (*GetLogLocationResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "GetLogLocation", "message", request)
	response := &GetLogLocationResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "GetLogLocation", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "GetLogLocation", "message", response)

	return response, nil
}

// Run full-scan data checksum on a tablet to verify data integrity.
//
// TODO: Consider refactoring this as a scan that runs a checksum aggregation
// function.

func (s *TabletServerServiceImpl) Checksum(request *ChecksumRequestPB) (*ChecksumResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "Checksum", "message", request)
	response := &ChecksumResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "Checksum", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "Checksum", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) ListTabletsForTabletServer(request *ListTabletsForTabletServerRequestPB) (*ListTabletsForTabletServerResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "ListTabletsForTabletServer", "message", request)
	response := &ListTabletsForTabletServerResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "ListTabletsForTabletServer", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "ListTabletsForTabletServer", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) ImportData(request *ImportDataRequestPB) (*ImportDataResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "ImportData", "message", request)
	response := &ImportDataResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "ImportData", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "ImportData", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) UpdateTransaction(request *UpdateTransactionRequestPB) (*UpdateTransactionResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "UpdateTransaction", "message", request)
	response := &UpdateTransactionResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "UpdateTransaction", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "UpdateTransaction", "message", response)

	return response, nil
}

// Returns transaction status at coordinator, i.e. PENDING, ABORTED, COMMITTED etc.

func (s *TabletServerServiceImpl) GetTransactionStatus(request *GetTransactionStatusRequestPB) (*GetTransactionStatusResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "GetTransactionStatus", "message", request)
	response := &GetTransactionStatusResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "GetTransactionStatus", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "GetTransactionStatus", "message", response)

	return response, nil
}

// Returns transaction status at participant, i.e. number of replicated batches or whether it was
// aborted.

func (s *TabletServerServiceImpl) GetTransactionStatusAtParticipant(request *GetTransactionStatusAtParticipantRequestPB) (*GetTransactionStatusAtParticipantResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "GetTransactionStatusAtParticipant", "message", request)
	response := &GetTransactionStatusAtParticipantResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "GetTransactionStatusAtParticipant", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "GetTransactionStatusAtParticipant", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) AbortTransaction(request *AbortTransactionRequestPB) (*AbortTransactionResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "AbortTransaction", "message", request)
	response := &AbortTransactionResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "AbortTransaction", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "AbortTransaction", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) Truncate(request *TruncateRequestPB) (*TruncateResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "Truncate", "message", request)
	response := &TruncateResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "Truncate", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "Truncate", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) GetTabletStatus(request *GetTabletStatusRequestPB) (*GetTabletStatusResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "GetTabletStatus", "message", request)
	response := &GetTabletStatusResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "GetTabletStatus", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "GetTabletStatus", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) GetMasterAddresses(request *GetMasterAddressesRequestPB) (*GetMasterAddressesResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "GetMasterAddresses", "message", request)
	response := &GetMasterAddressesResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "GetMasterAddresses", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "GetMasterAddresses", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) Publish(request *PublishRequestPB) (*PublishResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "Publish", "message", request)
	response := &PublishResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "Publish", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "Publish", "message", response)

	return response, nil
}

func (s *TabletServerServiceImpl) IsTabletServerReady(request *IsTabletServerReadyRequestPB) (*IsTabletServerReadyResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "IsTabletServerReady", "message", request)
	response := &IsTabletServerReadyResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "IsTabletServerReady", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "IsTabletServerReady", "message", response)

	return response, nil
}

// Takes precreated transaction from this tserver.

func (s *TabletServerServiceImpl) TakeTransaction(request *TakeTransactionRequestPB) (*TakeTransactionResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.tserver.TabletServerService", "method", "TakeTransaction", "message", request)
	response := &TakeTransactionResponsePB{}

	err := s.Messenger.SendMessage("yb.tserver.TabletServerService", "TakeTransaction", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.tserver.TabletServerService", "method", "TakeTransaction", "message", response)

	return response, nil
}