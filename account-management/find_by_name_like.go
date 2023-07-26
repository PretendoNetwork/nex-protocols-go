// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByNameLike sets the FindByNameLike handler function
func (protocol *Protocol) FindByNameLike(handler func(err error, client *nex.Client, callID uint32, uiGroups uint32, strLike string, resultRange *nex.ResultRange)) {
	protocol.findByNameLikeHandler = handler
}

func (protocol *Protocol) handleFindByNameLike(packet nex.PacketInterface) {
	if protocol.findByNameLikeHandler == nil {
		globals.Logger.Warning("AccountManagement::FindByNameLike not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.findByNameLikeHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), client, callID, 0, "", nil)
		return
	}

	strLike, err := parametersStream.ReadString()
	if err != nil {
		go protocol.findByNameLikeHandler(fmt.Errorf("Failed to read strLike from parameters. %s", err.Error()), client, callID, 0, "", nil)
		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findByNameLikeHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, 0, "", nil)
		return
	}

	go protocol.findByNameLikeHandler(nil, client, callID, uiGroups, strLike, resultRange.(*nex.ResultRange))
}
