// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindMatchmakeSessionBySingleGatheringID sets the FindMatchmakeSessionBySingleGatheringID handler function
func (protocol *Protocol) FindMatchmakeSessionBySingleGatheringID(handler func(err error, client *nex.Client, callID uint32, GID uint32)) {
	protocol.findMatchmakeSessionBySingleGatheringIDHandler = handler
}

func (protocol *Protocol) handleFindMatchmakeSessionBySingleGatheringID(packet nex.PacketInterface) {
	if protocol.findMatchmakeSessionBySingleGatheringIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindMatchmakeSessionBySingleGatheringID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.findMatchmakeSessionBySingleGatheringIDHandler(fmt.Errorf("Failed to read GID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.findMatchmakeSessionBySingleGatheringIDHandler(nil, client, callID, gid)
}
