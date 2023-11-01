// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LaunchSession sets the LaunchSession handler function
func (protocol *Protocol) LaunchSession(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strURL string) uint32) {
	protocol.launchSessionHandler = handler
}

func (protocol *Protocol) handleLaunchSession(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.launchSessionHandler == nil {
		globals.Logger.Warning("MatchMaking::LaunchSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.launchSessionHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strURL, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.launchSessionHandler(fmt.Errorf("Failed to read strURL from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.launchSessionHandler(nil, packet, callID, idGathering, strURL)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
