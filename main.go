package main

import (
	"batch-acctstatement/entity"
	"batch-acctstatement/helpers"
	"batch-acctstatement/pkg/acctstatement"
	"batch-acctstatement/service"
	"fmt"
	"log"
	"strings"
	"sync"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// var clientLive = config.NewArangoDBDatabase(config.DBURL, config.DBUSER, config.DBPASS)
	// var repo = repository.NewAcctStatementRepository(clientLive, config.DBNAME)
	// var service = service.NewAcctStatementService(repo)

	listAcctNo := strings.Split("20230326,20230327,20230328", ",")

	jobs := make(chan *entity.AcctStatement)
	wg := new(sync.WaitGroup)
	go service.ProcessWork(jobs, wg)

	for _, acctNo := range listAcctNo {
		var listAllTransactions []acctstatement.Transactions

		for i := 0; i < 900; i++ {
			listAllTransactions = append(listAllTransactions, acctstatement.Transactions{
				PostingDate:     helpers.TimeHostNow().Format("20060102") + "-" + fmt.Sprint(i),
				ReferenceNumber: "AP-" + fmt.Sprint(i),
			})
		}

		service.Process(listAllTransactions, acctNo, jobs, wg)
	}
	close(jobs)
	// wg.Wait()

	// helpers.SendSignKill()
	return
}
