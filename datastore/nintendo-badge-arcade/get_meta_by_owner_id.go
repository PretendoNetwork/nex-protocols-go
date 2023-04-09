package datastore_nintendo_badge_arcade

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetaByOwnerID sets the GetMetaByOwnerID function
func (protocol *DataStoreNintendoBadgeArcadeProtocol) GetMetaByOwnerID(handler func(err error, client *nex.Client, callID uint32, param *DataStoreGetMetaByOwnerIDParam)) {
	protocol.GetMetaByOwnerIDHandler = handler
}

func (protocol *DataStoreNintendoBadgeArcadeProtocol) HandleGetMetaByOwnerID(packet nex.PacketInterface) {
	if protocol.GetMetaByOwnerIDHandler == nil {
		globals.Logger.Warning("DataStoreBadgeArcade::GetMetaByOwnerID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStoreGetMetaByOwnerIDParam())
	if err != nil {
		go protocol.GetMetaByOwnerIDHandler(err, client, callID, nil)
		return
	}

	go protocol.GetMetaByOwnerIDHandler(nil, client, callID, param.(*DataStoreGetMetaByOwnerIDParam))
}
