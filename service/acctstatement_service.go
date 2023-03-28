package service

import (
	"batch-acctstatement/entity"
	"batch-acctstatement/pkg/acctstatement"
	"sync"
)

type (
	AcctStatementService interface {
		Process(transactions []acctstatement.Transactions, acctNo string, jobs chan<- *entity.AcctStatement, wg *sync.WaitGroup)
		ProcessWork(jobs <-chan *entity.AcctStatement, wg *sync.WaitGroup)
	}
)
