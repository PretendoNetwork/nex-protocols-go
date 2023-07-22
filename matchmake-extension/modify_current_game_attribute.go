// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ModifyCurrentGameAttribute sets the ModifyCurrentGameAttribute handler function
func (protocol *MatchmakeExtensionProtocol) ModifyCurrentGameAttribute(handler func(err error, client *nex.Client, callID uint32, gid uint32, attribIndex uint32, newValue uint32)) {
	protocol.modifyCurrentGameAttributeHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleModifyCurrentGameAttribute(packet nex.PacketInterface) {
	if protocol.modifyCurrentGameAttributeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::ModifyCurrentGameAttribute not implemented")
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
		go protocol.modifyCurrentGameAttributeHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		return
	}

	attribIndex, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.modifyCurrentGameAttributeHandler(fmt.Errorf("Failed to read attribIndex from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		return
	}

	newValue, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.modifyCurrentGameAttributeHandler(fmt.Errorf("Failed to read newValue from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		return
	}

	go protocol.modifyCurrentGameAttributeHandler(nil, client, callID, gid, attribIndex, newValue)
}
