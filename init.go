// Package nexproto provides all the main NEX protocols.
//
// Each folder contains a different package for that specific protocol,
// with all their types and methods needed to parse and build packets with RMC payloads
package nexproto

import (
	"github.com/PretendoNetwork/nex-go"
	account_management_types "github.com/PretendoNetwork/nex-protocols-go/account-management/types"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
	ticket_granting_types "github.com/PretendoNetwork/nex-protocols-go/ticket-granting/types"
)

func init() {
	nex.RegisterDataHolderType("NintendoCreateAccountData", account_management_types.NewNintendoCreateAccountData())
	nex.RegisterDataHolderType("AccountExtraInfo", account_management_types.NewAccountExtraInfo())
	nex.RegisterDataHolderType("NintendoLoginData", ticket_granting_types.NewNintendoLoginData())
	nex.RegisterDataHolderType("AuthenticationInfo", ticket_granting_types.NewAuthenticationInfo())
	nex.RegisterDataHolderType("Gathering", match_making_types.NewGathering())
	nex.RegisterDataHolderType("MatchmakeSession", match_making_types.NewMatchmakeSession())
}
