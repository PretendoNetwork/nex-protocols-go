// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveBlackList sets the RemoveBlackList handler function
func (protocol *Protocol) RemoveBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, pid *nex.PID) uint32) {
	protocol.removeBlackListHandler = handler
}

func (protocol *Protocol) handleRemoveBlackList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.removeBlackListHandler == nil {
		globals.Logger.Warning("FriendsWiiU::RemoveBlackList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadPID()
	if err != nil {
		errorCode = protocol.removeBlackListHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.removeBlackListHandler(nil, packet, callID, pid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
