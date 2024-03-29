// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item "github.com/PretendoNetwork/nex-protocols-go/service-item"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Service Item (Team Kirby Clash Deluxe) protocol
	ProtocolID = 0x77

	// MethodGetEnvironment is the method ID for the GetEnvironment method
	MethodGetEnvironment = 0x1

	// MethodHTTPGetRequest is the method ID for the HTTPGetRequest method
	MethodHTTPGetRequest = 0x2

	// MethodHTTPGetResponse is the method ID for the HTTPGetResponse method
	MethodHTTPGetResponse = 0x3

	// MethodPurchaseServiceItemRequest is the method ID for the PurchaseServiceItemRequest method
	MethodPurchaseServiceItemRequest = 0x4

	// MethodPurchaseServiceItemResponse is the method ID for the PurchaseServiceItemResponse method
	MethodPurchaseServiceItemResponse = 0x5

	// MethodListServiceItemRequest is the method ID for the ListServiceItemRequest method
	MethodListServiceItemRequest = 0x6

	// MethodListServiceItemResponse is the method ID for the ListServiceItemResponse method
	MethodListServiceItemResponse = 0x7

	// MethodGetBalanceRequest is the method ID for the GetBalanceRequest method
	MethodGetBalanceRequest = 0x8

	// MethodGetBalanceResponse is the method ID for the GetBalanceResponse method
	MethodGetBalanceResponse = 0x9

	// MethodGetPrepurchaseInfoRequest is the method ID for the GetPrepurchaseInfoRequest method
	MethodGetPrepurchaseInfoRequest = 0xA

	// MethodGetPrepurchaseInfoResponse is the method ID for the GetPrepurchaseInfoResponse method
	MethodGetPrepurchaseInfoResponse = 0xB

	// MethodGetServiceItemRightRequest is the method ID for the GetServiceItemRightRequest method
	MethodGetServiceItemRightRequest = 0xC

	// MethodGetServiceItemRightResponse is the method ID for the GetServiceItemRightResponse method
	MethodGetServiceItemRightResponse = 0xD

	// MethodGetPurchaseHistoryRequest is the method ID for the GetPurchaseHistoryRequest method
	MethodGetPurchaseHistoryRequest = 0xE

	// MethodGetPurchaseHistoryResponse is the method ID for the GetPurchaseHistoryResponse method
	MethodGetPurchaseHistoryResponse = 0xF

	// MethodPostRightBinaryByAccount is the method ID for the PostRightBinaryByAccount method
	MethodPostRightBinaryByAccount = 0x10

	// MethodUseServiceItemByAccountRequest is the method ID for the UseServiceItemByAccountRequest method
	MethodUseServiceItemByAccountRequest = 0x11

	// MethodUseServiceItemByAccountResponse is the method ID for the UseServiceItemByAccountResponse method
	MethodUseServiceItemByAccountResponse = 0x12

	// MethodAcquireServiceItemByAccount is the method ID for the AcquireServiceItemByAccount method
	MethodAcquireServiceItemByAccount = 0x13

	// MethodGetSupportID is the method ID for the GetSupportID method
	MethodGetSupportID = 0x14

	// MethodGetLawMessageRequest is the method ID for the GetLawMessageRequest method
	MethodGetLawMessageRequest = 0x15

	// MethodGetLawMessageResponse is the method ID for the GetLawMessageResponse method
	MethodGetLawMessageResponse = 0x16
)

var patchedMethods = []uint32{
	MethodGetEnvironment,
	MethodHTTPGetRequest,
	MethodHTTPGetResponse,
	MethodPurchaseServiceItemRequest,
	MethodPurchaseServiceItemResponse,
	MethodListServiceItemRequest,
	MethodListServiceItemResponse,
	MethodGetBalanceRequest,
	MethodGetBalanceResponse,
	MethodGetPrepurchaseInfoRequest,
	MethodGetPrepurchaseInfoResponse,
	MethodGetServiceItemRightRequest,
	MethodGetServiceItemRightResponse,
	MethodGetPurchaseHistoryRequest,
	MethodGetPurchaseHistoryResponse,
	MethodPostRightBinaryByAccount,
	MethodUseServiceItemByAccountRequest,
	MethodUseServiceItemByAccountResponse,
	MethodAcquireServiceItemByAccount,
	MethodGetSupportID,
	MethodGetLawMessageRequest,
	MethodGetLawMessageResponse,
}

type serviceItemProtocol = service_item.Protocol

