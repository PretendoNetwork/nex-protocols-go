// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	utility_types "github.com/PretendoNetwork/nex-protocols-go/utility/types"
)

const (
	// ProtocolID is the protocol ID for the Utility protocol
	ProtocolID = 0x6E

	// MethodAcquireNexUniqueID is the method ID for the method AcquireNexUniqueID
	MethodAcquireNexUniqueID = 0x1

	// MethodAcquireNexUniqueIDWithPassword is the method ID for the method AcquireNexUniqueIDWithPassword
	MethodAcquireNexUniqueIDWithPassword = 0x2

	// MethodAssociateNexUniqueIDWithMyPrincipalID is the method ID for the method AssociateNexUniqueIDWithMyPrincipalID
	MethodAssociateNexUniqueIDWithMyPrincipalID = 0x3

	// MethodAssociateNexUniqueIDsWithMyPrincipalID is the method ID for the method AssociateNexUniqueIDsWithMyPrincipalID
	MethodAssociateNexUniqueIDsWithMyPrincipalID = 0x4

	// MethodGetAssociatedNexUniqueIDWithMyPrincipalID is the method ID for the method GetAssociatedNexUniqueIDWithMyPrincipalID
	MethodGetAssociatedNexUniqueIDWithMyPrincipalID = 0x5

	// MethodGetAssociatedNexUniqueIDsWithMyPrincipalID is the method ID for the method GetAssociatedNexUniqueIDsWithMyPrincipalID
	MethodGetAssociatedNexUniqueIDsWithMyPrincipalID = 0x6

	// MethodGetIntegerSettings is the method ID for the method GetIntegerSettings
	MethodGetIntegerSettings = 0x7

	// MethodGetStringSettings is the method ID for the method GetStringSettings
	MethodGetStringSettings = 0x8
)

// Protocol stores all the RMC method handlers for the Utility protocol and listens for requests
type Protocol struct {
	Server                                     nex.ServerInterface
	AcquireNexUniqueID                         func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	AcquireNexUniqueIDWithPassword             func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	AssociateNexUniqueIDWithMyPrincipalID      func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo *utility_types.UniqueIDInfo) (*nex.RMCMessage, uint32)
	AssociateNexUniqueIDsWithMyPrincipalID     func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo []*utility_types.UniqueIDInfo) (*nex.RMCMessage, uint32)
	GetAssociatedNexUniqueIDWithMyPrincipalID  func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetAssociatedNexUniqueIDsWithMyPrincipalID func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetIntegerSettings                         func(err error, packet nex.PacketInterface, callID uint32, integerSettingIndex uint32) (*nex.RMCMessage, uint32)
	GetStringSettings                          func(err error, packet nex.PacketInterface, callID uint32, stringSettingIndex uint32) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			switch message.MethodID {
			case MethodAcquireNexUniqueID:
				protocol.handleAcquireNexUniqueID(packet)
			case MethodAcquireNexUniqueIDWithPassword:
				protocol.handleAcquireNexUniqueIDWithPassword(packet)
			case MethodAssociateNexUniqueIDWithMyPrincipalID:
				protocol.handleAssociateNexUniqueIDWithMyPrincipalID(packet)
			case MethodAssociateNexUniqueIDsWithMyPrincipalID:
				protocol.handleAssociateNexUniqueIDsWithMyPrincipalID(packet)
			case MethodGetAssociatedNexUniqueIDWithMyPrincipalID:
				protocol.handleGetAssociatedNexUniqueIDWithMyPrincipalID(packet)
			case MethodGetAssociatedNexUniqueIDsWithMyPrincipalID:
				protocol.handleGetAssociatedNexUniqueIDsWithMyPrincipalID(packet)
			case MethodGetIntegerSettings:
				protocol.handleGetIntegerSettings(packet)
			case MethodGetStringSettings:
				protocol.handleGetStringSettings(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Utility method ID: %#v\n", message.MethodID)
			}
		}
	})
}

// NewProtocol returns a new Utility protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
