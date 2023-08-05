// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPlayingSession sets the GetPlayingSession handler function
func (protocol *Protocol) GetPlayingSession(handler func(err error, client *nex.Client, callID uint32, lstPID []uint32) uint32) {
	protocol.getPlayingSessionHandler = handler
}

func (protocol *Protocol) handleGetPlayingSession(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPlayingSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetPlayingSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstPID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getPlayingSessionHandler(fmt.Errorf("Failed to read lstPID from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPlayingSessionHandler(nil, client, callID, lstPID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
