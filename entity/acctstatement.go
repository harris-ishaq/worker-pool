package entity

type (
	AcctStatement struct {
		Key               string `json:"_key"`
		PostingDate       string `json:"postingDate"`
		EffectiveDate     string `json:"effectiveDate"`
		TransactionAmount string `json:"transactionAmount"`
		DebitCredit       string `json:"debitCredit"`
		ReferenceNumber   string `json:"referenceNumber"`
		Description       string `json:"description"`
		TransactionType   string `json:"transactionType"`
		BranchCode        string `json:"branchCode"`
	}
)
