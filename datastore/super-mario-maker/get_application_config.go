package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetApplicationConfig sets the GetApplicationConfig handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetApplicationConfig(handler func(err error, client *nex.Client, callID uint32, applicationID uint32)) {
	protocol.GetApplicationConfigHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleGetApplicationConfig(packet nex.PacketInterface) {
	if protocol.GetApplicationConfigHandler == nil {
		globals.Logger.Warning("DataStoreSMM::GetApplicationConfig not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID := parametersStream.ReadUInt32LE()

	go protocol.GetApplicationConfigHandler(nil, client, callID, applicationID)
}
