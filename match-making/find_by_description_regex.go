// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByDescriptionRegex sets the FindByDescriptionRegex handler function
func (protocol *Protocol) FindByDescriptionRegex(handler func(err error, client *nex.Client, callID uint32, strDescriptionRegex string, resultRange *nex.ResultRange) uint32) {
	protocol.findByDescriptionRegexHandler = handler
}

func (protocol *Protocol) handleFindByDescriptionRegex(packet nex.PacketInterface) {
	if protocol.findByDescriptionRegexHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByDescriptionRegex not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
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
