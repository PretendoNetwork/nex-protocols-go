// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRound sets the GetRound handler function
func (protocol *Protocol) GetRound(handler func(err error, client *nex.Client, callID uint32, roundID uint64) uint32) {
	protocol.getRoundHandler = handler
}

func (protocol *Protocol) handleGetRound(packet nex.PacketInterface) {
	if protocol.getRoundHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetRound not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	roundID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.getRoundHandler(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getRoundHandler(nil, client, callID, roundID)
}
