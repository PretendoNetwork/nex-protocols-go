// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindMatchmakeSessionByGatheringID sets the FindMatchmakeSessionByGatheringID handler function
func (protocol *Protocol) FindMatchmakeSessionByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID []uint32) uint32) {
	protocol.findMatchmakeSessionByGatheringIDHandler = handler
}

func (protocol *Protocol) handleFindMatchmakeSessionByGatheringID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findMatchmakeSessionByGatheringIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindMatchmakeSessionByGatheringID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.findMatchmakeSessionByGatheringIDHandler(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findMatchmakeSessionByGatheringIDHandler(nil, packet, callID, lstGID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
