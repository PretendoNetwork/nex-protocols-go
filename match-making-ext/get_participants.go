// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetParticipants sets the GetParticipants handler function
func (protocol *Protocol) GetParticipants(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool)) {
	protocol.GetParticipantsHandler = handler
}

func (protocol *Protocol) handleGetParticipants(packet nex.PacketInterface) {
	if protocol.GetParticipantsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::GetParticipants not implemented")
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
		go protocol.GetParticipantsHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, false)
	}

	bOnlyActive, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.GetParticipantsHandler(fmt.Errorf("Failed to read bOnlyActive from parameters. %s", err.Error()), client, callID, 0, false)
	}

	go protocol.GetParticipantsHandler(nil, client, callID, idGathering, bOnlyActive)
}
