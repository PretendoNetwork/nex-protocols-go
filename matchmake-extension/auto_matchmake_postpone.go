// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AutoMatchmakePostpone sets the AutoMatchmakePostpone handler function
func (protocol *Protocol) AutoMatchmakePostpone(handler func(err error, client *nex.Client, callID uint32, anyGathering *nex.DataHolder, strMessage string) uint32) {
	protocol.autoMatchmakePostponeHandler = handler
}

func (protocol *Protocol) handleAutoMatchmakePostpone(packet nex.PacketInterface) {
	if protocol.autoMatchmakePostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakePostpone not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.autoMatchmakePostponeHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), client, callID, nil, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.autoMatchmakePostponeHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, nil, "")
		return
	}

	go protocol.autoMatchmakePostponeHandler(nil, client, callID, anyGathering, strMessage)
}
