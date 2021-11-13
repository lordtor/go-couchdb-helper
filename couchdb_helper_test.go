package couchdb_helper

import (
	"errors"
	"reflect"
	"testing"

	_ "github.com/go-kivik/couchdb/v4"
	kivik "github.com/go-kivik/kivik/v4"
	"github.com/go-kivik/kivik/v4/driver"
	"github.com/go-kivik/kivikmock/v4"
)

// TODO: Read https://github.com/go-kivik/kivikmock/blob/d4ebdf08086151ec454c0c18104688ae5d92f73e/db_test.go
func TestPutRowDB(t *testing.T) {
	c, mock, err := kivikmock.New()
	if err != nil {
		panic(err)
	}

	foo := mock.NewDB()
	bar := mock.NewDB()
	mock.ExpectDB().WithName("foo").WillReturn(foo)
	foo.ExpectPut().WillReturn("123")
	mock.ExpectDB().WithName("bar").WillReturn(bar)
	bar.ExpectPut().WillReturnError(errors.New("foo err"))

	type args struct {
		db  *kivik.DB
		id  string
		row interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRev string
		wantErr bool
	}{
		{"1", args{c.DB("foo"), "q", map[string]string{"id": "q", "f": "ff"}}, "123", false},
		{"2", args{c.DB("bar"), "q", map[string]string{"id": "q", "f": "ff"}}, "", true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRev, err := PutRowDB(tt.args.db, tt.args.id, tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutRowDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRev != tt.wantRev {
				t.Errorf("PutRowDB() = %v, want %v", gotRev, tt.wantRev)
			}
		})
	}
}

func TestGetStats(t *testing.T) {

	type args struct {
		con CouchDB
	}
	tests := []struct {
		name    string
		args    args
		want    *kivik.DBStats
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStats(tt.args.con)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateAddress(t *testing.T) {
	type args struct {
		con CouchDB
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateAddress(tt.args.con); got != tt.want {
				t.Errorf("GenerateAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateConnection(t *testing.T) {
	type args struct {
		con CouchDB
	}
	tests := []struct {
		name    string
		args    args
		want    *kivik.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateConnection(tt.args.con)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateClientDB(t *testing.T) {
	type args struct {
		con CouchDB
	}
	tests := []struct {
		name    string
		args    args
		want    *kivik.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateClientDB(tt.args.con)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateClientDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateClientDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateDBIndex(t *testing.T) {
	type args struct {
		con       CouchDB
		indexName string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateDBIndex(tt.args.con, tt.args.indexName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDBIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateDBIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateDBManyIndex(t *testing.T) {
	type args struct {
		con         CouchDB
		indexString []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateDBManyIndex(tt.args.con, tt.args.indexString)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDBManyIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateDBManyIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckDB(t *testing.T) {
	type args struct {
		con    CouchDB
		create bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckDB(tt.args.con, tt.args.create)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRowDB(t *testing.T) {
	c, mock, err := kivikmock.New()
	if err != nil {
		panic(err)
	}

	foo := mock.NewDB()
	bar := mock.NewDB()
	mock.ExpectDB().WithName("foo").WillReturn(foo)
	foo.ExpectGet().WillReturn(&driver.Document{Rev: "2-bar"})
	mock.ExpectDB().WithName("bar").WillReturn(bar)
	bar.ExpectGet().WillReturnError(errors.New("foo err"))

	type args struct {
		db *kivik.DB
		id string
	}
	tests := []struct {
		name string
		args args
		want *kivik.Row
	}{
		{"1", args{c.DB("foo"), "q"}, &kivik.Row{Rev: "2-bar"}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRowDB(tt.args.db, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRowDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindGetRow(t *testing.T) {
	type args struct {
		db          *kivik.DB
		queryFields map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		want1   int
		want2   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := FindGetRow(tt.args.db, tt.args.queryFields)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindGetRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindGetRow() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindGetRow() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("FindGetRow() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestFindRow(t *testing.T) {
	type args struct {
		db          *kivik.DB
		queryFields map[string]string
		bookmark    string
		fields      []string
		limit       int
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		want1   int
		want2   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := FindRow(tt.args.db, tt.args.queryFields, tt.args.bookmark, tt.args.fields, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRow() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindRow() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("FindRow() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
