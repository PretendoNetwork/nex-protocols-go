// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetParticipants sets the GetParticipants handler function
func (protocol *MatchMakingProtocol) GetParticipants(handler func(err error, client *nex.Client, callID uint32, idGathering uint32)) {
	protocol.getParticipantsHandler = handler
}

func (protocol *MatchMakingProtocol) handleGetParticipants(packet nex.PacketInterface) {
	if protocol.getParticipantsHandler == nil {
		globals.Logger.Warning("MatchMaking::GetParticipants not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getParticipantsHandler(fmt.Errorf("Failed to read gatheringID from parameters. %s", err.Error()), client, callID, 0)
	}

	go protocol.getParticipantsHandler(nil, client, callID, idGathering)
}
