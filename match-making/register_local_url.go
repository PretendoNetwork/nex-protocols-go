// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRegisterLocalURL(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RegisterLocalURL == nil {
		globals.Logger.Warning("MatchMaking::RegisterLocalURL not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.RegisterLocalURL(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	url, err := parametersStream.ReadStationURL()
	if err != nil {
		errorCode = protocol.RegisterLocalURL(fmt.Errorf("Failed to read url from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.RegisterLocalURL(nil, packet, callID, gid, url)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
