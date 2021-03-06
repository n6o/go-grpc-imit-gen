// Code generated by go-grpc-imit-gen; DO NOT EDIT.
// github.com/n6o/go-grpc-imit-gen

package imit_spanner

import (
	context "context"
	sync "sync"

	spanner "google.golang.org/genproto/googleapis/spanner/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var _ spanner.SpannerServer = &ImitSpannerServer{}

type ImitSpannerServer struct {
	spanner.UnimplementedSpannerServer

	// mocks
	mocks struct {
		BatchCreateSessions []func(P0 context.Context, P1 *spanner.BatchCreateSessionsRequest) (*spanner.BatchCreateSessionsResponse, error)
		BeginTransaction    []func(P0 context.Context, P1 *spanner.BeginTransactionRequest) (*spanner.Transaction, error)
		Commit              []func(P0 context.Context, P1 *spanner.CommitRequest) (*spanner.CommitResponse, error)
		CreateSession       []func(P0 context.Context, P1 *spanner.CreateSessionRequest) (*spanner.Session, error)
		DeleteSession       []func(P0 context.Context, P1 *spanner.DeleteSessionRequest) (*emptypb.Empty, error)
		ExecuteBatchDml     []func(P0 context.Context, P1 *spanner.ExecuteBatchDmlRequest) (*spanner.ExecuteBatchDmlResponse, error)
		ExecuteSql          []func(P0 context.Context, P1 *spanner.ExecuteSqlRequest) (*spanner.ResultSet, error)
		ExecuteStreamingSql []func(P0 *spanner.ExecuteSqlRequest, P1 spanner.Spanner_ExecuteStreamingSqlServer) error
		GetSession          []func(P0 context.Context, P1 *spanner.GetSessionRequest) (*spanner.Session, error)
		ListSessions        []func(P0 context.Context, P1 *spanner.ListSessionsRequest) (*spanner.ListSessionsResponse, error)
		PartitionQuery      []func(P0 context.Context, P1 *spanner.PartitionQueryRequest) (*spanner.PartitionResponse, error)
		PartitionRead       []func(P0 context.Context, P1 *spanner.PartitionReadRequest) (*spanner.PartitionResponse, error)
		Read                []func(P0 context.Context, P1 *spanner.ReadRequest) (*spanner.ResultSet, error)
		Rollback            []func(P0 context.Context, P1 *spanner.RollbackRequest) (*emptypb.Empty, error)
		StreamingRead       []func(P0 *spanner.ReadRequest, P1 spanner.Spanner_StreamingReadServer) error
	}

	// records
	records struct {
		BatchCreateSessions []struct {
			P0 context.Context
			P1 *spanner.BatchCreateSessionsRequest
		}
		BeginTransaction []struct {
			P0 context.Context
			P1 *spanner.BeginTransactionRequest
		}
		Commit []struct {
			P0 context.Context
			P1 *spanner.CommitRequest
		}
		CreateSession []struct {
			P0 context.Context
			P1 *spanner.CreateSessionRequest
		}
		DeleteSession []struct {
			P0 context.Context
			P1 *spanner.DeleteSessionRequest
		}
		ExecuteBatchDml []struct {
			P0 context.Context
			P1 *spanner.ExecuteBatchDmlRequest
		}
		ExecuteSql []struct {
			P0 context.Context
			P1 *spanner.ExecuteSqlRequest
		}
		ExecuteStreamingSql []struct {
			P0 *spanner.ExecuteSqlRequest
			P1 spanner.Spanner_ExecuteStreamingSqlServer
		}
		GetSession []struct {
			P0 context.Context
			P1 *spanner.GetSessionRequest
		}
		ListSessions []struct {
			P0 context.Context
			P1 *spanner.ListSessionsRequest
		}
		PartitionQuery []struct {
			P0 context.Context
			P1 *spanner.PartitionQueryRequest
		}
		PartitionRead []struct {
			P0 context.Context
			P1 *spanner.PartitionReadRequest
		}
		Read []struct {
			P0 context.Context
			P1 *spanner.ReadRequest
		}
		Rollback []struct {
			P0 context.Context
			P1 *spanner.RollbackRequest
		}
		StreamingRead []struct {
			P0 *spanner.ReadRequest
			P1 spanner.Spanner_StreamingReadServer
		}
	}

	// locks
	lockBatchCreateSessions sync.RWMutex
	lockBeginTransaction    sync.RWMutex
	lockCommit              sync.RWMutex
	lockCreateSession       sync.RWMutex
	lockDeleteSession       sync.RWMutex
	lockExecuteBatchDml     sync.RWMutex
	lockExecuteSql          sync.RWMutex
	lockExecuteStreamingSql sync.RWMutex
	lockGetSession          sync.RWMutex
	lockListSessions        sync.RWMutex
	lockPartitionQuery      sync.RWMutex
	lockPartitionRead       sync.RWMutex
	lockRead                sync.RWMutex
	lockRollback            sync.RWMutex
	lockStreamingRead       sync.RWMutex
}

// BatchCreateSessions Enqueue Mock
func (imit *ImitSpannerServer) EnqueueBatchCreateSessionsMock(m func(P0 context.Context, P1 *spanner.BatchCreateSessionsRequest) (*spanner.BatchCreateSessionsResponse, error)) {
	imit.lockBatchCreateSessions.Lock()
	imit.mocks.BatchCreateSessions = append(imit.mocks.BatchCreateSessions, m)
	imit.lockBatchCreateSessions.Unlock()
}

// BatchCreateSessions Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueBatchCreateSessionsMocks(ms []func(P0 context.Context, P1 *spanner.BatchCreateSessionsRequest) (*spanner.BatchCreateSessionsResponse, error)) {
	imit.lockBatchCreateSessions.Lock()
	imit.mocks.BatchCreateSessions = append(imit.mocks.BatchCreateSessions, ms...)
	imit.lockBatchCreateSessions.Unlock()
}

