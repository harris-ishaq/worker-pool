package service

import (
	"batch-acctstatement/entity"
	"batch-acctstatement/helpers"
	"batch-acctstatement/pkg/acctstatement"
	"log"
	"sync"
	"time"
)

// type acctStatementService struct {
// 	repoAcctStatement repository.AcctStatementRepository
// }

// type payloadPostingDate struct {
// 	PostingDate string `json:"postingDate"`
// 	PaymentID   string `json:"paymentId"`
// }

const totalWorkers = 10

// func NewAcctStatementService(acctStatementRepo repository.AcctStatementRepository) AcctStatementService {
// 	return &acctStatementService{
// 		repoAcctStatement: acctStatementRepo,
// 	}
// }

func Process(transactions []acctstatement.Transactions, acctNo string, jobs chan<- *entity.AcctStatement, wg *sync.WaitGroup) {
	var timeNow = helpers.TimeHostNow()

	for _, transaction := range transactions {
		data := entity.AcctStatement{
			Key:               acctNo + "_" + timeNow.Format("20060102"),
			PostingDate:       transaction.PostingDate,
			EffectiveDate:     transaction.EffectiveDate,
			TransactionAmount: transaction.TransactionAmount,
			DebitCredit:       transaction.DebitCredit,
			ReferenceNumber:   transaction.ReferenceNumber,
			Description:       transaction.Description,
			TransactionType:   transaction.TransactionType,
			BranchCode:        transaction.BranchCode,
		}

		wg.Add(1)
		jobs <- &data
	}

}

func ProcessWork(jobs <-chan *entity.AcctStatement, wg *sync.WaitGroup) {
	for worker := 0; worker < totalWorkers; worker++ {
		go func(jobs <-chan *entity.AcctStatement, wg *sync.WaitGroup, worker int) {
			data := <-jobs
			log.Println("worker ", worker, "started job id ", data.Key)
			time.Sleep(time.Second)
			log.Println("worker ", worker, "done working on job id ", data.Key)
			// log.Println("Data: ", data)
			// err := service.repoAcctStatement.Create(data)
			// if err != nil {
			// 	log.Println("Error while create data in collection , cause: ", err)
			// 	helpers.PubLogMsg("ERROR", fmt.Sprintf("Error while create data in collection because: %v", err))
			// }

			// pub, err := config.ConnectNats()
			// if err != nil {
			// 	log.Printf("error cause:%+v\n", err)
			// }

			// payloadData := payloadPostingDate{
			// 	PostingDate: data.PostingDate,
			// 	PaymentID:   data.ReferenceNumber,
			// }
			// payloadDataBytes, err := json.Marshal(payloadData)
			// if err != nil {
			// 	log.Printf("error cause:%+v\n", err)
			// }
			// if err := pub.Stan.Publish(config.CH_POSTINGDATE, payloadDataBytes); err != nil {
			// 	log.Printf("error cause:%+v\n", err)
			// }
			// if err := pub.Stan.Publish(config.CH_REPORTUPDATE, payloadDataBytes); err != nil {
			// 	log.Printf("error cause:%+v\n", err)
			// }

			wg.Done()
		}(jobs, wg, worker)
	}

}
