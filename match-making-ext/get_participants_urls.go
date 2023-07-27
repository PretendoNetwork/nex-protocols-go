// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetParticipantsURLs sets the GetParticipantsURLs handler function
func (protocol *Protocol) GetParticipantsURLs(handler func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)) {
	protocol.getParticipantsURLsHandler = handler
}

func (protocol *Protocol) handleGetParticipantsURLs(packet nex.PacketInterface) {
	if protocol.getParticipantsURLsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::GetParticipantsURLs not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGatherings, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getParticipantsURLsHandler(fmt.Errorf("Failed to read lstGatherings from parameters. %s", err.Error()), client, callID, nil)
	}

	go protocol.getParticipantsURLsHandler(nil, client, callID, lstGatherings)
}
