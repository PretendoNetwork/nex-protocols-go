// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LaunchSession sets the LaunchSession handler function
func (protocol *Protocol) LaunchSession(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, strURL string)) {
	protocol.launchSessionHandler = handler
}

func (protocol *Protocol) handleLaunchSession(packet nex.PacketInterface) {
	if protocol.launchSessionHandler == nil {
		globals.Logger.Warning("MatchMaking::LaunchSession not implemented")
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
		go protocol.launchSessionHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, "")
	}

	strURL, err := parametersStream.ReadString()
	if err != nil {
		go protocol.launchSessionHandler(fmt.Errorf("Failed to read strURL from parameters. %s", err.Error()), client, callID, 0, "")
	}

	go protocol.launchSessionHandler(nil, client, callID, idGathering, strURL)
}
