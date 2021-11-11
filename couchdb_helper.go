package couchdb_helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	_ "github.com/go-kivik/couchdb/v4" // The CouchDB driver
	kivik "github.com/go-kivik/kivik/v4"
	logging "github.com/lordtor/go-logging"
)

var (
	log = logging.Log
)

func PutRowDB(db *kivik.DB, id string, row interface{}) (rev string, err error) {
	rev, err = db.Put(context.TODO(), id, row)
	if err != nil {
		log.Error("PutRowDB", err)
		return "", err
	}
	return rev, nil
}

func GetStats(con CouchDB) (*kivik.DBStats, error) {
	db, err := CreateClientDB(con)
	if err != nil {
		log.Panic(err)
	}
	stat, err := db.Stats(context.TODO())
	if err != nil {
		return nil, err
	}
	db.Close(context.TODO())
	return stat, nil
}
func GenerateAddress(con CouchDB) string {
	var address string
	if con.Password != "" && con.Username != "" && con.Port != "" {
		address = fmt.Sprintf("%s://%v:%v@%v:%v/", con.Schema, con.Username,
			con.Password,
			con.Server,
			con.Port)
	} else if con.Password != "" && con.Username != "" && con.Port == "" {
		address = fmt.Sprintf("%s://%v:%v@%v/", con.Schema, con.Username,
			con.Password,
			con.Server)
	} else {
		address = fmt.Sprintf("%s://%v:%v/", con.Schema,
			con.Server,
			con.Port)
	}
	return address
}
func CreateConnection(con CouchDB) (*kivik.Client, error) {
	address := GenerateAddress(con)
	client, err := kivik.New("couch", address)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func CreateClientDB(con CouchDB) (*kivik.DB, error) {
	client, err := CreateConnection(con)
	if err != nil {
		return nil, err
	}
	db := client.DB(con.Name)
	return db, nil
}
func CreateDBIndex(con CouchDB, indexName string) (bool, error) {
	indexJSON := map[string]interface{}{}
	indexJSON["fields"] = []string{indexName}
	db, err := CreateClientDB(con)
	if err != nil {
		log.Panic(err)
	}
	err = db.CreateIndex(context.TODO(), "", "", indexJSON)
	if err != nil {
		log.Error("CreateDBIndex1", err)
		db.Close(context.TODO())
		return false, err
	}
	db.Close(context.TODO())
	return true, nil
}
func CreateDBManyIndex(con CouchDB, indexString []string) (bool, error) {
	indexJSON := map[string]interface{}{}
	indexJSON["fields"] = indexString
	db, err := CreateClientDB(con)
	if err != nil {
		log.Panic(err)
	}
	err = db.CreateIndex(context.TODO(), "", "", indexJSON)
	if err != nil {
		log.Error("CreateDBIndex1", err)
		db.Close(context.TODO())
		return false, err
	}
	db.Close(context.TODO())
	return true, nil
}
func CheckDB(con CouchDB, create bool) (bool, error) {
	client, err := CreateConnection(con)
	if err != nil {
		return false, err
	}
	exist, err := client.DBExists(context.TODO(), con.Name)
	if err != nil {
		return false, err
	}
	if exist {
		log.Debugf("CouchDB: Database: %v exist", con.Name)
		client.Close(context.TODO())
		return true, nil
	} else if !exist && create {
		log.Warningf("CouchDB: Database not found try create new: %v\n", con.Name)
		err := client.CreateDB(context.TODO(), con.Name)
		if err != nil {
			return false, err
		}
		client.Close(context.TODO())
		return true, nil
	} else {
		client.Close(context.TODO())
		return false, errors.New("CouchDB: Database not found!")
	}
}

func GetRowDB(db *kivik.DB, id string) *kivik.Row {
	row := db.Get(context.TODO(), id)
	return row
}
func FindGetRow(db *kivik.DB, queryFields map[string]interface{}) ([]interface{}, int, string, error) {
	rows, err := db.Find(context.TODO(), queryFields)
	if err != nil {
		return nil, 0, "", err
	}
	docs := []interface{}{}
	for rows.Next() {
		var doc interface{}
		if err := rows.ScanDoc(&doc); err != nil {
			return nil, 0, "", err
		}
		docs = append(docs, doc)
	}
	return docs, len(docs), rows.Bookmark(), nil
}

func FindRow(db *kivik.DB, queryFields map[string]string, bookmark string, fields []string, limit int) ([]interface{}, int, string, error) {
	querySelector := map[string]interface{}{
		"selector": queryFields,
		"fields":   []string{},
		"limit":    25,
		//"limit":    1000000000,
	}
	if bookmark != "" {
		querySelector["bookmark"] = bookmark
	}
	if len(fields) > 0 {
		querySelector["fields"] = fields
	}
	if limit > 0 {
		querySelector["limit"] = limit
	}

	query, err := json.Marshal(querySelector)
	if err != nil {
		return nil, 0, "", err
	}
	rows, err := db.Find(context.TODO(), query)

	if err != nil {
		return nil, 0, "", err
	}
	docs := []interface{}{}
	for rows.Next() {

		var doc interface{}
		if err := rows.ScanDoc(&doc); err != nil {
			return nil, 0, "", err
		}
		docs = append(docs, doc)
	}

	if rows.TotalRows() > 0 {
		log.Info("FindRow", rows.TotalRows())
	}
	return docs, len(docs), rows.Bookmark(), nil
}
