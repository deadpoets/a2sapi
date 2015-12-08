package models

// api_serverlist.go - Model for building list of server details

import "time"

// APIServerList represents the server detail list returned in response to
// building the master list or in response to building the list of server details
// via a user's API request.
type APIServerList struct {
	RetrievedAt        string       `json:"retrievalDate"`
	RetrievedTimeStamp int64        `json:"timestamp"`
	ServerCount        int          `json:"serverCount"`
	Servers            []*APIServer `json:"servers"`
	FailedCount        int          `json:"failedCount"`
	FailedServers      []string     `json:"failedServers"`
}

// APIServer represents an individual game server's information, including its
// A2S information as well as its geographical data. if available.
type APIServer struct {
	ID          int64      `json:"serverId"`
	Host        string     `json:"address"`
	Game        string     `json:"-"`
	IP          string     `json:"ip"`
	Port        int        `json:"port"`
	CountryInfo *DbCountry `json:"location"`
	// 'Info' by default was *ServerInfo, but nil pointers are encoded as
	// 'null' in JSON instead of an empty object, so use interface and handle appropriately
	Info        interface{}        `json:"info"`
	Players     []*SteamPlayerInfo `json:"players"`
	RealPlayers *RealPlayerInfo    `json:"realPlayers"`
	Rules       map[string]string  `json:"rules"`
}

// GetDefaultServerList Returns a default, empty, server list with the current
// date and time in response to a server detail list request that failed for
// whatever reason.
func GetDefaultServerList() *APIServerList {
	return &APIServerList{
		RetrievedAt:        time.Now().Format("Mon Jan _2 15:04:05 2006 EST"),
		RetrievedTimeStamp: time.Now().Unix(),
		ServerCount:        0,
		Servers:            make([]*APIServer, 0),
		FailedCount:        0,
		FailedServers:      make([]string, 0),
	}
}