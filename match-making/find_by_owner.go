// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByOwner sets the FindByOwner handler function
func (protocol *Protocol) FindByOwner(handler func(err error, client *nex.Client, callID uint32, id uint32, resultRange *nex.ResultRange)) {
	protocol.findByOwnerHandler = handler
}

func (protocol *Protocol) handleFindByOwner(packet nex.PacketInterface) {
	if protocol.findByOwnerHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByOwner not implemented")
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
		go protocol.findByOwnerHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findByOwnerHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	go protocol.findByOwnerHandler(nil, client, callID, id, resultRange.(*nex.ResultRange))
}
