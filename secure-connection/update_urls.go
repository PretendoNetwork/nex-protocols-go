// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateURLs(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateURLs == nil {
		globals.Logger.Warning("SecureConnection::UpdateURLs not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	vecMyURLs := types.NewList[*types.StationURL]()
	vecMyURLs.Type = types.NewStationURL("")
	err = vecMyURLs.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateURLs(fmt.Errorf("Failed to read vecMyURLs from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateURLs(nil, packet, callID, vecMyURLs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
