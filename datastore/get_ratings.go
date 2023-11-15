// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetRatings(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetRatings == nil {
		globals.Logger.Warning("DataStore::GetRatings not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		_, errorCode = protocol.GetRatings(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	accessPassword, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.GetRatings(fmt.Errorf("Failed to read accessPassword from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetRatings(nil, packet, callID, dataIDs, accessPassword)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
