// Package protocol implements the Subscription protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"

	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	subscription_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscription/types"
)

const (
	// ProtocolID is the protocol ID for the Subscription protocol
	ProtocolID = 0x75

	// MethodCreateMySubscriptionData is the method ID for the method CreateMySubscriptionDataID
	MethodCreateMySubscriptionData = 0x1

	// MethodUpdateMySubscriptionData is the method ID for the method UpdateMySubscriptionData
	MethodUpdateMySubscriptionData = 0x2

	// MethodClearMySubscriptionData is the method ID for the method ClearMySubscriptionDataID
	MethodClearMySubscriptionData = 0x3

	// MethodAddTarget is the method ID for the method AddTarget
	MethodAddTarget = 0x4

	// MethodDeleteTarget is the method ID for the method DeleteTarget
	MethodDeleteTarget = 0x5

	// MethodClearTarget is the method ID for the method ClearTarget
	MethodClearTarget = 0x6

	// MethodGetFriendSubscriptionData is the method ID for the method GetFriendSubscriptionData
	MethodGetFriendSubscriptionData = 0x7

	// MethodGetTargetSubscriptionData is the method ID for the method GetTargetSubscriptionData
	MethodGetTargetSubscriptionData = 0x8

	// MethodGetActivePlayerSubscriptionData is the method ID for the method GetActivePlayerSubscriptionData
	MethodGetActivePlayerSubscriptionData = 0x9

	// MethodGetSubscriptionData is the method ID for the method GetSubscriptionData
	MethodGetSubscriptionData = 0xA

	// MethodReplaceTargetAndGetSubscriptionData is the method ID for the method ReplaceTargetAndGetSubscriptionData
	MethodReplaceTargetAndGetSubscriptionData = 0xB

	// MethodSetPrivacyLevel is the method ID for the method SetPrivacyLevel
	MethodSetPrivacyLevel = 0xC

	// MethodGetPrivacyLevel is the method ID for the method GetPrivacyLevel
	MethodGetPrivacyLevel = 0xD
)

// Protocol handles the Subscription nex protocol
type Protocol struct {
	endpoint                            nex.EndpointInterface
	CreateMySubscriptionData            func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, param subscription_types.SubscriptionData, unknown2 types.Bool) (*nex.RMCMessage, *nex.Error)
	UpdateMySubscriptionData            func(err error, packet nex.PacketInterface, callID uint32, param subscription_types.SubscriptionData) (*nex.RMCMessage, *nex.Error)
	ClearMySubscriptionData             func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	AddTarget                           func(err error, packet nex.PacketInterface, callID uint32, targets types.List[types.PID]) (*nex.RMCMessage, *nex.Error)
	DeleteTarget                        func(err error, packet nex.PacketInterface, callID uint32, targets types.List[types.PID]) (*nex.RMCMessage, *nex.Error)
	ClearTarget                         func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetFriendSubscriptionData           func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetTargetSubscriptionData           func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetActivePlayerSubscriptionData     func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, unknown2 types.UInt32, unknown3 types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetSubscriptionData                 func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.PID]) (*nex.RMCMessage, *nex.Error)
	ReplaceTargetAndGetSubscriptionData func(err error, packet nex.PacketInterface, callID uint32, newTargets types.List[types.PID]) (*nex.RMCMessage, *nex.Error)
	SetPrivacyLevel                     func(err error, packet nex.PacketInterface, callID uint32, privacyLevel types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetPrivacyLevel                     func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Patches                             nex.ServiceProtocol
	PatchedMethods                      []uint32
}

