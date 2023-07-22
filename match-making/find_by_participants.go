// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByParticipants sets the FindByParticipants handler function
func (protocol *MatchMakingProtocol) FindByParticipants(handler func(err error, client *nex.Client, callID uint32, pid []uint32)) {
	protocol.findByParticipantsHandler = handler
}

func (protocol *MatchMakingProtocol) handleFindByParticipants(packet nex.PacketInterface) {
	if protocol.findByParticipantsHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByParticipants not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.findByParticipantsHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, nil)
	}

	go protocol.findByParticipantsHandler(nil, client, callID, pid)
}