// BatchCreateSessions Take Records
func (imit *ImitSpannerServer) TakeBatchCreateSessionsRecords() []struct {
	P0 context.Context
	P1 *spanner.BatchCreateSessionsRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.BatchCreateSessionsRequest
	}

	// clear records
	imit.lockBatchCreateSessions.Lock()
	records = imit.records.BatchCreateSessions
	imit.records.BatchCreateSessions = nil
	imit.lockBatchCreateSessions.Unlock()

	return records
}

// BatchCreateSessions Imitation
func (imit *ImitSpannerServer) BatchCreateSessions(P0 context.Context, P1 *spanner.BatchCreateSessionsRequest) (*spanner.BatchCreateSessionsResponse, error) {
	if len(imit.mocks.BatchCreateSessions) == 0 {
		panic("ImitSpannerServer.BatchCreateSessions mocks is nil but ImitSpannerServer.BatchCreateSessions was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.BatchCreateSessionsRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockBatchCreateSessions.Lock()
	// record call
	imit.records.BatchCreateSessions = append(imit.records.BatchCreateSessions, call)
	// pop mock
	mockFunc := imit.mocks.BatchCreateSessions[0]
	imit.mocks.BatchCreateSessions = imit.mocks.BatchCreateSessions[1:]
	imit.lockBatchCreateSessions.Unlock()

	return mockFunc(P0, P1)
}

// BeginTransaction Enqueue Mock
func (imit *ImitSpannerServer) EnqueueBeginTransactionMock(m func(P0 context.Context, P1 *spanner.BeginTransactionRequest) (*spanner.Transaction, error)) {
	imit.lockBeginTransaction.Lock()
	imit.mocks.BeginTransaction = append(imit.mocks.BeginTransaction, m)
	imit.lockBeginTransaction.Unlock()
}

// BeginTransaction Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueBeginTransactionMocks(ms []func(P0 context.Context, P1 *spanner.BeginTransactionRequest) (*spanner.Transaction, error)) {
	imit.lockBeginTransaction.Lock()
	imit.mocks.BeginTransaction = append(imit.mocks.BeginTransaction, ms...)
	imit.lockBeginTransaction.Unlock()
}

// BeginTransaction Take Records
func (imit *ImitSpannerServer) TakeBeginTransactionRecords() []struct {
	P0 context.Context
	P1 *spanner.BeginTransactionRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.BeginTransactionRequest
	}

	// clear records
	imit.lockBeginTransaction.Lock()
	records = imit.records.BeginTransaction
	imit.records.BeginTransaction = nil
	imit.lockBeginTransaction.Unlock()

	return records
}

