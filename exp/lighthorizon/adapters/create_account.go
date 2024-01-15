package adapters

import (
	"github.com/stellar/go/amount"
	"github.com/stellar/go/exp/lighthorizon/common"
	"github.com/stellar/go/protocols/horizon/operations"
)

func populateCreateAccountOperation(op *common.Operation, baseOp operations.Base) (operations.CreateAccount, error) {
	createAccount := op.Get().Body.MustCreateAccountOp()

	return operations.CreateAccount{
		Base:            baseOp,
		StartingBalance: amount.String(createAccount.StartingBalance),
		Funder:          op.SourceAccount().Address(),
		Account:         createAccount.Destination.Address(),
	}, nil
}
