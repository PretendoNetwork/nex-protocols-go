// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMatchmakeSessionAttribute sets the UpdateMatchmakeSessionAttribute handler function
func (protocol *Protocol) UpdateMatchmakeSessionAttribute(handler func(err error, client *nex.Client, callID uint32, gid uint32, attribs []uint32)) {
	protocol.updateMatchmakeSessionAttributeHandler = handler
}

func (protocol *Protocol) handleUpdateMatchmakeSessionAttribute(packet nex.PacketInterface) {
	if protocol.updateMatchmakeSessionAttributeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateMatchmakeSessionAttribute not implemented")
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
		go protocol.updateMatchmakeSessionAttributeHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	attribs, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.updateMatchmakeSessionAttributeHandler(fmt.Errorf("Failed to read attribs from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.updateMatchmakeSessionAttributeHandler(nil, client, callID, gid, attribs)
}
