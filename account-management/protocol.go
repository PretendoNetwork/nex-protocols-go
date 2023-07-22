// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	account_management_types "github.com/PretendoNetwork/nex-protocols-go/account-management/types"
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

// AccountManagementProtocol handles the Account Management NEX protocol
type AccountManagementProtocol struct {
	Server                             *nex.Server
	createAccountHandler               func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string)
	deleteAccountHandler               func(err error, client *nex.Client, callID uint32, idPrincipal uint32)
	disableAccountHandler              func(err error, client *nex.Client, callID uint32, idPrincipal uint32, dtUntil *nex.DateTime, strMessage string)
	changePasswordHandler              func(err error, client *nex.Client, callID uint32, strNewKey string)
	testCapabilityHandler              func(err error, client *nex.Client, callID uint32, uiCapability uint32)
	getNameHandler                     func(err error, client *nex.Client, callID uint32, idPrincipal uint32)
	getAccountDataHandler              func(err error, client *nex.Client, callID uint32)
	getPrivateDataHandler              func(err error, client *nex.Client, callID uint32)
	getPublicDataHandler               func(err error, client *nex.Client, callID uint32, idPrincipal uint32)
	getMultiplePublicDataHandler       func(err error, client *nex.Client, callID uint32, lstPrincipals []uint32)
	updateAccountNameHandler           func(err error, client *nex.Client, callID uint32, strName string)
	updateAccountEmailHandler          func(err error, client *nex.Client, callID uint32, strName string)
	updateCustomDataHandler            func(err error, client *nex.Client, callID uint32, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder)
	findByNameRegexHandler             func(err error, client *nex.Client, callID uint32, uiGroups uint32, strRegex string, resultRange *nex.ResultRange)
	updateAccountExpiryDateHandler     func(err error, client *nex.Client, callID uint32, idPrincipal uint32, dtExpiry *nex.DateTime, strExpiredMessage string)
	updateAccountEffectiveDateHandler  func(err error, client *nex.Client, callID uint32, idPrincipal uint32, dtEffectiveFrom *nex.DateTime, strNotEffectiveMessage string)
	updateStatusHandler                func(err error, client *nex.Client, callID uint32, strStatus string)
	getStatusHandler                   func(err error, client *nex.Client, callID uint32, idPrincipal uint32)
	getLastConnectionStatsHandler      func(err error, client *nex.Client, callID uint32, idPrincipal uint32)
	resetPasswordHandler               func(err error, client *nex.Client, callID uint32)
	createAccountWithCustomDataHandler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder)
	retrieveAccountHandler             func(err error, client *nex.Client, callID uint32)
	updateAccountHandler               func(err error, client *nex.Client, callID uint32, strKey string, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder)
	changePasswordByGuestHandler       func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, strEmail string)
	findByNameLikeHandler              func(err error, client *nex.Client, callID uint32, uiGroups uint32, strLike string, resultRange *nex.ResultRange)
	customCreateAccountHandler         func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder)
	nintendoCreateAccountHandler       func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder)
	lookupOrCreateAccountHandler       func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder)
	disconnectPrincipalHandler         func(err error, client *nex.Client, callID uint32, idPrincipal uint32)
	disconnectAllPrincipalsHandler     func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *AccountManagementProtocol) Setup() {
	nex.RegisterDataHolderType(account_management_types.NewNintendoCreateAccountData())
	nex.RegisterDataHolderType(account_management_types.NewAccountExtraInfo())

	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *AccountManagementProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
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
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported AccountManagement method ID: %#v\n", request.MethodID())
	}
}

// NewAccountManagementProtocol returns a new AccountManagementProtocol
func NewAccountManagementProtocol(server *nex.Server) *AccountManagementProtocol {
	protocol := &AccountManagementProtocol{Server: server}

	protocol.Setup()

	return protocol
}
