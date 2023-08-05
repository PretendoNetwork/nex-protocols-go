// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindMatchmakeSessionByGatheringID sets the FindMatchmakeSessionByGatheringID handler function
func (protocol *Protocol) FindMatchmakeSessionByGatheringID(handler func(err error, client *nex.Client, callID uint32, lstGID []uint32) uint32) {
	protocol.findMatchmakeSessionByGatheringIDHandler = handler
}

func (protocol *Protocol) handleFindMatchmakeSessionByGatheringID(packet nex.PacketInterface) {
	if protocol.findMatchmakeSessionByGatheringIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindMatchmakeSessionByGatheringID not implemented")
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
		go protocol.findMatchmakeSessionByGatheringIDHandler(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.findMatchmakeSessionByGatheringIDHandler(nil, client, callID, lstGID)
}