// BeginTransaction Imitation
func (imit *ImitSpannerServer) BeginTransaction(P0 context.Context, P1 *spanner.BeginTransactionRequest) (*spanner.Transaction, error) {
	if len(imit.mocks.BeginTransaction) == 0 {
		panic("ImitSpannerServer.BeginTransaction mocks is nil but ImitSpannerServer.BeginTransaction was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.BeginTransactionRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockBeginTransaction.Lock()
	// record call
	imit.records.BeginTransaction = append(imit.records.BeginTransaction, call)
	// pop mock
	mockFunc := imit.mocks.BeginTransaction[0]
	imit.mocks.BeginTransaction = imit.mocks.BeginTransaction[1:]
	imit.lockBeginTransaction.Unlock()

	return mockFunc(P0, P1)
}

// Commit Enqueue Mock
func (imit *ImitSpannerServer) EnqueueCommitMock(m func(P0 context.Context, P1 *spanner.CommitRequest) (*spanner.CommitResponse, error)) {
	imit.lockCommit.Lock()
	imit.mocks.Commit = append(imit.mocks.Commit, m)
	imit.lockCommit.Unlock()
}

// Commit Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueCommitMocks(ms []func(P0 context.Context, P1 *spanner.CommitRequest) (*spanner.CommitResponse, error)) {
	imit.lockCommit.Lock()
	imit.mocks.Commit = append(imit.mocks.Commit, ms...)
	imit.lockCommit.Unlock()
}

// Commit Take Records
func (imit *ImitSpannerServer) TakeCommitRecords() []struct {
	P0 context.Context
	P1 *spanner.CommitRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.CommitRequest
	}

	// clear records
	imit.lockCommit.Lock()
	records = imit.records.Commit
	imit.records.Commit = nil
	imit.lockCommit.Unlock()

	return records
}

// Commit Imitation
func (imit *ImitSpannerServer) Commit(P0 context.Context, P1 *spanner.CommitRequest) (*spanner.CommitResponse, error) {
	if len(imit.mocks.Commit) == 0 {
		panic("ImitSpannerServer.Commit mocks is nil but ImitSpannerServer.Commit was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.CommitRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockCommit.Lock()
	// record call
	imit.records.Commit = append(imit.records.Commit, call)
	// pop mock
	mockFunc := imit.mocks.Commit[0]
	imit.mocks.Commit = imit.mocks.Commit[1:]
	imit.lockCommit.Unlock()

	return mockFunc(P0, P1)
}

// CreateSession Enqueue Mock
func (imit *ImitSpannerServer) EnqueueCreateSessionMock(m func(P0 context.Context, P1 *spanner.CreateSessionRequest) (*spanner.Session, error)) {
	imit.lockCreateSession.Lock()
	imit.mocks.CreateSession = append(imit.mocks.CreateSession, m)
	imit.lockCreateSession.Unlock()
}

// CreateSession Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueCreateSessionMocks(ms []func(P0 context.Context, P1 *spanner.CreateSessionRequest) (*spanner.Session, error)) {
	imit.lockCreateSession.Lock()
	imit.mocks.CreateSession = append(imit.mocks.CreateSession, ms...)
	imit.lockCreateSession.Unlock()
}

// CreateSession Take Records
func (imit *ImitSpannerServer) TakeCreateSessionRecords() []struct {
	P0 context.Context
	P1 *spanner.CreateSessionRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.CreateSessionRequest
	}

	// clear records
	imit.lockCreateSession.Lock()
	records = imit.records.CreateSession
	imit.records.CreateSession = nil
	imit.lockCreateSession.Unlock()

	return records
}

// CreateSession Imitation
func (imit *ImitSpannerServer) CreateSession(P0 context.Context, P1 *spanner.CreateSessionRequest) (*spanner.Session, error) {
	if len(imit.mocks.CreateSession) == 0 {
		panic("ImitSpannerServer.CreateSession mocks is nil but ImitSpannerServer.CreateSession was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.CreateSessionRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockCreateSession.Lock()
	// record call
	imit.records.CreateSession = append(imit.records.CreateSession, call)
	// pop mock
	mockFunc := imit.mocks.CreateSession[0]
	imit.mocks.CreateSession = imit.mocks.CreateSession[1:]
	imit.lockCreateSession.Unlock()

	return mockFunc(P0, P1)
}

// DeleteSession Enqueue Mock
func (imit *ImitSpannerServer) EnqueueDeleteSessionMock(m func(P0 context.Context, P1 *spanner.DeleteSessionRequest) (*emptypb.Empty, error)) {
	imit.lockDeleteSession.Lock()
	imit.mocks.DeleteSession = append(imit.mocks.DeleteSession, m)
	imit.lockDeleteSession.Unlock()
}

// DeleteSession Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueDeleteSessionMocks(ms []func(P0 context.Context, P1 *spanner.DeleteSessionRequest) (*emptypb.Empty, error)) {
	imit.lockDeleteSession.Lock()
	imit.mocks.DeleteSession = append(imit.mocks.DeleteSession, ms...)
	imit.lockDeleteSession.Unlock()
}

// DeleteSession Take Records
func (imit *ImitSpannerServer) TakeDeleteSessionRecords() []struct {
	P0 context.Context
	P1 *spanner.DeleteSessionRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.DeleteSessionRequest
	}

	// clear records
	imit.lockDeleteSession.Lock()
	records = imit.records.DeleteSession
	imit.records.DeleteSession = nil
	imit.lockDeleteSession.Unlock()

	return records
}

