// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPasswordInfo sets the GetPasswordInfo handler function
func (protocol *Protocol) GetPasswordInfo(handler func(err error, client *nex.Client, callID uint32, dataID uint64)) {
	protocol.getPasswordInfoHandler = handler
}

func (protocol *Protocol) handleGetPasswordInfo(packet nex.PacketInterface) {
	if protocol.getPasswordInfoHandler == nil {
		globals.Logger.Warning("DataStore::GetPasswordInfo not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.getPasswordInfoHandler(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getPasswordInfoHandler(nil, client, callID, dataID)
}