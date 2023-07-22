// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"encoding/hex"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AutoMatchmakeWithGatheringID_Postpone sets the AutoMatchmakeWithGatheringID_Postpone handler function
func (protocol *MatchmakeExtensionProtocol) AutoMatchmakeWithGatheringID_Postpone(handler func(err error, client *nex.Client, callID uint32, lstGID []uint32, anyGathering *nex.DataHolder, strMessage string)) {
	protocol.autoMatchmakeWithGatheringID_PostponeHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleAutoMatchmakeWithGatheringID_Postpone(packet nex.PacketInterface) {
	if protocol.autoMatchmakeWithGatheringID_PostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithGatheringID_Postpone not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
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
		go protocol.autoMatchmakeWithGatheringID_PostponeHandler(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.autoMatchmakeWithGatheringID_PostponeHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.autoMatchmakeWithGatheringID_PostponeHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, nil, nil, "")
		return
	}

	go protocol.autoMatchmakeWithGatheringID_PostponeHandler(nil, client, callID, lstGID, anyGathering, strMessage)
}
