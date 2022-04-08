package forwarders

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/sqlx"
)

// Config encompasses config used by fwdmgr
//go:generate mockery --recursive --name Config --output ./mocks/ --case=underscore --structname Config --filename config.go
type Config interface {
	EvmUseForwarders() bool
	LogSQL() bool
}

type FwdMgr struct {
	logger          logger.Logger
	chainID         big.Int
	ORM             ORM
	forwarderMethod abi.Method
}

var fwdrMethod = makeFwdrMethod()

func NewFwdMgr(db *sqlx.DB, cfg Config, chainId big.Int, lggr logger.Logger) *FwdMgr {
	lggr = lggr.Named("FwdMgr")
	lggr.Infow("Initializing EVM forwarder manager")

	fwdMgr := FwdMgr{
		logger:  lggr,
		chainID: chainId,
		ORM:     NewORM(db, lggr, cfg),
	}
	return &fwdMgr
}

func (fm *FwdMgr) MakeForwardedPayload(to common.Address, EncodedPayload []byte) (fwdPayload []byte) {
	fwdPayload = append(fwdrMethod.ID, to.Bytes()...)
	fwdPayload = append(fwdPayload, EncodedPayload...)
	return fwdPayload
}

func makeFwdrMethod() abi.Method {
	addrTyp, _ := abi.NewType("address", "", nil)
	bytesTyp, _ := abi.NewType("bytes", "", nil)

	args := abi.Arguments{
		abi.Argument{
			Name: "to",
			Type: addrTyp,
		},
		abi.Argument{
			Name: "data",
			Type: bytesTyp,
		},
	}

	return abi.NewMethod("forward", "forward", abi.Function, "", false, false, args, nil)
}