// DeleteSession Imitation
func (imit *ImitSpannerServer) DeleteSession(P0 context.Context, P1 *spanner.DeleteSessionRequest) (*emptypb.Empty, error) {
	if len(imit.mocks.DeleteSession) == 0 {
		panic("ImitSpannerServer.DeleteSession mocks is nil but ImitSpannerServer.DeleteSession was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.DeleteSessionRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockDeleteSession.Lock()
	// record call
	imit.records.DeleteSession = append(imit.records.DeleteSession, call)
	// pop mock
	mockFunc := imit.mocks.DeleteSession[0]
	imit.mocks.DeleteSession = imit.mocks.DeleteSession[1:]
	imit.lockDeleteSession.Unlock()

	return mockFunc(P0, P1)
}

// ExecuteBatchDml Enqueue Mock
func (imit *ImitSpannerServer) EnqueueExecuteBatchDmlMock(m func(P0 context.Context, P1 *spanner.ExecuteBatchDmlRequest) (*spanner.ExecuteBatchDmlResponse, error)) {
	imit.lockExecuteBatchDml.Lock()
	imit.mocks.ExecuteBatchDml = append(imit.mocks.ExecuteBatchDml, m)
	imit.lockExecuteBatchDml.Unlock()
}

// ExecuteBatchDml Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueExecuteBatchDmlMocks(ms []func(P0 context.Context, P1 *spanner.ExecuteBatchDmlRequest) (*spanner.ExecuteBatchDmlResponse, error)) {
	imit.lockExecuteBatchDml.Lock()
	imit.mocks.ExecuteBatchDml = append(imit.mocks.ExecuteBatchDml, ms...)
	imit.lockExecuteBatchDml.Unlock()
}

// ExecuteBatchDml Take Records
func (imit *ImitSpannerServer) TakeExecuteBatchDmlRecords() []struct {
	P0 context.Context
	P1 *spanner.ExecuteBatchDmlRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.ExecuteBatchDmlRequest
	}

	// clear records
	imit.lockExecuteBatchDml.Lock()
	records = imit.records.ExecuteBatchDml
	imit.records.ExecuteBatchDml = nil
	imit.lockExecuteBatchDml.Unlock()

	return records
}

// ExecuteBatchDml Imitation
func (imit *ImitSpannerServer) ExecuteBatchDml(P0 context.Context, P1 *spanner.ExecuteBatchDmlRequest) (*spanner.ExecuteBatchDmlResponse, error) {
	if len(imit.mocks.ExecuteBatchDml) == 0 {
		panic("ImitSpannerServer.ExecuteBatchDml mocks is nil but ImitSpannerServer.ExecuteBatchDml was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.ExecuteBatchDmlRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockExecuteBatchDml.Lock()
	// record call
	imit.records.ExecuteBatchDml = append(imit.records.ExecuteBatchDml, call)
	// pop mock
	mockFunc := imit.mocks.ExecuteBatchDml[0]
	imit.mocks.ExecuteBatchDml = imit.mocks.ExecuteBatchDml[1:]
	imit.lockExecuteBatchDml.Unlock()

	return mockFunc(P0, P1)
}

// ExecuteSql Enqueue Mock
func (imit *ImitSpannerServer) EnqueueExecuteSqlMock(m func(P0 context.Context, P1 *spanner.ExecuteSqlRequest) (*spanner.ResultSet, error)) {
	imit.lockExecuteSql.Lock()
	imit.mocks.ExecuteSql = append(imit.mocks.ExecuteSql, m)
	imit.lockExecuteSql.Unlock()
}

// ExecuteSql Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueExecuteSqlMocks(ms []func(P0 context.Context, P1 *spanner.ExecuteSqlRequest) (*spanner.ResultSet, error)) {
	imit.lockExecuteSql.Lock()
	imit.mocks.ExecuteSql = append(imit.mocks.ExecuteSql, ms...)
	imit.lockExecuteSql.Unlock()
}

// ExecuteSql Take Records
func (imit *ImitSpannerServer) TakeExecuteSqlRecords() []struct {
	P0 context.Context
	P1 *spanner.ExecuteSqlRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.ExecuteSqlRequest
	}

	// clear records
	imit.lockExecuteSql.Lock()
	records = imit.records.ExecuteSql
	imit.records.ExecuteSql = nil
	imit.lockExecuteSql.Unlock()

	return records
}

