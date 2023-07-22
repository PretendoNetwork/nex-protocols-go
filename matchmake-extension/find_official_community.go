// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindOfficialCommunity sets the FindOfficialCommunity handler function
func (protocol *MatchmakeExtensionProtocol) FindOfficialCommunity(handler func(err error, client *nex.Client, callID uint32, isAvailableOnly bool, resultRange *nex.ResultRange)) {
	protocol.findOfficialCommunityHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleFindOfficialCommunity(packet nex.PacketInterface) {
	if protocol.findOfficialCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindOfficialCommunity not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	isAvailableOnly, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.findOfficialCommunityHandler(fmt.Errorf("Failed to read isAvailableOnly from parameters. %s", err.Error()), client, callID, false, nil)
		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findOfficialCommunityHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, false, nil)
		return
	}

	go protocol.findOfficialCommunityHandler(nil, client, callID, isAvailableOnly, resultRange.(*nex.ResultRange))
}
