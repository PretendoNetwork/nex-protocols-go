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
	Server                                            *nex.Server
	acquireNexUniqueIDHandler                         func(err error, packet nex.PacketInterface, callID uint32) uint32
	acquireNexUniqueIDWithPasswordHandler             func(err error, packet nex.PacketInterface, callID uint32) uint32
	associateNexUniqueIDWithMyPrincipalIDHandler      func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo *utility_types.UniqueIDInfo) uint32
	associateNexUniqueIDsWithMyPrincipalIDHandler     func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo []*utility_types.UniqueIDInfo) uint32
	getAssociatedNexUniqueIDWithMyPrincipalIDHandler  func(err error, packet nex.PacketInterface, callID uint32) uint32
	getAssociatedNexUniqueIDsWithMyPrincipalIDHandler func(err error, packet nex.PacketInterface, callID uint32) uint32
	getIntegerSettingsHandler                         func(err error, packet nex.PacketInterface, callID uint32, integerSettingIndex uint32) uint32
	getStringSettingsHandler                          func(err error, packet nex.PacketInterface, callID uint32, stringSettingIndex uint32) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodAcquireNexUniqueID:
				go protocol.handleAcquireNexUniqueID(packet)
			case MethodAcquireNexUniqueIDWithPassword:
				go protocol.handleAcquireNexUniqueIDWithPassword(packet)
			case MethodAssociateNexUniqueIDWithMyPrincipalID:
				go protocol.handleAssociateNexUniqueIDWithMyPrincipalID(packet)
			case MethodAssociateNexUniqueIDsWithMyPrincipalID:
				go protocol.handleAssociateNexUniqueIDsWithMyPrincipalID(packet)
			case MethodGetAssociatedNexUniqueIDWithMyPrincipalID:
				go protocol.handleGetAssociatedNexUniqueIDWithMyPrincipalID(packet)
			case MethodGetAssociatedNexUniqueIDsWithMyPrincipalID:
				go protocol.handleGetAssociatedNexUniqueIDsWithMyPrincipalID(packet)
			case MethodGetIntegerSettings:
				go protocol.handleGetIntegerSettings(packet)
			case MethodGetStringSettings:
				go protocol.handleGetStringSettings(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Utility method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewProtocol returns a new Utility protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
