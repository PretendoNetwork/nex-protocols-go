// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Account Management protocol
	ProtocolID = 0x19

	// MethodCreateAccount is the method ID for method CreateAccount
	MethodCreateAccount = 0x1

	// MethodDeleteAccount is the method ID for method DeleteAccount
	MethodDeleteAccount = 0x2

	// MethodDisableAccount is the method ID for method DisableAccount
	MethodDisableAccount = 0x3

	// MethodChangePassword is the method ID for method ChangePassword
	MethodChangePassword = 0x4

	// MethodTestCapability is the method ID for method TestCapability
	MethodTestCapability = 0x5

	// MethodGetName is the method ID for method GetName
	MethodGetName = 0x6

	// MethodGetAccountData is the method ID for method GetAccountData
	MethodGetAccountData = 0x7

	// MethodGetPrivateData is the method ID for method GetPrivateData
	MethodGetPrivateData = 0x8

	// MethodGetPublicData is the method ID for method GetPublicData
	MethodGetPublicData = 0x9

	// MethodGetMultiplePublicData is the method ID for method GetMultiplePublicData
	MethodGetMultiplePublicData = 0xA

	// MethodUpdateAccountName is the method ID for method UpdateAccountName
	MethodUpdateAccountName = 0xB

	// MethodUpdateAccountEmail is the method ID for method UpdateAccountEmail
	MethodUpdateAccountEmail = 0xC

	// MethodUpdateCustomData is the method ID for method UpdateCustomData
	MethodUpdateCustomData = 0xD

	// MethodFindByNameRegex is the method ID for method FindByNameRegex
	MethodFindByNameRegex = 0xE

	// MethodUpdateAccountExpiryDate is the method ID for method UpdateAccountExpiryDate
	MethodUpdateAccountExpiryDate = 0xF

	// MethodUpdateAccountEffectiveDate is the method ID for method UpdateAccountEffectiveDate
	MethodUpdateAccountEffectiveDate = 0x10

	// MethodUpdateStatus is the method ID for method UpdateStatus
	MethodUpdateStatus = 0x11

	// MethodGetStatus is the method ID for method GetStatus
	MethodGetStatus = 0x12

	// MethodGetLastConnectionStats is the method ID for method GetLastConnectionStats
	MethodGetLastConnectionStats = 0x13

	// MethodResetPassword is the method ID for method ResetPassword
	MethodResetPassword = 0x14

	// MethodCreateAccountWithCustomData is the method ID for method CreateAccountWithCustomData
	MethodCreateAccountWithCustomData = 0x15

	// MethodRetrieveAccount is the method ID for method RetrieveAccount
	MethodRetrieveAccount = 0x16

	// MethodUpdateAccount is the method ID for method UpdateAccount
	MethodUpdateAccount = 0x17

	// MethodChangePasswordByGuest is the method ID for method ChangePasswordByGuest
	MethodChangePasswordByGuest = 0x18

	// MethodFindByNameLike is the method ID for method FindByNameLike
	MethodFindByNameLike = 0x19

	// MethodCustomCreateAccount is the method ID for method CustomCreateAccount
	MethodCustomCreateAccount = 0x1A

	// MethodNintendoCreateAccount is the method ID for method NintendoCreateAccount
	MethodNintendoCreateAccount = 0x1B

	// MethodLookupOrCreateAccount is the method ID for method LookupOrCreateAccount
	MethodLookupOrCreateAccount = 0x1C

	// MethodDisconnectPrincipal is the method ID for method DisconnectPrincipal
	MethodDisconnectPrincipal = 0x1D

	// MethodDisconnectAllPrincipals is the method ID for method DisconnectAllPrincipals
	MethodDisconnectAllPrincipals = 0x1E
)

