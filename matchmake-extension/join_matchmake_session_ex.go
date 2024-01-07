// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleJoinMatchmakeSessionEx(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.JoinMatchmakeSessionEx == nil {
		globals.Logger.Warning("MatchmakeExtension::JoinMatchmakeSessionEx not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gid := types.NewPrimitiveU32(0)
	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, "", false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage := types.NewString("")
	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, 0, "", false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dontCareMyBlockList := types.NewPrimitiveBool(false)
	err = dontCareMyBlockList.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read dontCareMyBlockList from parameters. %s", err.Error()), packet, callID, 0, "", false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	participationCount := types.NewPrimitiveU16(0)
	err = participationCount.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read participationCount from parameters. %s", err.Error()), packet, callID, 0, "", false, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.JoinMatchmakeSessionEx(nil, packet, callID, gid, strMessage, dontCareMyBlockList, participationCount)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
