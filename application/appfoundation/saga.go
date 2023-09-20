package appfoundation

import "github.com/insolar/insolar/insolar"

type SagaAcceptInfo struct {
	Amount     string
	FromMember insolar.Reference
	Request    insolar.Reference
}
