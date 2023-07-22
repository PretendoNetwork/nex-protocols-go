// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddParticipants sets the AddParticipants handler function
func (protocol *MatchMakingProtocol) AddParticipants(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, lstPrincipals []uint32, strMessage string)) {
	protocol.addParticipantsHandler = handler
}

func (protocol *MatchMakingProtocol) handleAddParticipants(packet nex.PacketInterface) {
	if protocol.addParticipantsHandler == nil {
		globals.Logger.Warning("MatchMaking::AddParticipants not implemented")
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
		go protocol.addParticipantsHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, nil, "")
	}

	lstPrincipals, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.addParticipantsHandler(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), client, callID, 0, nil, "")
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.addParticipantsHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, nil, "")
	}

	go protocol.addParticipantsHandler(nil, client, callID, idGathering, lstPrincipals, strMessage)
}
