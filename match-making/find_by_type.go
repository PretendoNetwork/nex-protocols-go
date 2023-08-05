// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByType sets the FindByType handler function
func (protocol *Protocol) FindByType(handler func(err error, client *nex.Client, callID uint32, strType string, resultRange *nex.ResultRange) uint32) {
	protocol.findByTypeHandler = handler
}

func (protocol *Protocol) handleFindByType(packet nex.PacketInterface) {
	if protocol.findByTypeHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByType not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strType, err := parametersStream.ReadString()
	if err != nil {
		go protocol.findByTypeHandler(fmt.Errorf("Failed to read strType from parameters. %s", err.Error()), client, callID, "", nil)
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findByTypeHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, "", nil)
	}

	go protocol.findByTypeHandler(nil, client, callID, strType, resultRange.(*nex.ResultRange))
}
