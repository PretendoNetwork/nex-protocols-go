// Package protocol implements the Uplay Win protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Uplay Win protocol
	ProtocolID = 0x31

	// MethodGetActions is the method ID for the method GetActions
	MethodGetActions = 0x1

	// MethodGetActionsCompleted is the method ID for the method GetActionsCompleted
	MethodGetActionsCompleted = 0x2

	// MethodGetActionsCount is the method ID for the method GetActionsCount
	MethodGetActionsCount = 0x3

	// MethodGetActionsCompletedCount is the method ID for the method GetActionsCompletedCount
	MethodGetActionsCompletedCount = 0x4

	// MethodGetRewards is the method ID for the method GetRewards
	MethodGetRewards = 0x5

	// MethodGetRewardsPurchased is the method ID for the method GetRewardsPurchased
	MethodGetRewardsPurchased = 0x6

	// MethodUplayWelcome is the method ID for the method UplayWelcome
	MethodUplayWelcome = 0x7

	// MethodSetActionCompleted is the method ID for the method SetActionCompleted
	MethodSetActionCompleted = 0x8

	// MethodSetActionsCompleted is the method ID for the method SetActionsCompleted
	MethodSetActionsCompleted = 0x9

	// MethodGetUserToken is the method ID for the method GetUserToken
	MethodGetUserToken = 0xA

	// MethodGetVirtualCurrencyUserBalance is the method ID for the method GetVirtualCurrencyUserBalance
	MethodGetVirtualCurrencyUserBalance = 0xB

	// MethodGetSectionsByKey is the method ID for the method GetSectionsByKey
	MethodGetSectionsByKey = 0xC
)

// Protocol stores all the RMC method handlers for the Uplay win protocol and listens for requests
type Protocol struct {
	endpoint                      nex.EndpointInterface
	GetActions                    func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)
	GetActionsCompleted           func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)
	GetActionsCount               func(err error, packet nex.PacketInterface, callID uint32, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)
	GetActionsCompletedCount      func(err error, packet nex.PacketInterface, callID uint32, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)
	GetRewards                    func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)
	GetRewardsPurchased           func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)
	UplayWelcome                  func(err error, packet nex.PacketInterface, callID uint32, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)
	SetActionCompleted            func(err error, packet nex.PacketInterface, callID uint32, actionCode types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)
	SetActionsCompleted           func(err error, packet nex.PacketInterface, callID uint32, actionCodeList types.List[types.String], cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)
	GetUserToken                  func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetVirtualCurrencyUserBalance func(err error, packet nex.PacketInterface, callID uint32, platformCode types.String) (*nex.RMCMessage, *nex.Error)
	GetSectionsByKey              func(err error, packet nex.PacketInterface, callID uint32, cultureName types.String, sectionKey types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)
	Patches                       nex.ServiceProtocol
	PatchedMethods                []uint32
}

// Interface implements the methods present on the Uplay win protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerGetActions(handler func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetActionsCompleted(handler func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetActionsCount(handler func(err error, packet nex.PacketInterface, callID uint32, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetActionsCompletedCount(handler func(err error, packet nex.PacketInterface, callID uint32, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRewards(handler func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRewardsPurchased(handler func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerUplayWelcome(handler func(err error, packet nex.PacketInterface, callID uint32, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerSetActionCompleted(handler func(err error, packet nex.PacketInterface, callID uint32, actionCode types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerSetActionsCompleted(handler func(err error, packet nex.PacketInterface, callID uint32, actionCodeList types.List[types.String], cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetUserToken(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetVirtualCurrencyUserBalance(handler func(err error, packet nex.PacketInterface, callID uint32, platformCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSectionsByKey(handler func(err error, packet nex.PacketInterface, callID uint32, cultureName types.String, sectionKey types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerGetActions sets the handler for the GetActions method
func (protocol *Protocol) SetHandlerGetActions(handler func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetActions = handler
}

