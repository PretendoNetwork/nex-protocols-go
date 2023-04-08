package datastore_nintendo_badge_arcade

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetaByOwnerId sets the GetMetaByOwnerId function
func (protocol *DataStoreNintendoBadgeArcadeProtocol) GetMetaByOwnerId(handler func(err error, client *nex.Client, callID uint32, param *DataStoreGetMetaByOwnerIdParam)) {
	protocol.GetMetaByOwnerIdHandler = handler
}

func (protocol *DataStoreNintendoBadgeArcadeProtocol) HandleGetMetaByOwnerId(packet nex.PacketInterface) {
	if protocol.GetMetaByOwnerIdHandler == nil {
		globals.Logger.Warning("DataStoreBadgeArcade::GetMetaByOwnerId not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStoreGetMetaByOwnerIdParam())
	if err != nil {
		go protocol.GetMetaByOwnerIdHandler(err, client, callID, nil)
		return
	}

	go protocol.GetMetaByOwnerIdHandler(nil, client, callID, param.(*DataStoreGetMetaByOwnerIdParam))
}
