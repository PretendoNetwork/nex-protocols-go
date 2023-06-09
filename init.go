package nexproto

import (
	"github.com/PretendoNetwork/nex-go"
	account_management "github.com/PretendoNetwork/nex-protocols-go/account-management"
	"github.com/PretendoNetwork/nex-protocols-go/authentication"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
)

func init() {
	nex.RegisterDataHolderType(account_management.NewNintendoCreateAccountData())
	nex.RegisterDataHolderType(account_management.NewAccountExtraInfo())
	nex.RegisterDataHolderType(authentication.NewNintendoLoginData())
	nex.RegisterDataHolderType(authentication.NewAuthenticationInfo())
	nex.RegisterDataHolderType(match_making.NewGathering())
	nex.RegisterDataHolderType(match_making.NewMatchmakeSession())
}