// SetHandlerGetActionsCompleted sets the handler for the GetActionsCompleted method
func (protocol *Protocol) SetHandlerGetActionsCompleted(handler func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetActionsCompleted = handler
}

// SetHandlerGetActionsCount sets the handler for the GetActionsCount method
func (protocol *Protocol) SetHandlerGetActionsCount(handler func(err error, packet nex.PacketInterface, callID uint32, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetActionsCount = handler
}

// SetHandlerGetActionsCompletedCount sets the handler for the GetActionsCompletedCount method
func (protocol *Protocol) SetHandlerGetActionsCompletedCount(handler func(err error, packet nex.PacketInterface, callID uint32, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetActionsCompletedCount = handler
}

// SetHandlerGetRewards sets the handler for the GetRewards method
func (protocol *Protocol) SetHandlerGetRewards(handler func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRewards = handler
}

// SetHandlerGetRewardsPurchased sets the handler for the GetRewardsPurchased method
func (protocol *Protocol) SetHandlerGetRewardsPurchased(handler func(err error, packet nex.PacketInterface, callID uint32, startRowIndex types.Int32, maximumRows types.Int32, sortExpression types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRewardsPurchased = handler
}

// SetHandlerUplayWelcome sets the handler for the UplayWelcome method
func (protocol *Protocol) SetHandlerUplayWelcome(handler func(err error, packet nex.PacketInterface, callID uint32, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.UplayWelcome = handler
}

// SetHandlerSetActionCompleted sets the handler for the SetActionCompleted method
func (protocol *Protocol) SetHandlerSetActionCompleted(handler func(err error, packet nex.PacketInterface, callID uint32, actionCode types.String, cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.SetActionCompleted = handler
}

// SetHandlerSetActionsCompleted sets the handler for the SetActionsCompleted method
func (protocol *Protocol) SetHandlerSetActionsCompleted(handler func(err error, packet nex.PacketInterface, callID uint32, actionCodeList types.List[types.String], cultureName types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.SetActionsCompleted = handler
}

// SetHandlerGetUserToken sets the handler for the GetUserToken method
func (protocol *Protocol) SetHandlerGetUserToken(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetUserToken = handler
}

// SetHandlerGetVirtualCurrencyUserBalance sets the handler for the GetVirtualCurrencyUserBalance method
func (protocol *Protocol) SetHandlerGetVirtualCurrencyUserBalance(handler func(err error, packet nex.PacketInterface, callID uint32, platformCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetVirtualCurrencyUserBalance = handler
}

// SetHandlerGetSectionsByKey sets the handler for the GetSectionsByKey method
func (protocol *Protocol) SetHandlerGetSectionsByKey(handler func(err error, packet nex.PacketInterface, callID uint32, cultureName types.String, sectionKey types.String, platformCode types.String, gameCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSectionsByKey = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodGetActions:
		protocol.handleGetActions(packet)
	case MethodGetActionsCompleted:
		protocol.handleGetActionsCompleted(packet)
	case MethodGetActionsCount:
		protocol.handleGetActionsCount(packet)
	case MethodGetActionsCompletedCount:
		protocol.handleGetActionsCompletedCount(packet)
	case MethodGetRewards:
		protocol.handleGetRewards(packet)
	case MethodGetRewardsPurchased:
		protocol.handleGetRewardsPurchased(packet)
	case MethodUplayWelcome:
		protocol.handleUplayWelcome(packet)
	case MethodSetActionCompleted:
		protocol.handleSetActionCompleted(packet)
	case MethodSetActionsCompleted:
		protocol.handleSetActionsCompleted(packet)
	case MethodGetUserToken:
		protocol.handleGetUserToken(packet)
	case MethodGetVirtualCurrencyUserBalance:
		protocol.handleGetVirtualCurrencyUserBalance(packet)
	case MethodGetSectionsByKey:
		protocol.handleGetSectionsByKey(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Uplay Win method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Uplay Win protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
