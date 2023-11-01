// Package protocol implements the Matchmake tension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinMatchmakeSession sets the JoinMatchmakeSession handler function
func (protocol *Protocol) JoinMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, gid uint32, strMessage string) uint32) {
	protocol.joinMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleJoinMatchmakeSession(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.joinMatchmakeSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::JoinMatchmakeSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.joinMatchmakeSessionHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.joinMatchmakeSessionHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.joinMatchmakeSessionHandler(nil, packet, callID, gid, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
