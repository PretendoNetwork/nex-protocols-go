// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetState sets the GetState handler function
func (protocol *Protocol) GetState(handler func(err error, client *nex.Client, callID uint32, idGathering uint32) uint32) {
	protocol.getStateHandler = handler
}

func (protocol *Protocol) handleGetState(packet nex.PacketInterface) {
	if protocol.getStateHandler == nil {
		globals.Logger.Warning("MatchMaking::GetState not implemented")
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
		go protocol.getStateHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0)
	}

	go protocol.getStateHandler(nil, client, callID, idGathering)
}
