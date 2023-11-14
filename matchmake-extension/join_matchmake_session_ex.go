// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleJoinMatchmakeSessionEx(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.JoinMatchmakeSessionEx == nil {
		globals.Logger.Warning("MatchmakeExtension::JoinMatchmakeSessionEx not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, "", false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, 0, "", false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dontCareMyBlockList, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read dontCareMyBlockList from parameters. %s", err.Error()), packet, callID, 0, "", false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	participationCount, err := parametersStream.ReadUInt16LE()
	if err != nil {
		errorCode = protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read participationCount from parameters. %s", err.Error()), packet, callID, 0, "", false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.JoinMatchmakeSessionEx(nil, packet, callID, gid, strMessage, dontCareMyBlockList, participationCount)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
