// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"encoding/hex"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AutoMatchmakeWithGatheringIDPostpone sets the AutoMatchmakeWithGatheringIDPostpone handler function
func (protocol *Protocol) AutoMatchmakeWithGatheringIDPostpone(handler func(err error, client *nex.Client, callID uint32, lstGID []uint32, anyGathering *nex.DataHolder, strMessage string) uint32) {
	protocol.autoMatchmakeWithGatheringIDPostponeHandler = handler
}

func (protocol *Protocol) handleAutoMatchmakeWithGatheringIDPostpone(packet nex.PacketInterface) {
	if protocol.autoMatchmakeWithGatheringIDPostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithGatheringIDPostpone not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	globals.Logger.Info(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.autoMatchmakeWithGatheringIDPostponeHandler(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.autoMatchmakeWithGatheringIDPostponeHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.autoMatchmakeWithGatheringIDPostponeHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	go protocol.autoMatchmakeWithGatheringIDPostponeHandler(nil, client, callID, lstGID, anyGathering, strMessage)
}
