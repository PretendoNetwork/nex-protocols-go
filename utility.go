package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

// UniqueIDInfo holds parameters for a matchmake session
type UniqueIDInfo struct {
	NexUniqueId         uint64
	NexUniqueIdPassword uint64

	*nex.Structure
}

// Bytes encodes the UniqueIDInfo and returns a byte array
func (uniqueIDInfo *UniqueIDInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(uniqueIDInfo.NexUniqueId)
	stream.WriteUInt64LE(uniqueIDInfo.NexUniqueIdPassword)

	return stream.Bytes()
}

// ExtractFromStream extracts a UniqueIDInfo structure from a stream
func (uniqueIDInfo *UniqueIDInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	uniqueIDInfo.NexUniqueId = stream.ReadUInt64LE();
	uniqueIDInfo.NexUniqueIdPassword = stream.ReadUInt64LE();

	if err != nil {
		return err
	}

	return nil
}

// NewUniqueIDInfo returns a new UniqueIDInfo
func NewUniqueIDInfo() *UniqueIDInfo {
	return &UniqueIDInfo{}
}

const (
	// UtilityProtocolID is the protocol ID for the Utility protocol
	UtilityProtocolID = 0x6e

	// UtilityMethodAcquireNexUniqueId is the method ID for the method AcquireNexUniqueId
	UtilityMethodAcquireNexUniqueId = 0x1

	// UtilityMethodAcquireNexUniqueIdWithPassword is the method ID for the method AcquireNexUniqueIdWithPassword
	UtilityMethodAcquireNexUniqueIdWithPassword = 0x2

	// UtilityMethodAssociateNexUniqueIdWithMyPrincipalId is the method ID for the method AssociateNexUniqueIdWithMyPrincipalId
	UtilityMethodAssociateNexUniqueIdWithMyPrincipalId = 0x3

	// UtilityMethodAssociateNexUniqueIdsWithMyPrincipalId is the method ID for the method AssociateNexUniqueIdsWithMyPrincipalId
	UtilityMethodAssociateNexUniqueIdsWithMyPrincipalId = 0x4

	// UtilityMethodGetAssociatedNexUniqueIdWithMyPrincipalId is the method ID for the method GetAssociatedNexUniqueIdWithMyPrincipalId
	UtilityMethodGetAssociatedNexUniqueIdWithMyPrincipalId = 0x5

	// UtilityMethodGetAssociatedNexUniqueIdsWithMyPrincipalId is the method ID for the method GetAssociatedNexUniqueIdsWithMyPrincipalId
	UtilityMethodGetAssociatedNexUniqueIdsWithMyPrincipalId = 0x6

	// UtilityMethodGetIntegerSettings is the method ID for the method GetIntegerSettings
	UtilityMethodGetIntegerSettings = 0x7

	// UtilityMethodGetStringSettings is the method ID for the method GetStringSettings
	UtilityMethodGetStringSettings = 0x8
)

// UtilityProtocol handles the Utility nex protocol
type UtilityProtocol struct {
	server                                            *nex.Server
	AcquireNexUniqueIdHandler                         func(err error, client *nex.Client, callID uint32)
	AcquireNexUniqueIdWithPasswordHandler             func(err error, client *nex.Client, callID uint32)
	AssociateNexUniqueIdWithMyPrincipalIdHandler      func(err error, client *nex.Client, callID uint32, uniqueIDInfo *UniqueIDInfo)
	AssociateNexUniqueIdsWithMyPrincipalIdHandler     func(err error, client *nex.Client, callID uint32, uniqueIDInfo []*UniqueIDInfo)
	GetAssociatedNexUniqueIdWithMyPrincipalIdHandler  func(err error, client *nex.Client, callID uint32)
	GetAssociatedNexUniqueIdsWithMyPrincipalIdHandler func(err error, client *nex.Client, callID uint32)
	GetIntegerSettingsHandler                         func(err error, client *nex.Client, callID uint32, integerSettingIndex uint32)
	GetStringSettingsHandler                          func(err error, client *nex.Client, callID uint32, stringSettingIndex uint32)
}

