// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveFriend sets the RemoveFriend handler function
func (protocol *Protocol) RemoveFriend(handler func(err error, packet nex.PacketInterface, callID uint32, pid uint32) uint32) {
	protocol.removeFriendHandler = handler
}

func (protocol *Protocol) handleRemoveFriend(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.removeFriendHandler == nil {
		globals.Logger.Warning("FriendsWiiU::RemoveFriend not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.removeFriendHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.removeFriendHandler(nil, packet, callID, pid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
