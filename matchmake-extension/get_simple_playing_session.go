// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSimplePlayingSession sets the GetSimplePlayingSession handler function
func (protocol *Protocol) GetSimplePlayingSession(handler func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool) uint32) {
	protocol.getSimplePlayingSessionHandler = handler
}

func (protocol *Protocol) handleGetSimplePlayingSession(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getSimplePlayingSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetSimplePlayingSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	listPID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getSimplePlayingSessionHandler(fmt.Errorf("Failed to read listPID from parameters. %s", err.Error()), client, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	includeLoginUser, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.getSimplePlayingSessionHandler(fmt.Errorf("Failed to read includeLoginUser from parameters. %s", err.Error()), client, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getSimplePlayingSessionHandler(nil, client, callID, listPID, includeLoginUser)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