// Setup initializes the protocol
func (utilityProtocol *UtilityProtocol) Setup() {
	nexServer := utilityProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if UtilityProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case UtilityMethodAcquireNexUniqueId:
				go utilityProtocol.handleAcquireNexUniqueId(packet)
			case UtilityMethodAcquireNexUniqueIdWithPassword:
				go utilityProtocol.handleAcquireNexUniqueIdWithPassword(packet)
			case UtilityMethodAssociateNexUniqueIdWithMyPrincipalId:
				go utilityProtocol.handleAssociateNexUniqueIdWithMyPrincipalId(packet)
			case UtilityMethodAssociateNexUniqueIdsWithMyPrincipalId:
				go utilityProtocol.handleAssociateNexUniqueIdsWithMyPrincipalId(packet)
			case UtilityMethodGetAssociatedNexUniqueIdWithMyPrincipalId:
				go utilityProtocol.handleGetAssociatedNexUniqueIdWithMyPrincipalId(packet)
			case UtilityMethodGetIntegerSettings:
				go utilityProtocol.handleGetIntegerSettings(packet)
			case UtilityMethodGetStringSettings:
				go utilityProtocol.handleGetStringSettings(packet)
			default:
				go respondNotImplemented(packet, UtilityProtocolID)
				fmt.Printf("Unsupported Utility method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// AcquireNexUniqueId sets the AcquireNexUniqueId handler function
func (utilityProtocol *UtilityProtocol) AcquireNexUniqueId(handler func(err error, client *nex.Client, callID uint32)) {
	utilityProtocol.AcquireNexUniqueIdHandler = handler
}

// AcquireNexUniqueIdWithPassword sets the AcquireNexUniqueIdWithPassword handler function
func (utilityProtocol *UtilityProtocol) AcquireNexUniqueIdWithPassword(handler func(err error, client *nex.Client, callID uint32)) {
	utilityProtocol.AcquireNexUniqueIdWithPasswordHandler = handler
}

// AssociateNexUniqueIdWithMyPrincipalId sets the AssociateNexUniqueIdWithMyPrincipalId handler function
func (utilityProtocol *UtilityProtocol) AssociateNexUniqueIdWithMyPrincipalId(handler func(err error, client *nex.Client, callID uint32, uniqueIDInfo *UniqueIDInfo)) {
	utilityProtocol.AssociateNexUniqueIdWithMyPrincipalIdHandler = handler
}

// AssociateNexUniqueIdsWithMyPrincipalId sets the AssociateNexUniqueIdsWithMyPrincipalId handler function
func (utilityProtocol *UtilityProtocol) AssociateNexUniqueIdsWithMyPrincipalId(handler func(err error, client *nex.Client, callID uint32, uniqueIDInfo []*UniqueIDInfo)) {
	utilityProtocol.AssociateNexUniqueIdsWithMyPrincipalIdHandler = handler
}

// GetAssociatedNexUniqueIdWithMyPrincipalId sets the GetAssociatedNexUniqueIdWithMyPrincipalId handler function
func (utilityProtocol *UtilityProtocol) GetAssociatedNexUniqueIdWithMyPrincipalId(handler func(err error, client *nex.Client, callID uint32)) {
	utilityProtocol.GetAssociatedNexUniqueIdWithMyPrincipalIdHandler = handler
}

// GetAssociatedNexUniqueIdsWithMyPrincipalId sets the GetAssociatedNexUniqueIdsWithMyPrincipalId handler function
func (utilityProtocol *UtilityProtocol) GetAssociatedNexUniqueIdsWithMyPrincipalId(handler func(err error, client *nex.Client, callID uint32)) {
	utilityProtocol.GetAssociatedNexUniqueIdsWithMyPrincipalIdHandler = handler
}

// GetIntegerSettings sets the GetIntegerSettings handler function
func (utilityProtocol *UtilityProtocol) GetIntegerSettings(handler func(err error, client *nex.Client, callID uint32, integerSettingIndex uint32)) {
	utilityProtocol.GetIntegerSettingsHandler = handler
}

// GetStringSettings sets the GetStringSettings handler function
func (utilityProtocol *UtilityProtocol) GetStringSettings(handler func(err error, client *nex.Client, callID uint32, stringSettingIndex uint32)) {
	utilityProtocol.GetStringSettingsHandler = handler
}

func (utilityProtocol *UtilityProtocol) handleAcquireNexUniqueId(packet nex.PacketInterface) {
	if utilityProtocol.AcquireNexUniqueIdHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::AcquireNexUniqueId not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go utilityProtocol.AcquireNexUniqueIdHandler(nil, client, callID)
}

func (utilityProtocol *UtilityProtocol) handleAcquireNexUniqueIdWithPassword(packet nex.PacketInterface) {
	if utilityProtocol.AcquireNexUniqueIdWithPasswordHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::AcquireNexUniqueIdWithPassword not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go utilityProtocol.AcquireNexUniqueIdWithPasswordHandler(nil, client, callID)
}

func (utilityProtocol *UtilityProtocol) handleAssociateNexUniqueIdWithMyPrincipalId(packet nex.PacketInterface) {
	if utilityProtocol.AssociateNexUniqueIdWithMyPrincipalIdHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::AssociateNexUniqueIdWithMyPrincipalId not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, utilityProtocol.server)

	uniqueIDInfoStructureInterface, err := parametersStream.ReadStructure(NewUniqueIDInfo())
	if err != nil {
		go utilityProtocol.AssociateNexUniqueIdWithMyPrincipalIdHandler(nil, client, callID, nil)
		return
	}
	uniqueIDInfo := uniqueIDInfoStructureInterface.(*UniqueIDInfo)

	go utilityProtocol.AssociateNexUniqueIdWithMyPrincipalIdHandler(nil, client, callID, uniqueIDInfo)
}

func (utilityProtocol *UtilityProtocol) handleAssociateNexUniqueIdsWithMyPrincipalId(packet nex.PacketInterface) {
	if utilityProtocol.GetAssociatedNexUniqueIdWithMyPrincipalIdHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::AssociateNexUniqueIdsWithMyPrincipalId not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, utilityProtocol.server)
	structureCount := (int)(parametersStream.ReadUInt32LE());
	uniqueIDInfo := make([]*UniqueIDInfo, structureCount)
	
	for i := 0; i < structureCount; i++ {
		uniqueIDInfoStructureInterface, err := parametersStream.ReadStructure(NewUniqueIDInfo())
		if err != nil {
			go utilityProtocol.AssociateNexUniqueIdsWithMyPrincipalIdHandler(nil, client, callID, nil)
			return
		}
		uniqueIDInfo[i] = uniqueIDInfoStructureInterface.(*UniqueIDInfo)
	}

	go utilityProtocol.AssociateNexUniqueIdsWithMyPrincipalIdHandler(nil, client, callID, uniqueIDInfo)
}

