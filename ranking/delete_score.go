// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteScore sets the DeleteScore handler function
func (protocol *Protocol) DeleteScore(handler func(err error, client *nex.Client, callID uint32, category uint32, uniqueID uint64)) {
	protocol.deleteScoreHandler = handler
}

func (protocol *Protocol) handleDeleteScore(packet nex.PacketInterface) {
	if protocol.deleteScoreHandler == nil {
		globals.Logger.Warning("Ranking::DeleteScore not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.deleteScoreHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.deleteScoreHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.deleteScoreHandler(nil, client, callID, category, uniqueID)
}
