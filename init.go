// Package nexproto provides all the main NEX protocols.
//
// Each folder contains a different package for that specific protocol,
// with all their types and methods needed to parse and build packets with RMC payloads
package nexproto

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	account_management_types "github.com/PretendoNetwork/nex-protocols-go/v2/account-management/types"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/v2/messaging/types"
	nintendo_notifications_types "github.com/PretendoNetwork/nex-protocols-go/v2/nintendo-notifications/types"
	ticket_granting_types "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting/types"
)

func init() {
	types.RegisterObjectHolderType(account_management_types.NewNintendoCreateAccountData())
	types.RegisterObjectHolderType(account_management_types.NewAccountExtraInfo())
	types.RegisterObjectHolderType(ticket_granting_types.NewNintendoLoginData())
	types.RegisterObjectHolderType(ticket_granting_types.NewAuthenticationInfo())
	types.RegisterObjectHolderType(match_making_types.NewGathering())
	types.RegisterObjectHolderType(match_making_types.NewMatchmakeSession())
	types.RegisterObjectHolderType(messaging_types.NewUserMessage())
	types.RegisterObjectHolderType(messaging_types.NewTextMessage())
	types.RegisterObjectHolderType(messaging_types.NewBinaryMessage())
	types.RegisterObjectHolderType(nintendo_notifications_types.NewNintendoNotificationEventGeneral())
}
