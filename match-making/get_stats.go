// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetStats sets the GetStats handler function
func (protocol *Protocol) GetStats(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, lstParticipants []uint32, lstColumns []byte)) {
	protocol.getStatsHandler = handler
}

func (protocol *Protocol) handleGetStats(packet nex.PacketInterface) {
	if protocol.getStatsHandler == nil {
		globals.Logger.Warning("MatchMaking::GetStats not implemented")
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
		go protocol.getStatsHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, nil, nil)
	}

	lstParticipants, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getStatsHandler(fmt.Errorf("Failed to read lstParticipants from parameters. %s", err.Error()), client, callID, 0, nil, nil)
	}

	lstColumns, err := parametersStream.ReadBuffer() // * This is documented as List<byte>, but that's justs a buffer so...
	if err != nil {
		go protocol.getStatsHandler(fmt.Errorf("Failed to read lstColumns from parameters. %s", err.Error()), client, callID, 0, nil, nil)
	}

	go protocol.getStatsHandler(nil, client, callID, idGathering, lstParticipants, lstColumns)
}
