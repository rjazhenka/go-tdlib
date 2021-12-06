// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
)

// CallServer Describes a server for relaying call data
type CallServer struct {
	tdCommon
	ID          JSONInt64      `json:"id"`           // Server identifier
	IPAddress   string         `json:"ip_address"`   // Server IPv4 address
	IPv6Address string         `json:"ipv6_address"` // Server IPv6 address
	Port        int32          `json:"port"`         // Server port number
	Type        CallServerType `json:"type"`         // Server type
}

// MessageType return the string telegram-type of CallServer
func (callServer *CallServer) MessageType() string {
	return "callServer"
}

// NewCallServer creates a new CallServer
//
// @param iD Server identifier
// @param iPAddress Server IPv4 address
// @param iPv6Address Server IPv6 address
// @param port Server port number
// @param typeParam Server type
func NewCallServer(iD JSONInt64, iPAddress string, iPv6Address string, port int32, typeParam CallServerType) *CallServer {
	callServerTemp := CallServer{
		tdCommon:    tdCommon{Type: "callServer"},
		ID:          iD,
		IPAddress:   iPAddress,
		IPv6Address: iPv6Address,
		Port:        port,
		Type:        typeParam,
	}

	return &callServerTemp
}

// UnmarshalJSON unmarshal to json
func (callServer *CallServer) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID          JSONInt64 `json:"id"`           // Server identifier
		IPAddress   string    `json:"ip_address"`   // Server IPv4 address
		IPv6Address string    `json:"ipv6_address"` // Server IPv6 address
		Port        int32     `json:"port"`         // Server port number

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	callServer.tdCommon = tempObj.tdCommon
	callServer.ID = tempObj.ID
	callServer.IPAddress = tempObj.IPAddress
	callServer.IPv6Address = tempObj.IPv6Address
	callServer.Port = tempObj.Port

	fieldType, _ := unmarshalCallServerType(objMap["type"])
	callServer.Type = fieldType

	return nil
}