// ExecuteSql Imitation
func (imit *ImitSpannerServer) ExecuteSql(P0 context.Context, P1 *spanner.ExecuteSqlRequest) (*spanner.ResultSet, error) {
	if len(imit.mocks.ExecuteSql) == 0 {
		panic("ImitSpannerServer.ExecuteSql mocks is nil but ImitSpannerServer.ExecuteSql was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.ExecuteSqlRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockExecuteSql.Lock()
	// record call
	imit.records.ExecuteSql = append(imit.records.ExecuteSql, call)
	// pop mock
	mockFunc := imit.mocks.ExecuteSql[0]
	imit.mocks.ExecuteSql = imit.mocks.ExecuteSql[1:]
	imit.lockExecuteSql.Unlock()

	return mockFunc(P0, P1)
}

// ExecuteStreamingSql Enqueue Mock
func (imit *ImitSpannerServer) EnqueueExecuteStreamingSqlMock(m func(P0 *spanner.ExecuteSqlRequest, P1 spanner.Spanner_ExecuteStreamingSqlServer) error) {
	imit.lockExecuteStreamingSql.Lock()
	imit.mocks.ExecuteStreamingSql = append(imit.mocks.ExecuteStreamingSql, m)
	imit.lockExecuteStreamingSql.Unlock()
}

// ExecuteStreamingSql Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueExecuteStreamingSqlMocks(ms []func(P0 *spanner.ExecuteSqlRequest, P1 spanner.Spanner_ExecuteStreamingSqlServer) error) {
	imit.lockExecuteStreamingSql.Lock()
	imit.mocks.ExecuteStreamingSql = append(imit.mocks.ExecuteStreamingSql, ms...)
	imit.lockExecuteStreamingSql.Unlock()
}

// ExecuteStreamingSql Take Records
func (imit *ImitSpannerServer) TakeExecuteStreamingSqlRecords() []struct {
	P0 *spanner.ExecuteSqlRequest
	P1 spanner.Spanner_ExecuteStreamingSqlServer
} {
	var records []struct {
		P0 *spanner.ExecuteSqlRequest
		P1 spanner.Spanner_ExecuteStreamingSqlServer
	}

	// clear records
	imit.lockExecuteStreamingSql.Lock()
	records = imit.records.ExecuteStreamingSql
	imit.records.ExecuteStreamingSql = nil
	imit.lockExecuteStreamingSql.Unlock()

	return records
}

// ExecuteStreamingSql Imitation
func (imit *ImitSpannerServer) ExecuteStreamingSql(P0 *spanner.ExecuteSqlRequest, P1 spanner.Spanner_ExecuteStreamingSqlServer) error {
	if len(imit.mocks.ExecuteStreamingSql) == 0 {
		panic("ImitSpannerServer.ExecuteStreamingSql mocks is nil but ImitSpannerServer.ExecuteStreamingSql was just called")
	}

	call := struct {
		P0 *spanner.ExecuteSqlRequest
		P1 spanner.Spanner_ExecuteStreamingSqlServer
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockExecuteStreamingSql.Lock()
	// record call
	imit.records.ExecuteStreamingSql = append(imit.records.ExecuteStreamingSql, call)
	// pop mock
	mockFunc := imit.mocks.ExecuteStreamingSql[0]
	imit.mocks.ExecuteStreamingSql = imit.mocks.ExecuteStreamingSql[1:]
	imit.lockExecuteStreamingSql.Unlock()

	return mockFunc(P0, P1)
}

// GetSession Enqueue Mock
func (imit *ImitSpannerServer) EnqueueGetSessionMock(m func(P0 context.Context, P1 *spanner.GetSessionRequest) (*spanner.Session, error)) {
	imit.lockGetSession.Lock()
	imit.mocks.GetSession = append(imit.mocks.GetSession, m)
	imit.lockGetSession.Unlock()
}

// GetSession Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueGetSessionMocks(ms []func(P0 context.Context, P1 *spanner.GetSessionRequest) (*spanner.Session, error)) {
	imit.lockGetSession.Lock()
	imit.mocks.GetSession = append(imit.mocks.GetSession, ms...)
	imit.lockGetSession.Unlock()
}

// GetSession Take Records
func (imit *ImitSpannerServer) TakeGetSessionRecords() []struct {
	P0 context.Context
	P1 *spanner.GetSessionRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.GetSessionRequest
	}

	// clear records
	imit.lockGetSession.Lock()
	records = imit.records.GetSession
	imit.records.GetSession = nil
	imit.lockGetSession.Unlock()

	return records
}

