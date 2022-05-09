package nexproto

import (
	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/plogger"
)

var logger = plogger.NewLogger()

func init() {
	nex.RegisterDataHolderType(NewNintendoCreateAccountData())
	nex.RegisterDataHolderType(NewAccountExtraInfo())
}
