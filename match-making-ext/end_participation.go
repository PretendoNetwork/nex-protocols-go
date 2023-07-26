// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EndParticipation sets the EndParticipation handler function
func (protocol *Protocol) EndParticipation(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string)) {
	protocol.EndParticipationHandler = handler
}

func (protocol *Protocol) handleEndParticipation(packet nex.PacketInterface) {
	if protocol.EndParticipationHandler == nil {
		globals.Logger.Warning("MatchMakingExt::EndParticipation not implemented")
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
		go protocol.EndParticipationHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.EndParticipationHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, "")
	}

	go protocol.EndParticipationHandler(nil, client, callID, idGathering, strMessage)
}
