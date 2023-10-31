// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveFriendByPrincipalID sets the RemoveFriendByPrincipalID handler function
func (protocol *Protocol) RemoveFriendByPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, pid uint32) uint32) {
	protocol.removeFriendByPrincipalIDHandler = handler
}

func (protocol *Protocol) handleRemoveFriendByPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.removeFriendByPrincipalIDHandler == nil {
		globals.Logger.Warning("Friends3DS::RemoveFriendByPrincipalID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.removeFriendByPrincipalIDHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.removeFriendByPrincipalIDHandler(nil, packet, callID, pid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
