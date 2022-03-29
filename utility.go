package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

// UniqueIDInfo holds parameters for a matchmake session
type UniqueIDInfo struct {
	NexUniqueID         uint64
	NexUniqueIDPassword uint64

	*nex.Structure
}

// Bytes encodes the UniqueIDInfo and returns a byte array
func (uniqueIDInfo *UniqueIDInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(uniqueIDInfo.NexUniqueID)
	stream.WriteUInt64LE(uniqueIDInfo.NexUniqueIDPassword)

	return stream.Bytes()
}

// ExtractFromStream extracts a UniqueIDInfo structure from a stream
func (uniqueIDInfo *UniqueIDInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	uniqueIDInfo.NexUniqueID = stream.ReadUInt64LE()
	uniqueIDInfo.NexUniqueIDPassword = stream.ReadUInt64LE()

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

	// UtilityMethodAcquireNexUniqueID is the method ID for the method AcquireNexUniqueID
	UtilityMethodAcquireNexUniqueID = 0x1

	// UtilityMethodAcquireNexUniqueIDWithPassword is the method ID for the method AcquireNexUniqueIDWithPassword
	UtilityMethodAcquireNexUniqueIDWithPassword = 0x2

	// UtilityMethodAssociateNexUniqueIDWithMyPrincipalID is the method ID for the method AssociateNexUniqueIDWithMyPrincipalID
	UtilityMethodAssociateNexUniqueIDWithMyPrincipalID = 0x3

	// UtilityMethodAssociateNexUniqueIDsWithMyPrincipalID is the method ID for the method AssociateNexUniqueIDsWithMyPrincipalID
	UtilityMethodAssociateNexUniqueIDsWithMyPrincipalID = 0x4

	// UtilityMethodGetAssociatedNexUniqueIDWithMyPrincipalID is the method ID for the method GetAssociatedNexUniqueIDWithMyPrincipalID
	UtilityMethodGetAssociatedNexUniqueIDWithMyPrincipalID = 0x5

	// UtilityMethodGetAssociatedNexUniqueIDsWithMyPrincipalID is the method ID for the method GetAssociatedNexUniqueIDsWithMyPrincipalID
	UtilityMethodGetAssociatedNexUniqueIDsWithMyPrincipalID = 0x6

	// UtilityMethodGetIntegerSettings is the method ID for the method GetIntegerSettings
	UtilityMethodGetIntegerSettings = 0x7

	// UtilityMethodGetStringSettings is the method ID for the method GetStringSettings
	UtilityMethodGetStringSettings = 0x8
)

// UtilityProtocol handles the Utility nex protocol
type UtilityProtocol struct {
	server                                            *nex.Server
	AcquireNexUniqueIDHandler                         func(err error, client *nex.Client, callID uint32)
	AcquireNexUniqueIDWithPasswordHandler             func(err error, client *nex.Client, callID uint32)
	AssociateNexUniqueIDWithMyPrincipalIDHandler      func(err error, client *nex.Client, callID uint32, uniqueIDInfo *UniqueIDInfo)
	AssociateNexUniqueIDsWithMyPrincipalIDHandler     func(err error, client *nex.Client, callID uint32, uniqueIDInfo []*UniqueIDInfo)
	GetAssociatedNexUniqueIDWithMyPrincipalIDHandler  func(err error, client *nex.Client, callID uint32)
	GetAssociatedNexUniqueIDsWithMyPrincipalIDHandler func(err error, client *nex.Client, callID uint32)
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
			case UtilityMethodAcquireNexUniqueID:
				go utilityProtocol.handleAcquireNexUniqueID(packet)
			case UtilityMethodAcquireNexUniqueIDWithPassword:
				go utilityProtocol.handleAcquireNexUniqueIDWithPassword(packet)
			case UtilityMethodAssociateNexUniqueIDWithMyPrincipalID:
				go utilityProtocol.handleAssociateNexUniqueIDWithMyPrincipalID(packet)
			case UtilityMethodAssociateNexUniqueIDsWithMyPrincipalID:
				go utilityProtocol.handleAssociateNexUniqueIDsWithMyPrincipalID(packet)
			case UtilityMethodGetAssociatedNexUniqueIDWithMyPrincipalID:
				go utilityProtocol.handleGetAssociatedNexUniqueIDWithMyPrincipalID(packet)
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

// AcquireNexUniqueID sets the AcquireNexUniqueID handler function
func (utilityProtocol *UtilityProtocol) AcquireNexUniqueID(handler func(err error, client *nex.Client, callID uint32)) {
	utilityProtocol.AcquireNexUniqueIDHandler = handler
}

// AcquireNexUniqueIDWithPassword sets the AcquireNexUniqueIDWithPassword handler function
func (utilityProtocol *UtilityProtocol) AcquireNexUniqueIDWithPassword(handler func(err error, client *nex.Client, callID uint32)) {
	utilityProtocol.AcquireNexUniqueIDWithPasswordHandler = handler
}

// AssociateNexUniqueIDWithMyPrincipalID sets the AssociateNexUniqueIDWithMyPrincipalID handler function
func (utilityProtocol *UtilityProtocol) AssociateNexUniqueIDWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32, uniqueIDInfo *UniqueIDInfo)) {
	utilityProtocol.AssociateNexUniqueIDWithMyPrincipalIDHandler = handler
}