func (utilityProtocol *UtilityProtocol) handleGetAssociatedNexUniqueIdWithMyPrincipalId(packet nex.PacketInterface) {
	if utilityProtocol.GetAssociatedNexUniqueIdWithMyPrincipalIdHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::GetAssociatedNexUniqueIdWithMyPrincipalId not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go utilityProtocol.GetAssociatedNexUniqueIdWithMyPrincipalIdHandler(nil, client, callID)
}

func (utilityProtocol *UtilityProtocol) handleGetAssociatedNexUniqueIdsWithMyPrincipalId(packet nex.PacketInterface) {
	if utilityProtocol.GetAssociatedNexUniqueIdsWithMyPrincipalIdHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::GetAssociatedNexUniqueIdsWithMyPrincipalId not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go utilityProtocol.GetAssociatedNexUniqueIdsWithMyPrincipalIdHandler(nil, client, callID)
}

func (utilityProtocol *UtilityProtocol) handleGetIntegerSettings(packet nex.PacketInterface) {
	if utilityProtocol.GetIntegerSettingsHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::GetIntegerSettings not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, utilityProtocol.server)
	integerSettingIndex := parametersStream.ReadUInt32LE();

	go utilityProtocol.GetIntegerSettingsHandler(nil, client, callID, integerSettingIndex)
}

func (utilityProtocol *UtilityProtocol) handleGetStringSettings(packet nex.PacketInterface) {
	if utilityProtocol.GetStringSettingsHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::GetStringSettings not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, utilityProtocol.server)
	stringSettingIndex := parametersStream.ReadUInt32LE();

	go utilityProtocol.GetStringSettingsHandler(nil, client, callID, stringSettingIndex)
}

// NewUtilityProtocol returns a new UtilityProtocol
func NewUtilityProtocol(server *nex.Server) *UtilityProtocol {
	utilityProtocol := &UtilityProtocol{server: server}

	utilityProtocol.Setup()

	return utilityProtocol
}
