// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityByGatheringID sets the FindCommunityByGatheringID handler function
func (protocol *Protocol) FindCommunityByGatheringID(handler func(err error, client *nex.Client, callID uint32, lstGID []uint32) uint32) {
	protocol.findCommunityByGatheringIDHandler = handler
}

func (protocol *Protocol) handleFindCommunityByGatheringID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findCommunityByGatheringIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindCommunityByGatheringID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.findCommunityByGatheringIDHandler(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findCommunityByGatheringIDHandler(nil, client, callID, lstGID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
