// Package protocol implements the Subscription protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
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

	// MethodGetSubscriptionUserFriendList is the method ID for the method GetSubscriptionUserFriendList
	MethodGetSubscriptionUserFriendList = 0xE

	// MethodGetPrivacyLevels is the method ID for the method GetPrivacyLevels
	MethodGetPrivacyLevels = 0xF
)

// Protocol handles the Subscription nex protocol
type Protocol struct {
	server                              nex.ServerInterface
	CreateMySubscriptionData            func(err error, packet nex.PacketInterface, callID uint32, unk *types.PrimitiveU64, content []byte) (*nex.RMCMessage, uint32)
	UpdateMySubscriptionData            func(err error, packet nex.PacketInterface, callID uint32, unk *types.PrimitiveU32, content []byte) (*nex.RMCMessage, uint32)
	GetFriendSubscriptionData           func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetTargetSubscriptionData           func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetActivePlayerSubscriptionData     func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetSubscriptionData                 func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)
	ReplaceTargetAndGetSubscriptionData func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetPrivacyLevels                    func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Subscription protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerCreateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, unk *types.PrimitiveU64, content []byte) (*nex.RMCMessage, uint32))
	SetHandlerUpdateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, unk *types.PrimitiveU32, content []byte) (*nex.RMCMessage, uint32))
	SetHandlerGetFriendSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerGetTargetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerGetActivePlayerSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerGetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32))
	SetHandlerReplaceTargetAndGetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerGetPrivacyLevels(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerCreateMySubscriptionData sets the handler for the CreateMySubscriptionData method
func (protocol *Protocol) SetHandlerCreateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, unk *types.PrimitiveU64, content []byte) (*nex.RMCMessage, uint32)) {
	protocol.CreateMySubscriptionData = handler
}

// SetHandlerUpdateMySubscriptionData sets the handler for the UpdateMySubscriptionData method
func (protocol *Protocol) SetHandlerUpdateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, unk *types.PrimitiveU32, content []byte) (*nex.RMCMessage, uint32)) {
	protocol.UpdateMySubscriptionData = handler
}

// SetHandlerGetFriendSubscriptionData sets the handler for the GetFriendSubscriptionData method
func (protocol *Protocol) SetHandlerGetFriendSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetFriendSubscriptionData = handler
}

// SetHandlerGetTargetSubscriptionData sets the handler for the GetTargetSubscriptionData method
func (protocol *Protocol) SetHandlerGetTargetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetTargetSubscriptionData = handler
}

// SetHandlerGetActivePlayerSubscriptionData sets the handler for the GetActivePlayerSubscriptionData method
func (protocol *Protocol) SetHandlerGetActivePlayerSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetActivePlayerSubscriptionData = handler
}

// SetHandlerGetSubscriptionData sets the handler for the GetSubscriptionData method
func (protocol *Protocol) SetHandlerGetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)) {
	protocol.GetSubscriptionData = handler
}

// SetHandlerReplaceTargetAndGetSubscriptionData sets the handler for the ReplaceTargetAndGetSubscriptionData method
func (protocol *Protocol) SetHandlerReplaceTargetAndGetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.ReplaceTargetAndGetSubscriptionData = handler
}

// SetHandlerGetPrivacyLevels sets the handler for the GetPrivacyLevels method
func (protocol *Protocol) SetHandlerGetPrivacyLevels(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetPrivacyLevels = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if ProtocolID == request.ProtocolID {
			switch request.MethodID {
			case MethodCreateMySubscriptionData:
				protocol.handleCreateMySubscriptionData(packet)
			case MethodUpdateMySubscriptionData:
				protocol.handleUpdateMySubscriptionData(packet)
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
			case MethodGetPrivacyLevels:
				protocol.handleGetPrivacyLevels(packet)
			default:
				globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Subscription method ID: %#v\n", request.MethodID)
			}
		}
	})
}

// NewProtocol returns a new Protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
