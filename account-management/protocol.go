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
	Server                      nex.ServerInterface
	CreateAccount               func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string) (*nex.RMCMessage, uint32)
	DeleteAccount               func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	DisableAccount              func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtUntil *nex.DateTime, strMessage string) (*nex.RMCMessage, uint32)
	ChangePassword              func(err error, packet nex.PacketInterface, callID uint32, strNewKey string) (*nex.RMCMessage, uint32)
	TestCapability              func(err error, packet nex.PacketInterface, callID uint32, uiCapability uint32) (*nex.RMCMessage, uint32)
	GetName                     func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	GetAccountData              func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetPrivateData              func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetPublicData               func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	GetMultiplePublicData       func(err error, packet nex.PacketInterface, callID uint32, lstPrincipals []*nex.PID) (*nex.RMCMessage, uint32)
	UpdateAccountName           func(err error, packet nex.PacketInterface, callID uint32, strName string) (*nex.RMCMessage, uint32)
	UpdateAccountEmail          func(err error, packet nex.PacketInterface, callID uint32, strName string) (*nex.RMCMessage, uint32)
	UpdateCustomData            func(err error, packet nex.PacketInterface, callID uint32, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32)
	FindByNameRegex             func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strRegex string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	UpdateAccountExpiryDate     func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtExpiry *nex.DateTime, strExpiredMessage string) (*nex.RMCMessage, uint32)
	UpdateAccountEffectiveDate  func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtEffectiveFrom *nex.DateTime, strNotEffectiveMessage string) (*nex.RMCMessage, uint32)
	UpdateStatus                func(err error, packet nex.PacketInterface, callID uint32, strStatus string) (*nex.RMCMessage, uint32)
	GetStatus                   func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	GetLastConnectionStats      func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	ResetPassword               func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	CreateAccountWithCustomData func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32)
	RetrieveAccount             func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	UpdateAccount               func(err error, packet nex.PacketInterface, callID uint32, strKey string, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32)
	ChangePasswordByGuest       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, strEmail string) (*nex.RMCMessage, uint32)
	FindByNameLike              func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strLike string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	CustomCreateAccount         func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32)
	NintendoCreateAccount       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32)
	LookupOrCreateAccount       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32)
	DisconnectPrincipal         func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	DisconnectAllPrincipals     func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodCreateAccount:
		protocol.handleCreateAccount(packet)
	case MethodDeleteAccount:
		protocol.handleDeleteAccount(packet)
	case MethodDisableAccount:
		protocol.handleDisableAccount(packet)
	case MethodChangePassword:
		protocol.handleChangePassword(packet)
	case MethodTestCapability:
		protocol.handleTestCapability(packet)
	case MethodGetName:
		protocol.handleGetName(packet)
	case MethodGetAccountData:
		protocol.handleGetAccountData(packet)
	case MethodGetPrivateData:
		protocol.handleGetPrivateData(packet)
	case MethodGetPublicData:
		protocol.handleGetPublicData(packet)
	case MethodGetMultiplePublicData:
		protocol.handleGetMultiplePublicData(packet)
	case MethodUpdateAccountName:
		protocol.handleUpdateAccountName(packet)
	case MethodUpdateAccountEmail:
		protocol.handleUpdateAccountEmail(packet)
	case MethodUpdateCustomData:
		protocol.handleUpdateCustomData(packet)
	case MethodFindByNameRegex:
		protocol.handleFindByNameRegex(packet)
	case MethodUpdateAccountExpiryDate:
		protocol.handleUpdateAccountExpiryDate(packet)
	case MethodUpdateAccountEffectiveDate:
		protocol.handleUpdateAccountEffectiveDate(packet)
	case MethodUpdateStatus:
		protocol.handleUpdateStatus(packet)
	case MethodGetStatus:
		protocol.handleGetStatus(packet)
	case MethodGetLastConnectionStats:
		protocol.handleGetLastConnectionStats(packet)
	case MethodResetPassword:
		protocol.handleResetPassword(packet)
	case MethodCreateAccountWithCustomData:
		protocol.handleCreateAccountWithCustomData(packet)
	case MethodRetrieveAccount:
		protocol.handleRetrieveAccount(packet)
	case MethodUpdateAccount:
		protocol.handleUpdateAccount(packet)
	case MethodChangePasswordByGuest:
		protocol.handleChangePasswordByGuest(packet)
	case MethodFindByNameLike:
		protocol.handleFindByNameLike(packet)
	case MethodCustomCreateAccount:
		protocol.handleCustomCreateAccount(packet)
	case MethodNintendoCreateAccount:
		protocol.handleNintendoCreateAccount(packet)
	case MethodLookupOrCreateAccount:
		protocol.handleLookupOrCreateAccount(packet)
	case MethodDisconnectPrincipal:
		protocol.handleDisconnectPrincipal(packet)
	case MethodDisconnectAllPrincipals:
		protocol.handleDisconnectAllPrincipals(packet)
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
