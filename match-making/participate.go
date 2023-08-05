// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Participate sets the Participate handler function
func (protocol *Protocol) Participate(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string) uint32) {
	protocol.participateHandler = handler
}

func (protocol *Protocol) handleParticipate(packet nex.PacketInterface) {
	if protocol.participateHandler == nil {
		globals.Logger.Warning("MatchMaking::Participate not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.participateHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, "")
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.participateHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, "")
	}

	go protocol.participateHandler(nil, client, callID, idGathering, strMessage)
}
