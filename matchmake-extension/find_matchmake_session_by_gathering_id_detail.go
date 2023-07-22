// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindMatchmakeSessionByGatheringIDDetail sets the FindMatchmakeSessionByGatheringIDDetail handler function
func (protocol *MatchmakeExtensionProtocol) FindMatchmakeSessionByGatheringIDDetail(handler func(err error, client *nex.Client, callID uint32, gid uint32)) {
	protocol.findMatchmakeSessionByGatheringIDDetailHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleFindMatchmakeSessionByGatheringIDDetail(packet nex.PacketInterface) {
	if protocol.findMatchmakeSessionByGatheringIDDetailHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindMatchmakeSessionByGatheringIDDetail not implemented")
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
		go protocol.findMatchmakeSessionByGatheringIDDetailHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.findMatchmakeSessionByGatheringIDDetailHandler(nil, client, callID, gid)
}
