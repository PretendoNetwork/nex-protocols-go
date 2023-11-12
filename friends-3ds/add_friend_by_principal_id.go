// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendByPrincipalID sets the AddFriendByPrincipalID handler function
func (protocol *Protocol) AddFriendByPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, lfc uint64, pid uint32) uint32) {
	protocol.addFriendByPrincipalIDHandler = handler
}

func (protocol *Protocol) handleAddFriendByPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.addFriendByPrincipalIDHandler == nil {
		globals.Logger.Warning("Friends3DS::AddFriendByPrincipalID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.addFriendByPrincipalIDHandler(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.addFriendByPrincipalIDHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.addFriendByPrincipalIDHandler(nil, packet, callID, lfc, pid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