// Protocol stores all the RMC method handlers for the Service Item (Team Kirby Clash Deluxe) protocol and listens for requests
// Embeds the Service Item protocol
type Protocol struct {
	Server *nex.Server
	serviceItemProtocol
	getEnvironmentHandler                  func(err error, packet nex.PacketInterface, callID uint32, uniqueID string, platform uint8) uint32
	httpGetRequestHandler                  func(err error, packet nex.PacketInterface, callID uint32, url *service_item_team_kirby_clash_deluxe_types.ServiceItemHTTPGetParam) uint32
	httpGetResponseHandler                 func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	purchaseServiceItemRequestHandler      func(err error, packet nex.PacketInterface, callID uint32, purchaseServiceItemParam *service_item_team_kirby_clash_deluxe_types.ServiceItemPurchaseServiceItemParam) uint32
	purchaseServiceItemResponseHandler     func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	listServiceItemRequestHandler          func(err error, packet nex.PacketInterface, callID uint32, listServiceItemParam *service_item_team_kirby_clash_deluxe_types.ServiceItemListServiceItemParam) uint32
	listServiceItemResponseHandler         func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	getBalanceRequestHandler               func(err error, packet nex.PacketInterface, callID uint32, getBalanceParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetBalanceParam) uint32
	getBalanceResponseHandler              func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	getPrepurchaseInfoRequestHandler       func(err error, packet nex.PacketInterface, callID uint32, getPrepurchaseInfoParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetPrepurchaseInfoParam) uint32
	getPrepurchaseInfoResponseHandler      func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	getServiceItemRightRequestHandler      func(err error, packet nex.PacketInterface, callID uint32, getServiceItemRightParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetServiceItemRightParam, withoutRightBinary bool) uint32
	getServiceItemRightResponseHandler     func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	getPurchaseHistoryRequestHandler       func(err error, packet nex.PacketInterface, callID uint32, getPurchaseHistoryParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetPurchaseHistoryParam) uint32
	getPurchaseHistoryResponseHandler      func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	postRightBinaryByAccountHandler        func(err error, packet nex.PacketInterface, callID uint32, postRightBinaryByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemPostRightBinaryByAccountParam) uint32
	useServiceItemByAccountRequestHandler  func(err error, packet nex.PacketInterface, callID uint32, useServiceItemByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemUseServiceItemByAccountParam) uint32
	useServiceItemByAccountResponseHandler func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
	acquireServiceItemByAccountHandler     func(err error, packet nex.PacketInterface, callID uint32, acquireServiceItemByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemAcquireServiceItemByAccountParam) uint32
	getSupportIDHandler                    func(err error, packet nex.PacketInterface, callID uint32, getSuppordIDParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetSupportIDParam) uint32
	getLawMessageRequestHandler            func(err error, packet nex.PacketInterface, callID uint32, getLawMessageParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetLawMessageParam) uint32
	getLawMessageResponseHandler           func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.serviceItemProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetEnvironment:
		go protocol.handleGetEnvironment(packet)
	case MethodHTTPGetRequest:
		go protocol.handleHTTPGetRequest(packet)
	case MethodHTTPGetResponse:
		go protocol.handleHTTPGetResponse(packet)
	case MethodPurchaseServiceItemRequest:
		go protocol.handlePurchaseServiceItemRequest(packet)
	case MethodPurchaseServiceItemResponse:
		go protocol.handlePurchaseServiceItemResponse(packet)
	case MethodListServiceItemRequest:
		go protocol.handleListServiceItemRequest(packet)
	case MethodListServiceItemResponse:
		go protocol.handleListServiceItemResponse(packet)
	case MethodGetBalanceRequest:
		go protocol.handleGetBalanceRequest(packet)
	case MethodGetBalanceResponse:
		go protocol.handleGetBalanceResponse(packet)
	case MethodGetPrepurchaseInfoRequest:
		go protocol.handleGetPrepurchaseInfoRequest(packet)
	case MethodGetPrepurchaseInfoResponse:
		go protocol.handleGetPrepurchaseInfoResponse(packet)
	case MethodGetServiceItemRightRequest:
		go protocol.handleGetServiceItemRightRequest(packet)
	case MethodGetServiceItemRightResponse:
		go protocol.handleGetServiceItemRightResponse(packet)
	case MethodGetPurchaseHistoryRequest:
		go protocol.handleGetPurchaseHistoryRequest(packet)
	case MethodGetPurchaseHistoryResponse:
		go protocol.handleGetPurchaseHistoryResponse(packet)
	case MethodPostRightBinaryByAccount:
		go protocol.handlePostRightBinaryByAccount(packet)
	case MethodUseServiceItemByAccountRequest:
		go protocol.handleUseServiceItemByAccountRequest(packet)
	case MethodUseServiceItemByAccountResponse:
		go protocol.handleUseServiceItemByAccountResponse(packet)
	case MethodAcquireServiceItemByAccount:
		go protocol.handleAcquireServiceItemByAccount(packet)
	case MethodGetSupportID:
		go protocol.handleGetSupportID(packet)
	case MethodGetLawMessageRequest:
		go protocol.handleGetLawMessageRequest(packet)
	case MethodGetLawMessageResponse:
		go protocol.handleGetLawMessageResponse(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Service Item (Team Kirby Clash Deluxe) method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new ServiceItemTeamKirbyClashDeluxe protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.serviceItemProtocol.Server = server

	protocol.Setup()

	return protocol
}