// GetSession Imitation
func (imit *ImitSpannerServer) GetSession(P0 context.Context, P1 *spanner.GetSessionRequest) (*spanner.Session, error) {
	if len(imit.mocks.GetSession) == 0 {
		panic("ImitSpannerServer.GetSession mocks is nil but ImitSpannerServer.GetSession was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.GetSessionRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockGetSession.Lock()
	// record call
	imit.records.GetSession = append(imit.records.GetSession, call)
	// pop mock
	mockFunc := imit.mocks.GetSession[0]
	imit.mocks.GetSession = imit.mocks.GetSession[1:]
	imit.lockGetSession.Unlock()

	return mockFunc(P0, P1)
}

// ListSessions Enqueue Mock
func (imit *ImitSpannerServer) EnqueueListSessionsMock(m func(P0 context.Context, P1 *spanner.ListSessionsRequest) (*spanner.ListSessionsResponse, error)) {
	imit.lockListSessions.Lock()
	imit.mocks.ListSessions = append(imit.mocks.ListSessions, m)
	imit.lockListSessions.Unlock()
}

// ListSessions Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueListSessionsMocks(ms []func(P0 context.Context, P1 *spanner.ListSessionsRequest) (*spanner.ListSessionsResponse, error)) {
	imit.lockListSessions.Lock()
	imit.mocks.ListSessions = append(imit.mocks.ListSessions, ms...)
	imit.lockListSessions.Unlock()
}

// ListSessions Take Records
func (imit *ImitSpannerServer) TakeListSessionsRecords() []struct {
	P0 context.Context
	P1 *spanner.ListSessionsRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.ListSessionsRequest
	}

	// clear records
	imit.lockListSessions.Lock()
	records = imit.records.ListSessions
	imit.records.ListSessions = nil
	imit.lockListSessions.Unlock()

	return records
}

