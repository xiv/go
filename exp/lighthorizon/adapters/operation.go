package adapters

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/stellar/go/exp/lighthorizon/common"
	"github.com/stellar/go/protocols/horizon/operations"
	"github.com/stellar/go/support/render/hal"
	"github.com/stellar/go/xdr"
)

func PopulateOperation(r *http.Request, op *common.Operation) (operations.Operation, error) {
	hash, err := op.TransactionHash()
	if err != nil {
		return nil, err
	}

	toid := strconv.FormatInt(op.TOID(), 10)
	baseOp := operations.Base{
		ID:                    toid,
		PT:                    toid,
		TransactionSuccessful: op.TransactionResult.Successful(),
		SourceAccount:         op.SourceAccount().Address(),
		LedgerCloseTime:       time.Unix(int64(op.LedgerHeader.ScpValue.CloseTime), 0).UTC(),
		TransactionHash:       hash,
		Type:                  operations.TypeNames[op.Get().Body.Type],
		TypeI:                 int32(op.Get().Body.Type),
	}

	lb := hal.LinkBuilder{Base: r.URL}
	self := fmt.Sprintf("/operations/%s", toid)
	baseOp.Links.Self = lb.Link(self)
	baseOp.Links.Succeeds = lb.Linkf("/effects?order=desc&cursor=%s", baseOp.PT)
	baseOp.Links.Precedes = lb.Linkf("/effects?order=asc&cursor=%s", baseOp.PT)
	baseOp.Links.Transaction = lb.Linkf("/transactions/%s", hash)
	baseOp.Links.Effects = lb.Link(self, "effects")

	switch op.Get().Body.Type {
	case xdr.OperationTypeCreateAccount:
		return populateCreateAccountOperation(op, baseOp)
	case xdr.OperationTypePayment:
		return populatePaymentOperation(op, baseOp)
	case xdr.OperationTypePathPaymentStrictReceive:
		return populatePathPaymentStrictReceiveOperation(op, baseOp)
	case xdr.OperationTypePathPaymentStrictSend:
		return populatePathPaymentStrictSendOperation(op, baseOp)
	case xdr.OperationTypeManageBuyOffer:
		return populateManageBuyOfferOperation(op, baseOp)
	case xdr.OperationTypeManageSellOffer:
		return populateManageSellOfferOperation(op, baseOp)
	case xdr.OperationTypeCreatePassiveSellOffer:
		return populateCreatePassiveSellOfferOperation(op, baseOp)
	case xdr.OperationTypeSetOptions:
		return populateSetOptionsOperation(op, baseOp)
	case xdr.OperationTypeChangeTrust:
		return populateChangeTrustOperation(op, baseOp)
	case xdr.OperationTypeAllowTrust:
		return populateAllowTrustOperation(op, baseOp)
	case xdr.OperationTypeAccountMerge:
		return populateAccountMergeOperation(op, baseOp)
	case xdr.OperationTypeInflation:
		return populateInflationOperation(op, baseOp)
	case xdr.OperationTypeManageData:
		return populateManageDataOperation(op, baseOp)
	case xdr.OperationTypeBumpSequence:
		return populateBumpSequenceOperation(op, baseOp)
	case xdr.OperationTypeCreateClaimableBalance:
		return populateCreateClaimableBalanceOperation(op, baseOp)
	case xdr.OperationTypeClaimClaimableBalance:
		return populateClaimClaimableBalanceOperation(op, baseOp)
	case xdr.OperationTypeBeginSponsoringFutureReserves:
		return populateBeginSponsoringFutureReservesOperation(op, baseOp)
	case xdr.OperationTypeEndSponsoringFutureReserves:
		return populateEndSponsoringFutureReservesOperation(op, baseOp)
	case xdr.OperationTypeRevokeSponsorship:
		return populateRevokeSponsorshipOperation(op, baseOp)
	case xdr.OperationTypeClawback:
		return populateClawbackOperation(op, baseOp)
	case xdr.OperationTypeClawbackClaimableBalance:
		return populateClawbackClaimableBalanceOperation(op, baseOp)
	case xdr.OperationTypeSetTrustLineFlags:
		return populateSetTrustLineFlagsOperation(op, baseOp)
	case xdr.OperationTypeLiquidityPoolDeposit:
		return populateLiquidityPoolDepositOperation(op, baseOp)
	case xdr.OperationTypeLiquidityPoolWithdraw:
		return populateLiquidityPoolWithdrawOperation(op, baseOp)
	default:
		return nil, fmt.Errorf("unknown operation type: %s", op.Get().Body.Type)
	}
}
