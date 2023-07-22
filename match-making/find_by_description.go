// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByDescription sets the FindByDescription handler function
func (protocol *MatchMakingProtocol) FindByDescription(handler func(err error, client *nex.Client, callID uint32, strDescription string, resultRange *nex.ResultRange)) {
	protocol.findByDescriptionHandler = handler
}

func (protocol *MatchMakingProtocol) handleFindByDescription(packet nex.PacketInterface) {
	if protocol.findByDescriptionHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByDescription not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strDescription, err := parametersStream.ReadString()
	if err != nil {
		go protocol.findByDescriptionHandler(fmt.Errorf("Failed to read strDescription from parameters. %s", err.Error()), client, callID, "", nil)
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findByDescriptionHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, "", nil)
	}

	go protocol.findByDescriptionHandler(nil, client, callID, strDescription, resultRange.(*nex.ResultRange))
}
