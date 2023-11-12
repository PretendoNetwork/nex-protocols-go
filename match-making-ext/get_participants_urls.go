// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetParticipantsURLs sets the GetParticipantsURLs handler function
func (protocol *Protocol) GetParticipantsURLs(handler func(err error, packet nex.PacketInterface, callID uint32, lstGatherings []uint32) uint32) {
	protocol.getParticipantsURLsHandler = handler
}

func (protocol *Protocol) handleGetParticipantsURLs(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getParticipantsURLsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::GetParticipantsURLs not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGatherings, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getParticipantsURLsHandler(fmt.Errorf("Failed to read lstGatherings from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getParticipantsURLsHandler(nil, packet, callID, lstGatherings)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
