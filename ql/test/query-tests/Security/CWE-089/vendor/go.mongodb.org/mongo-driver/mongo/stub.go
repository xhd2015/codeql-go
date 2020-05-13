// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for go.mongodb.org/mongo-driver/mongo, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: go.mongodb.org/mongo-driver/mongo (exports: Pipeline; functions: Connect)

// Package mongo is a stub of go.mongodb.org/mongo-driver/mongo, generated by depstubber.
package mongo

import (
	context "context"
	time "time"
)

type BulkWriteResult struct {
	InsertedCount int64
	MatchedCount  int64
	ModifiedCount int64
	DeletedCount  int64
	UpsertedCount int64
	UpsertedIDs   map[int64]interface{}
}

type ChangeStream struct {
	Current interface{}
}

func (_ *ChangeStream) Close(_ context.Context) error {
	return nil
}

func (_ *ChangeStream) Decode(_ interface{}) error {
	return nil
}

func (_ *ChangeStream) Err() error {
	return nil
}

func (_ *ChangeStream) ID() int64 {
	return 0
}

func (_ *ChangeStream) Next(_ context.Context) bool {
	return false
}

func (_ *ChangeStream) ResumeToken() interface{} {
	return nil
}

func (_ *ChangeStream) TryNext(_ context.Context) bool {
	return false
}

type Client struct{}

func (_ *Client) Connect(_ context.Context) error {
	return nil
}

func (_ *Client) Database(_ string, _ ...*interface{}) *Database {
	return nil
}

func (_ *Client) Disconnect(_ context.Context) error {
	return nil
}

func (_ *Client) ListDatabaseNames(_ context.Context, _ interface{}, _ ...*interface{}) ([]string, error) {
	return nil, nil
}

func (_ *Client) ListDatabases(_ context.Context, _ interface{}, _ ...*interface{}) (ListDatabasesResult, error) {
	return ListDatabasesResult{}, nil
}

func (_ *Client) NumberSessionsInProgress() int {
	return 0
}

func (_ *Client) Ping(_ context.Context, _ *interface{}) error {
	return nil
}

func (_ *Client) StartSession(_ ...*interface{}) (Session, error) {
	return nil, nil
}

func (_ *Client) UseSession(_ context.Context, _ func(SessionContext) error) error {
	return nil
}

func (_ *Client) UseSessionWithOptions(_ context.Context, _ *interface{}, _ func(SessionContext) error) error {
	return nil
}

func (_ *Client) Watch(_ context.Context, _ interface{}, _ ...*interface{}) (*ChangeStream, error) {
	return nil, nil
}

type Collection struct{}

func (_ *Collection) Aggregate(_ context.Context, _ interface{}, _ ...*interface{}) (*Cursor, error) {
	return nil, nil
}

func (_ *Collection) BulkWrite(_ context.Context, _ []WriteModel, _ ...*interface{}) (*BulkWriteResult, error) {
	return nil, nil
}

func (_ *Collection) Clone(_ ...*interface{}) (*Collection, error) {
	return nil, nil
}

func (_ *Collection) CountDocuments(_ context.Context, _ interface{}, _ ...*interface{}) (int64, error) {
	return 0, nil
}

func (_ *Collection) Database() *Database {
	return nil
}

func (_ *Collection) DeleteMany(_ context.Context, _ interface{}, _ ...*interface{}) (*DeleteResult, error) {
	return nil, nil
}

func (_ *Collection) DeleteOne(_ context.Context, _ interface{}, _ ...*interface{}) (*DeleteResult, error) {
	return nil, nil
}

func (_ *Collection) Distinct(_ context.Context, _ string, _ interface{}, _ ...*interface{}) ([]interface{}, error) {
	return nil, nil
}

func (_ *Collection) Drop(_ context.Context) error {
	return nil
}

func (_ *Collection) EstimatedDocumentCount(_ context.Context, _ ...*interface{}) (int64, error) {
	return 0, nil
}

func (_ *Collection) Find(_ context.Context, _ interface{}, _ ...*interface{}) (*Cursor, error) {
	return nil, nil
}

func (_ *Collection) FindOne(_ context.Context, _ interface{}, _ ...*interface{}) *SingleResult {
	return nil
}

func (_ *Collection) FindOneAndDelete(_ context.Context, _ interface{}, _ ...*interface{}) *SingleResult {
	return nil
}

func (_ *Collection) FindOneAndReplace(_ context.Context, _ interface{}, _ interface{}, _ ...*interface{}) *SingleResult {
	return nil
}

func (_ *Collection) FindOneAndUpdate(_ context.Context, _ interface{}, _ interface{}, _ ...*interface{}) *SingleResult {
	return nil
}

func (_ *Collection) Indexes() IndexView {
	return IndexView{}
}

func (_ *Collection) InsertMany(_ context.Context, _ []interface{}, _ ...*interface{}) (*InsertManyResult, error) {
	return nil, nil
}

func (_ *Collection) InsertOne(_ context.Context, _ interface{}, _ ...*interface{}) (*InsertOneResult, error) {
	return nil, nil
}

func (_ *Collection) Name() string {
	return ""
}

func (_ *Collection) ReplaceOne(_ context.Context, _ interface{}, _ interface{}, _ ...*interface{}) (*UpdateResult, error) {
	return nil, nil
}

func (_ *Collection) UpdateMany(_ context.Context, _ interface{}, _ interface{}, _ ...*interface{}) (*UpdateResult, error) {
	return nil, nil
}

