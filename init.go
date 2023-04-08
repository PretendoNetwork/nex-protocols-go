package nexproto

import (
	"github.com/PretendoNetwork/nex-go"
	account_management "github.com/PretendoNetwork/nex-protocols-go/account-management"
	"github.com/PretendoNetwork/nex-protocols-go/authentication"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
	"github.com/PretendoNetwork/plogger-go"
)

var logger = plogger.NewLogger()

func init() {
	nex.RegisterDataHolderType(account_management.NewNintendoCreateAccountData())
	nex.RegisterDataHolderType(authentication.NewNintendoLoginData())
	nex.RegisterDataHolderType(account_management.NewAccountExtraInfo())
	nex.RegisterDataHolderType(match_making.NewGathering())
}
