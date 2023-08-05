// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchSimpleSearchObjectByObjectIDs sets the SearchSimpleSearchObjectByObjectIDs handler function
func (protocol *Protocol) SearchSimpleSearchObjectByObjectIDs(handler func(err error, client *nex.Client, callID uint32, objectIDs []uint32) uint32) {
	protocol.searchSimpleSearchObjectByObjectIDsHandler = handler
}

func (protocol *Protocol) handleSearchSimpleSearchObjectByObjectIDs(packet nex.PacketInterface) {
	if protocol.searchSimpleSearchObjectByObjectIDsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::SearchSimpleSearchObjectByObjectIDs not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	objectIDs, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.searchSimpleSearchObjectByObjectIDsHandler(fmt.Errorf("Failed to read objectIDs from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.searchSimpleSearchObjectByObjectIDsHandler(nil, client, callID, objectIDs)
}