// AssociateNexUniqueIDsWithMyPrincipalID sets the AssociateNexUniqueIDsWithMyPrincipalID handler function
func (utilityProtocol *UtilityProtocol) AssociateNexUniqueIDsWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32, uniqueIDInfo []*UniqueIDInfo)) {
	utilityProtocol.AssociateNexUniqueIDsWithMyPrincipalIDHandler = handler
}

// GetAssociatedNexUniqueIDWithMyPrincipalID sets the GetAssociatedNexUniqueIDWithMyPrincipalID handler function
func (utilityProtocol *UtilityProtocol) GetAssociatedNexUniqueIDWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32)) {
	utilityProtocol.GetAssociatedNexUniqueIDWithMyPrincipalIDHandler = handler
}

// GetAssociatedNexUniqueIDsWithMyPrincipalID sets the GetAssociatedNexUniqueIDsWithMyPrincipalID handler function
func (utilityProtocol *UtilityProtocol) GetAssociatedNexUniqueIDsWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32)) {
	utilityProtocol.GetAssociatedNexUniqueIDsWithMyPrincipalIDHandler = handler
}

// GetIntegerSettings sets the GetIntegerSettings handler function
func (utilityProtocol *UtilityProtocol) GetIntegerSettings(handler func(err error, client *nex.Client, callID uint32, integerSettingIndex uint32)) {
	utilityProtocol.GetIntegerSettingsHandler = handler
}

// GetStringSettings sets the GetStringSettings handler function
func (utilityProtocol *UtilityProtocol) GetStringSettings(handler func(err error, client *nex.Client, callID uint32, stringSettingIndex uint32)) {
	utilityProtocol.GetStringSettingsHandler = handler
}

func (utilityProtocol *UtilityProtocol) handleAcquireNexUniqueID(packet nex.PacketInterface) {
	if utilityProtocol.AcquireNexUniqueIDHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::AcquireNexUniqueID not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go utilityProtocol.AcquireNexUniqueIDHandler(nil, client, callID)
}

func (utilityProtocol *UtilityProtocol) handleAcquireNexUniqueIDWithPassword(packet nex.PacketInterface) {
	if utilityProtocol.AcquireNexUniqueIDWithPasswordHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::AcquireNexUniqueIDWithPassword not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go utilityProtocol.AcquireNexUniqueIDWithPasswordHandler(nil, client, callID)
}

func (utilityProtocol *UtilityProtocol) handleAssociateNexUniqueIDWithMyPrincipalID(packet nex.PacketInterface) {
	if utilityProtocol.AssociateNexUniqueIDWithMyPrincipalIDHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::AssociateNexUniqueIDWithMyPrincipalID not implemented")
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
		go utilityProtocol.AssociateNexUniqueIDWithMyPrincipalIDHandler(nil, client, callID, nil)
		return
	}
	uniqueIDInfo := uniqueIDInfoStructureInterface.(*UniqueIDInfo)

	go utilityProtocol.AssociateNexUniqueIDWithMyPrincipalIDHandler(nil, client, callID, uniqueIDInfo)
}

func (utilityProtocol *UtilityProtocol) handleAssociateNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	if utilityProtocol.GetAssociatedNexUniqueIDWithMyPrincipalIDHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::AssociateNexUniqueIDsWithMyPrincipalID not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, utilityProtocol.server)
	structureCount := (int)(parametersStream.ReadUInt32LE())
	uniqueIDInfo := make([]*UniqueIDInfo, structureCount)

	for i := 0; i < structureCount; i++ {
		uniqueIDInfoStructureInterface, err := parametersStream.ReadStructure(NewUniqueIDInfo())
		if err != nil {
			go utilityProtocol.AssociateNexUniqueIDsWithMyPrincipalIDHandler(nil, client, callID, nil)
			return
		}
		uniqueIDInfo[i] = uniqueIDInfoStructureInterface.(*UniqueIDInfo)
	}

	go utilityProtocol.AssociateNexUniqueIDsWithMyPrincipalIDHandler(nil, client, callID, uniqueIDInfo)
}

func (utilityProtocol *UtilityProtocol) handleGetAssociatedNexUniqueIDWithMyPrincipalID(packet nex.PacketInterface) {
	if utilityProtocol.GetAssociatedNexUniqueIDWithMyPrincipalIDHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::GetAssociatedNexUniqueIDWithMyPrincipalID not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go utilityProtocol.GetAssociatedNexUniqueIDWithMyPrincipalIDHandler(nil, client, callID)
}

func (utilityProtocol *UtilityProtocol) handleGetAssociatedNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	if utilityProtocol.GetAssociatedNexUniqueIDsWithMyPrincipalIDHandler == nil {
		fmt.Println("[Warning] UtilityProtocol::GetAssociatedNexUniqueIDsWithMyPrincipalID not implemented")
		go respondNotImplemented(packet, UtilityProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go utilityProtocol.GetAssociatedNexUniqueIDsWithMyPrincipalIDHandler(nil, client, callID)
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
	integerSettingIndex := parametersStream.ReadUInt32LE()

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
	stringSettingIndex := parametersStream.ReadUInt32LE()

	go utilityProtocol.GetStringSettingsHandler(nil, client, callID, stringSettingIndex)
}

// NewUtilityProtocol returns a new UtilityProtocol
func NewUtilityProtocol(server *nex.Server) *UtilityProtocol {
	utilityProtocol := &UtilityProtocol{server: server}

	utilityProtocol.Setup()

	return utilityProtocol
}
