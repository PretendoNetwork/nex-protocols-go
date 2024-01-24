// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRegisterLocalURLs(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.RegisterLocalURLs == nil {
		globals.Logger.Warning("MatchMaking::RegisterLocalURLs not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gid := types.NewPrimitiveU32(0)
	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RegisterLocalURLs(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstURLs := types.NewList[*types.StationURL]()
	lstURLs.Type = types.NewStationURL("")
	err = lstURLs.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RegisterLocalURLs(fmt.Errorf("Failed to read lstURLs from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RegisterLocalURLs(nil, packet, callID, gid, lstURLs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
