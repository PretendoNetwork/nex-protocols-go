// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetDetailedParticipants sets the GetDetailedParticipants handler function
func (protocol *Protocol) GetDetailedParticipants(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool)) {
	protocol.getDetailedParticipantsHandler = handler
}

func (protocol *Protocol) handleGetDetailedParticipants(packet nex.PacketInterface) {
	if protocol.getDetailedParticipantsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::GetDetailedParticipants not implemented")
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
		go protocol.getDetailedParticipantsHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, false)
	}

	bOnlyActive, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.getDetailedParticipantsHandler(fmt.Errorf("Failed to read bOnlyActive from parameters. %s", err.Error()), client, callID, 0, false)
	}

	go protocol.getDetailedParticipantsHandler(nil, client, callID, idGathering, bOnlyActive)
}