// Interface implements the methods present on the Subscription protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerCreateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, param subscription_types.SubscriptionData, unknown2 types.Bool) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, param subscription_types.SubscriptionData) (*nex.RMCMessage, *nex.Error))
	SetHandlerClearMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddTarget(handler func(err error, packet nex.PacketInterface, callID uint32, targets types.List[types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteTarget(handler func(err error, packet nex.PacketInterface, callID uint32, targets types.List[types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerClearTarget(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetFriendSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetTargetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetActivePlayerSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, unknown2 types.UInt32, unknown3 types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReplaceTargetAndGetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, newTargets types.List[types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerSetPrivacyLevel(handler func(err error, packet nex.PacketInterface, callID uint32, privacyLevel types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPrivacyLevel(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerCreateMySubscriptionData sets the handler for the CreateMySubscriptionData method
func (protocol *Protocol) SetHandlerCreateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, param subscription_types.SubscriptionData, unknown2 types.Bool) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateMySubscriptionData = handler
}

// SetHandlerUpdateMySubscriptionData sets the handler for the UpdateMySubscriptionData method
func (protocol *Protocol) SetHandlerUpdateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, param subscription_types.SubscriptionData) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateMySubscriptionData = handler
}

// SetHandlerClearMySubscriptionData sets the handler for the ClearMySubscriptionData method
func (protocol *Protocol) SetHandlerClearMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ClearMySubscriptionData = handler
}

// SetHandlerAddTarget sets the handler for the AddTarget method
func (protocol *Protocol) SetHandlerAddTarget(handler func(err error, packet nex.PacketInterface, callID uint32, targets types.List[types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddTarget = handler
}

// SetHandlerDeleteTarget sets the handler for the DeleteTarget method
func (protocol *Protocol) SetHandlerDeleteTarget(handler func(err error, packet nex.PacketInterface, callID uint32, targets types.List[types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteTarget = handler
}

// SetHandlerClearTarget sets the handler for the ClearTarget method
func (protocol *Protocol) SetHandlerClearTarget(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ClearMySubscriptionData = handler
}

// SetHandlerGetFriendSubscriptionData sets the handler for the GetFriendSubscriptionData method
func (protocol *Protocol) SetHandlerGetFriendSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetFriendSubscriptionData = handler
}

// SetHandlerGetTargetSubscriptionData sets the handler for the GetTargetSubscriptionData method
func (protocol *Protocol) SetHandlerGetTargetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetTargetSubscriptionData = handler
}

// SetHandlerGetActivePlayerSubscriptionData sets the handler for the GetActivePlayerSubscriptionData method
func (protocol *Protocol) SetHandlerGetActivePlayerSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, unknown2 types.UInt32, unknown3 types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetActivePlayerSubscriptionData = handler
}

// SetHandlerGetSubscriptionData sets the handler for the GetSubscriptionData method
func (protocol *Protocol) SetHandlerGetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSubscriptionData = handler
}

// SetHandlerReplaceTargetAndGetSubscriptionData sets the handler for the ReplaceTargetAndGetSubscriptionData method
func (protocol *Protocol) SetHandlerReplaceTargetAndGetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, newTargets types.List[types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReplaceTargetAndGetSubscriptionData = handler
}

// SetHandlerSetPrivacyLevel sets the handler for the SetPrivacyLevel method
func (protocol *Protocol) SetHandlerSetPrivacyLevel(handler func(err error, packet nex.PacketInterface, callID uint32, privacyLevel types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.SetPrivacyLevel = handler
}

// SetHandlerGetPrivacyLevel sets the handler for the GetPrivacyLevel method
func (protocol *Protocol) SetHandlerGetPrivacyLevel(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPrivacyLevel = handler
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
	case MethodCreateMySubscriptionData:
		protocol.handleCreateMySubscriptionData(packet)
	case MethodUpdateMySubscriptionData:
		protocol.handleUpdateMySubscriptionData(packet)
	case MethodClearMySubscriptionData:
		protocol.handleClearMySubscriptionData(packet)
	case MethodAddTarget:
		protocol.handleAddTarget(packet)
	case MethodDeleteTarget:
		protocol.handleDeleteTarget(packet)
	case MethodClearTarget:
		protocol.handleClearTarget(packet)
	case MethodGetFriendSubscriptionData:
		protocol.handleGetFriendSubscriptionData(packet)
	case MethodGetTargetSubscriptionData:
		protocol.handleGetTargetSubscriptionData(packet)
	case MethodGetActivePlayerSubscriptionData:
		protocol.handleGetActivePlayerSubscriptionData(packet)
	case MethodGetSubscriptionData:
		protocol.handleGetSubscriptionData(packet)
	case MethodReplaceTargetAndGetSubscriptionData:
		protocol.handleReplaceTargetAndGetSubscriptionData(packet)
	case MethodSetPrivacyLevel:
		protocol.handleSetPrivacyLevel(packet)
	case MethodGetPrivacyLevel:
		protocol.handleGetPrivacyLevel(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Subscription method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
