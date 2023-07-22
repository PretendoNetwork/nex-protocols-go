// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindMatchmakeSessionByOwner sets the FindMatchmakeSessionByOwner handler function
func (protocol *MatchmakeExtensionProtocol) FindMatchmakeSessionByOwner(handler func(err error, client *nex.Client, callID uint32, id uint32, resultRange *nex.ResultRange)) {
	protocol.findMatchmakeSessionByOwnerHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleFindMatchmakeSessionByOwner(packet nex.PacketInterface) {
	if protocol.findMatchmakeSessionByOwnerHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindMatchmakeSessionByOwner not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.findMatchmakeSessionByOwnerHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findMatchmakeSessionByOwnerHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.findMatchmakeSessionByOwnerHandler(nil, client, callID, id, resultRange.(*nex.ResultRange))
}