// Protocol stores all the RMC method handlers for the Account Management protocol and listens for requests
type Protocol struct {
	Server                             nex.ServerInterface
	createAccountHandler               func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string) uint32
	deleteAccountHandler               func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) uint32
	disableAccountHandler              func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtUntil *nex.DateTime, strMessage string) uint32
	changePasswordHandler              func(err error, packet nex.PacketInterface, callID uint32, strNewKey string) uint32
	testCapabilityHandler              func(err error, packet nex.PacketInterface, callID uint32, uiCapability uint32) uint32
	getNameHandler                     func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) uint32
	getAccountDataHandler              func(err error, packet nex.PacketInterface, callID uint32) uint32
	getPrivateDataHandler              func(err error, packet nex.PacketInterface, callID uint32) uint32
	getPublicDataHandler               func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) uint32
	getMultiplePublicDataHandler       func(err error, packet nex.PacketInterface, callID uint32, lstPrincipals []*nex.PID) uint32
	updateAccountNameHandler           func(err error, packet nex.PacketInterface, callID uint32, strName string) uint32
	updateAccountEmailHandler          func(err error, packet nex.PacketInterface, callID uint32, strName string) uint32
	updateCustomDataHandler            func(err error, packet nex.PacketInterface, callID uint32, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) uint32
	findByNameRegexHandler             func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strRegex string, resultRange *nex.ResultRange) uint32
	updateAccountExpiryDateHandler     func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtExpiry *nex.DateTime, strExpiredMessage string) uint32
	updateAccountEffectiveDateHandler  func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtEffectiveFrom *nex.DateTime, strNotEffectiveMessage string) uint32
	updateStatusHandler                func(err error, packet nex.PacketInterface, callID uint32, strStatus string) uint32
	getStatusHandler                   func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) uint32
	getLastConnectionStatsHandler      func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) uint32
	resetPasswordHandler               func(err error, packet nex.PacketInterface, callID uint32) uint32
	createAccountWithCustomDataHandler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) uint32
	retrieveAccountHandler             func(err error, packet nex.PacketInterface, callID uint32) uint32
	updateAccountHandler               func(err error, packet nex.PacketInterface, callID uint32, strKey string, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) uint32
	changePasswordByGuestHandler       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, strEmail string) uint32
	findByNameLikeHandler              func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strLike string, resultRange *nex.ResultRange) uint32
	customCreateAccountHandler         func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) uint32
	nintendoCreateAccountHandler       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) uint32
	lookupOrCreateAccountHandler       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) uint32
	disconnectPrincipalHandler         func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) uint32
	disconnectAllPrincipalsHandler     func(err error, packet nex.PacketInterface, callID uint32) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodCreateAccount:
		go protocol.handleCreateAccount(packet)
	case MethodDeleteAccount:
		go protocol.handleDeleteAccount(packet)
	case MethodDisableAccount:
		go protocol.handleDisableAccount(packet)
	case MethodChangePassword:
		go protocol.handleChangePassword(packet)
	case MethodTestCapability:
		go protocol.handleTestCapability(packet)
	case MethodGetName:
		go protocol.handleGetName(packet)
	case MethodGetAccountData:
		go protocol.handleGetAccountData(packet)
	case MethodGetPrivateData:
		go protocol.handleGetPrivateData(packet)
	case MethodGetPublicData:
		go protocol.handleGetPublicData(packet)
	case MethodGetMultiplePublicData:
		go protocol.handleGetMultiplePublicData(packet)
	case MethodUpdateAccountName:
		go protocol.handleUpdateAccountName(packet)
	case MethodUpdateAccountEmail:
		go protocol.handleUpdateAccountEmail(packet)
	case MethodUpdateCustomData:
		go protocol.handleUpdateCustomData(packet)
	case MethodFindByNameRegex:
		go protocol.handleFindByNameRegex(packet)
	case MethodUpdateAccountExpiryDate:
		go protocol.handleUpdateAccountExpiryDate(packet)
	case MethodUpdateAccountEffectiveDate:
		go protocol.handleUpdateAccountEffectiveDate(packet)
	case MethodUpdateStatus:
		go protocol.handleUpdateStatus(packet)
	case MethodGetStatus:
		go protocol.handleGetStatus(packet)
	case MethodGetLastConnectionStats:
		go protocol.handleGetLastConnectionStats(packet)
	case MethodResetPassword:
		go protocol.handleResetPassword(packet)
	case MethodCreateAccountWithCustomData:
		go protocol.handleCreateAccountWithCustomData(packet)
	case MethodRetrieveAccount:
		go protocol.handleRetrieveAccount(packet)
	case MethodUpdateAccount:
		go protocol.handleUpdateAccount(packet)
	case MethodChangePasswordByGuest:
		go protocol.handleChangePasswordByGuest(packet)
	case MethodFindByNameLike:
		go protocol.handleFindByNameLike(packet)
	case MethodCustomCreateAccount:
		go protocol.handleCustomCreateAccount(packet)
	case MethodNintendoCreateAccount:
		go protocol.handleNintendoCreateAccount(packet)
	case MethodLookupOrCreateAccount:
		go protocol.handleLookupOrCreateAccount(packet)
	case MethodDisconnectPrincipal:
		go protocol.handleDisconnectPrincipal(packet)
	case MethodDisconnectAllPrincipals:
		go protocol.handleDisconnectAllPrincipals(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported AccountManagement method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Account Management protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
