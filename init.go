// Package nexproto provides all the main NEX protocols.
//
// Each folder contains a different package for that specific protocol,
// with all their types and methods needed to parse and build packets with RMC payloads
package nexproto

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	account_management_types "github.com/PretendoNetwork/nex-protocols-go/v2/account-management/types"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	ticket_granting_types "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting/types"
)

func init() {
	types.RegisterObjectHolderType(account_management_types.NewNintendoCreateAccountData())
	types.RegisterObjectHolderType(account_management_types.NewAccountExtraInfo())
	types.RegisterObjectHolderType(ticket_granting_types.NewNintendoLoginData())
	types.RegisterObjectHolderType(ticket_granting_types.NewAuthenticationInfo())
	types.RegisterObjectHolderType(match_making_types.NewGathering())
	types.RegisterObjectHolderType(match_making_types.NewMatchmakeSession())
}