// ListSessions Imitation
func (imit *ImitSpannerServer) ListSessions(P0 context.Context, P1 *spanner.ListSessionsRequest) (*spanner.ListSessionsResponse, error) {
	if len(imit.mocks.ListSessions) == 0 {
		panic("ImitSpannerServer.ListSessions mocks is nil but ImitSpannerServer.ListSessions was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.ListSessionsRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockListSessions.Lock()
	// record call
	imit.records.ListSessions = append(imit.records.ListSessions, call)
	// pop mock
	mockFunc := imit.mocks.ListSessions[0]
	imit.mocks.ListSessions = imit.mocks.ListSessions[1:]
	imit.lockListSessions.Unlock()

	return mockFunc(P0, P1)
}

// PartitionQuery Enqueue Mock
func (imit *ImitSpannerServer) EnqueuePartitionQueryMock(m func(P0 context.Context, P1 *spanner.PartitionQueryRequest) (*spanner.PartitionResponse, error)) {
	imit.lockPartitionQuery.Lock()
	imit.mocks.PartitionQuery = append(imit.mocks.PartitionQuery, m)
	imit.lockPartitionQuery.Unlock()
}

// PartitionQuery Enqueue Mocks
func (imit *ImitSpannerServer) EnqueuePartitionQueryMocks(ms []func(P0 context.Context, P1 *spanner.PartitionQueryRequest) (*spanner.PartitionResponse, error)) {
	imit.lockPartitionQuery.Lock()
	imit.mocks.PartitionQuery = append(imit.mocks.PartitionQuery, ms...)
	imit.lockPartitionQuery.Unlock()
}

// PartitionQuery Take Records
func (imit *ImitSpannerServer) TakePartitionQueryRecords() []struct {
	P0 context.Context
	P1 *spanner.PartitionQueryRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.PartitionQueryRequest
	}

	// clear records
	imit.lockPartitionQuery.Lock()
	records = imit.records.PartitionQuery
	imit.records.PartitionQuery = nil
	imit.lockPartitionQuery.Unlock()

	return records
}

// PartitionQuery Imitation
func (imit *ImitSpannerServer) PartitionQuery(P0 context.Context, P1 *spanner.PartitionQueryRequest) (*spanner.PartitionResponse, error) {
	if len(imit.mocks.PartitionQuery) == 0 {
		panic("ImitSpannerServer.PartitionQuery mocks is nil but ImitSpannerServer.PartitionQuery was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.PartitionQueryRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockPartitionQuery.Lock()
	// record call
	imit.records.PartitionQuery = append(imit.records.PartitionQuery, call)
	// pop mock
	mockFunc := imit.mocks.PartitionQuery[0]
	imit.mocks.PartitionQuery = imit.mocks.PartitionQuery[1:]
	imit.lockPartitionQuery.Unlock()

	return mockFunc(P0, P1)
}

// PartitionRead Enqueue Mock
func (imit *ImitSpannerServer) EnqueuePartitionReadMock(m func(P0 context.Context, P1 *spanner.PartitionReadRequest) (*spanner.PartitionResponse, error)) {
	imit.lockPartitionRead.Lock()
	imit.mocks.PartitionRead = append(imit.mocks.PartitionRead, m)
	imit.lockPartitionRead.Unlock()
}

// PartitionRead Enqueue Mocks
func (imit *ImitSpannerServer) EnqueuePartitionReadMocks(ms []func(P0 context.Context, P1 *spanner.PartitionReadRequest) (*spanner.PartitionResponse, error)) {
	imit.lockPartitionRead.Lock()
	imit.mocks.PartitionRead = append(imit.mocks.PartitionRead, ms...)
	imit.lockPartitionRead.Unlock()
}

// PartitionRead Take Records
func (imit *ImitSpannerServer) TakePartitionReadRecords() []struct {
	P0 context.Context
	P1 *spanner.PartitionReadRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.PartitionReadRequest
	}

	// clear records
	imit.lockPartitionRead.Lock()
	records = imit.records.PartitionRead
	imit.records.PartitionRead = nil
	imit.lockPartitionRead.Unlock()

	return records
}

// PartitionRead Imitation
func (imit *ImitSpannerServer) PartitionRead(P0 context.Context, P1 *spanner.PartitionReadRequest) (*spanner.PartitionResponse, error) {
	if len(imit.mocks.PartitionRead) == 0 {
		panic("ImitSpannerServer.PartitionRead mocks is nil but ImitSpannerServer.PartitionRead was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.PartitionReadRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockPartitionRead.Lock()
	// record call
	imit.records.PartitionRead = append(imit.records.PartitionRead, call)
	// pop mock
	mockFunc := imit.mocks.PartitionRead[0]
	imit.mocks.PartitionRead = imit.mocks.PartitionRead[1:]
	imit.lockPartitionRead.Unlock()

	return mockFunc(P0, P1)
}

// Read Enqueue Mock
func (imit *ImitSpannerServer) EnqueueReadMock(m func(P0 context.Context, P1 *spanner.ReadRequest) (*spanner.ResultSet, error)) {
	imit.lockRead.Lock()
	imit.mocks.Read = append(imit.mocks.Read, m)
	imit.lockRead.Unlock()
}

// Read Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueReadMocks(ms []func(P0 context.Context, P1 *spanner.ReadRequest) (*spanner.ResultSet, error)) {
	imit.lockRead.Lock()
	imit.mocks.Read = append(imit.mocks.Read, ms...)
	imit.lockRead.Unlock()
}

// Read Take Records
func (imit *ImitSpannerServer) TakeReadRecords() []struct {
	P0 context.Context
	P1 *spanner.ReadRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.ReadRequest
	}

	// clear records
	imit.lockRead.Lock()
	records = imit.records.Read
	imit.records.Read = nil
	imit.lockRead.Unlock()

	return records
}

// Read Imitation
func (imit *ImitSpannerServer) Read(P0 context.Context, P1 *spanner.ReadRequest) (*spanner.ResultSet, error) {
	if len(imit.mocks.Read) == 0 {
		panic("ImitSpannerServer.Read mocks is nil but ImitSpannerServer.Read was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.ReadRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockRead.Lock()
	// record call
	imit.records.Read = append(imit.records.Read, call)
	// pop mock
	mockFunc := imit.mocks.Read[0]
	imit.mocks.Read = imit.mocks.Read[1:]
	imit.lockRead.Unlock()

	return mockFunc(P0, P1)
}

// Rollback Enqueue Mock
func (imit *ImitSpannerServer) EnqueueRollbackMock(m func(P0 context.Context, P1 *spanner.RollbackRequest) (*emptypb.Empty, error)) {
	imit.lockRollback.Lock()
	imit.mocks.Rollback = append(imit.mocks.Rollback, m)
	imit.lockRollback.Unlock()
}

// Rollback Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueRollbackMocks(ms []func(P0 context.Context, P1 *spanner.RollbackRequest) (*emptypb.Empty, error)) {
	imit.lockRollback.Lock()
	imit.mocks.Rollback = append(imit.mocks.Rollback, ms...)
	imit.lockRollback.Unlock()
}

// Rollback Take Records
func (imit *ImitSpannerServer) TakeRollbackRecords() []struct {
	P0 context.Context
	P1 *spanner.RollbackRequest
} {
	var records []struct {
		P0 context.Context
		P1 *spanner.RollbackRequest
	}

	// clear records
	imit.lockRollback.Lock()
	records = imit.records.Rollback
	imit.records.Rollback = nil
	imit.lockRollback.Unlock()

	return records
}

// Rollback Imitation
func (imit *ImitSpannerServer) Rollback(P0 context.Context, P1 *spanner.RollbackRequest) (*emptypb.Empty, error) {
	if len(imit.mocks.Rollback) == 0 {
		panic("ImitSpannerServer.Rollback mocks is nil but ImitSpannerServer.Rollback was just called")
	}

	call := struct {
		P0 context.Context
		P1 *spanner.RollbackRequest
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockRollback.Lock()
	// record call
	imit.records.Rollback = append(imit.records.Rollback, call)
	// pop mock
	mockFunc := imit.mocks.Rollback[0]
	imit.mocks.Rollback = imit.mocks.Rollback[1:]
	imit.lockRollback.Unlock()

	return mockFunc(P0, P1)
}

// StreamingRead Enqueue Mock
func (imit *ImitSpannerServer) EnqueueStreamingReadMock(m func(P0 *spanner.ReadRequest, P1 spanner.Spanner_StreamingReadServer) error) {
	imit.lockStreamingRead.Lock()
	imit.mocks.StreamingRead = append(imit.mocks.StreamingRead, m)
	imit.lockStreamingRead.Unlock()
}

// StreamingRead Enqueue Mocks
func (imit *ImitSpannerServer) EnqueueStreamingReadMocks(ms []func(P0 *spanner.ReadRequest, P1 spanner.Spanner_StreamingReadServer) error) {
	imit.lockStreamingRead.Lock()
	imit.mocks.StreamingRead = append(imit.mocks.StreamingRead, ms...)
	imit.lockStreamingRead.Unlock()
}

// StreamingRead Take Records
func (imit *ImitSpannerServer) TakeStreamingReadRecords() []struct {
	P0 *spanner.ReadRequest
	P1 spanner.Spanner_StreamingReadServer
} {
	var records []struct {
		P0 *spanner.ReadRequest
		P1 spanner.Spanner_StreamingReadServer
	}

	// clear records
	imit.lockStreamingRead.Lock()
	records = imit.records.StreamingRead
	imit.records.StreamingRead = nil
	imit.lockStreamingRead.Unlock()

	return records
}

// StreamingRead Imitation
func (imit *ImitSpannerServer) StreamingRead(P0 *spanner.ReadRequest, P1 spanner.Spanner_StreamingReadServer) error {
	if len(imit.mocks.StreamingRead) == 0 {
		panic("ImitSpannerServer.StreamingRead mocks is nil but ImitSpannerServer.StreamingRead was just called")
	}

	call := struct {
		P0 *spanner.ReadRequest
		P1 spanner.Spanner_StreamingReadServer
	}{
		P0: P0,
		P1: P1,
	}

	imit.lockStreamingRead.Lock()
	// record call
	imit.records.StreamingRead = append(imit.records.StreamingRead, call)
	// pop mock
	mockFunc := imit.mocks.StreamingRead[0]
	imit.mocks.StreamingRead = imit.mocks.StreamingRead[1:]
	imit.lockStreamingRead.Unlock()

	return mockFunc(P0, P1)
}