func (_ *Collection) UpdateOne(_ context.Context, _ interface{}, _ interface{}, _ ...*interface{}) (*UpdateResult, error) {
	return nil, nil
}

func (_ *Collection) Watch(_ context.Context, _ interface{}, _ ...*interface{}) (*ChangeStream, error) {
	return nil, nil
}

func Connect(_ context.Context, _ ...interface{}) (*Client, error) {
	return nil, nil
}

type Cursor struct {
	Current interface{}
}

func (_ *Cursor) All(_ context.Context, _ interface{}) error {
	return nil
}

func (_ *Cursor) Close(_ context.Context) error {
	return nil
}

func (_ *Cursor) Decode(_ interface{}) error {
	return nil
}

func (_ *Cursor) Err() error {
	return nil
}

func (_ *Cursor) ID() int64 {
	return 0
}

func (_ *Cursor) Next(_ context.Context) bool {
	return false
}

func (_ *Cursor) TryNext(_ context.Context) bool {
	return false
}

type Database struct{}

func (_ *Database) Aggregate(_ context.Context, _ interface{}, _ ...*interface{}) (*Cursor, error) {
	return nil, nil
}

func (_ *Database) Client() *Client {
	return nil
}

func (_ *Database) Collection(_ string, _ ...*interface{}) *Collection {
	return nil
}

func (_ *Database) Drop(_ context.Context) error {
	return nil
}

func (_ *Database) ListCollectionNames(_ context.Context, _ interface{}, _ ...*interface{}) ([]string, error) {
	return nil, nil
}

func (_ *Database) ListCollections(_ context.Context, _ interface{}, _ ...*interface{}) (*Cursor, error) {
	return nil, nil
}

func (_ *Database) Name() string {
	return ""
}

func (_ *Database) ReadConcern() *interface{} {
	return nil
}

func (_ *Database) ReadPreference() *interface{} {
	return nil
}

func (_ *Database) RunCommand(_ context.Context, _ interface{}, _ ...*interface{}) *SingleResult {
	return nil
}

func (_ *Database) RunCommandCursor(_ context.Context, _ interface{}, _ ...*interface{}) (*Cursor, error) {
	return nil, nil
}

func (_ *Database) Watch(_ context.Context, _ interface{}, _ ...*interface{}) (*ChangeStream, error) {
	return nil, nil
}

func (_ *Database) WriteConcern() *interface{} {
	return nil
}

type DatabaseSpecification struct {
	Name       string
	SizeOnDisk int64
	Empty      bool
}

type DeleteResult struct {
	DeletedCount int64
}

type IndexModel struct {
	Keys    interface{}
	Options *interface{}
}

type IndexView struct{}

func (_ IndexView) CreateMany(_ context.Context, _ []IndexModel, _ ...*interface{}) ([]string, error) {
	return nil, nil
}

func (_ IndexView) CreateOne(_ context.Context, _ IndexModel, _ ...*interface{}) (string, error) {
	return "", nil
}

func (_ IndexView) DropAll(_ context.Context, _ ...*interface{}) (interface{}, error) {
	return nil, nil
}

func (_ IndexView) DropOne(_ context.Context, _ string, _ ...*interface{}) (interface{}, error) {
	return nil, nil
}

func (_ IndexView) List(_ context.Context, _ ...*interface{}) (*Cursor, error) {
	return nil, nil
}

type InsertManyResult struct {
	InsertedIDs []interface{}
}

type InsertOneResult struct {
	InsertedID interface{}
}

type ListDatabasesResult struct {
	Databases []DatabaseSpecification
	TotalSize int64
}

type Pipeline []interface{}

type Session interface {
	AbortTransaction(_ context.Context) error
	AdvanceClusterTime(_ interface{}) error
	AdvanceOperationTime(_ *interface{}) error
	Client() *Client
	ClusterTime() interface{}
	CommitTransaction(_ context.Context) error
	EndSession(_ context.Context)
	OperationTime() *interface{}
	StartTransaction(_ ...*interface{}) error
	WithTransaction(_ context.Context, _ func(SessionContext) (interface{}, error), _ ...*interface{}) (interface{}, error)
}

type SessionContext interface {
	AbortTransaction(_ context.Context) error
	AdvanceClusterTime(_ interface{}) error
	AdvanceOperationTime(_ *interface{}) error
	Client() *Client
	ClusterTime() interface{}
	CommitTransaction(_ context.Context) error
	Deadline() (time.Time, bool)
	Done() <-chan struct{}
	EndSession(_ context.Context)
	Err() error
	OperationTime() *interface{}
	StartTransaction(_ ...*interface{}) error
	Value(_ interface{}) interface{}
	WithTransaction(_ context.Context, _ func(SessionContext) (interface{}, error), _ ...*interface{}) (interface{}, error)
}

type SingleResult struct{}

func (_ *SingleResult) Decode(_ interface{}) error {
	return nil
}

func (_ *SingleResult) DecodeBytes() (interface{}, error) {
	return nil, nil
}

func (_ *SingleResult) Err() error {
	return nil
}

type UpdateResult struct {
	MatchedCount  int64
	ModifiedCount int64
	UpsertedCount int64
	UpsertedID    interface{}
}

func (_ *UpdateResult) UnmarshalBSON(_ []byte) error {
	return nil
}

type WriteModel interface{}
