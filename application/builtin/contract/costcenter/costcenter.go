package costcenter

import (
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
)

type CostCenter struct {
	foundation.BaseContract
	FeeMember *insolar.Reference
}

// New creates new CostCenter.
func New(feeMember *insolar.Reference) (*CostCenter, error) {
	return &CostCenter{
		FeeMember: feeMember,
	}, nil
}

// GetFeeMember gets fee member reference.
// ins:immutable
func (cc *CostCenter) GetFeeMember() (*insolar.Reference, error) {
	return cc.FeeMember, nil
}

const Fee = "100000000"

// CalcFee calculates fee for amount. Returns fee.
// ins:immutable
func (cc *CostCenter) CalcFee(amountStr string) (string, error) {
	return Fee, nil
}
