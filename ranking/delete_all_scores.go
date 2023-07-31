// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteAllScores sets the DeleteAllScores handler function
func (protocol *Protocol) DeleteAllScores(handler func(err error, client *nex.Client, callID uint32, uniqueID uint64)) {
	protocol.deleteAllScoresHandler = handler
}

func (protocol *Protocol) handleDeleteAllScores(packet nex.PacketInterface) {
	if protocol.deleteAllScoresHandler == nil {
		globals.Logger.Warning("Ranking::DeleteAllScores not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.deleteAllScoresHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.deleteAllScoresHandler(nil, client, callID, uniqueID)
}
