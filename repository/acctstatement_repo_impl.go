package repository

import (
	"batch-acctstatement/entity"
	"context"
	"log"
	"time"

	"github.com/arangodb/go-driver"
)

type acctStatementRepository struct {
	Live           driver.Database
	LiveCollection string
}

func getContext(timeout int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
}

func NewAcctStatementRepository(clientLive driver.Client, dbName string) AcctStatementRepository {
	var dbLive, err = clientLive.Database(context.Background(), dbName)
	if err != nil {
		log.Fatal(err)
	}

	return &acctStatementRepository{
		Live:           dbLive,
		LiveCollection: "account_statement",
	}
}

func (db *acctStatementRepository) Create(model *entity.AcctStatement) error {
	var ctx, cancel = getContext(15)
	defer cancel()

	col, err := db.Live.Collection(ctx, db.LiveCollection)
	if err != nil {
		log.Println("error, cause: ", err)
		return err
	}
	meta, err := col.CreateDocument(ctx, model)
	if err != nil {
		log.Printf("Error while creating document, cause: %+v\n", err)
		return err
	}

	log.Printf("Created document with key '%s', revision '%s'\n", meta.Key, meta.Rev)
	return nil
}
