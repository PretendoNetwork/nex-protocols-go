// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByDescriptionRegex sets the FindByDescriptionRegex handler function
func (protocol *MatchMakingProtocol) FindByDescriptionRegex(handler func(err error, client *nex.Client, callID uint32, strDescriptionRegex string, resultRange *nex.ResultRange)) {
	protocol.findByDescriptionRegexHandler = handler
}

func (protocol *MatchMakingProtocol) handleFindByDescriptionRegex(packet nex.PacketInterface) {
	if protocol.findByDescriptionRegexHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByDescriptionRegex not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strDescriptionRegex, err := parametersStream.ReadString()
	if err != nil {
		go protocol.findByDescriptionRegexHandler(fmt.Errorf("Failed to read strDescriptionRegex from parameters. %s", err.Error()), client, callID, "", nil)
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findByDescriptionRegexHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, "", nil)
	}

	go protocol.findByDescriptionRegexHandler(nil, client, callID, strDescriptionRegex, resultRange.(*nex.ResultRange))
}
