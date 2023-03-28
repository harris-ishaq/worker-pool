package repository

import "batch-acctstatement/entity"

type (
	AcctStatementRepository interface {
		Create(model *entity.AcctStatement) error
	}
)
