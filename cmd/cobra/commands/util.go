package commands

import (
	"github.com/bytom/crypto/ed25519/chainkd"
	"github.com/bytom/encoding/json"
)

// accountIns is used for account related request.
type accountIns struct {
	RootXPubs   []chainkd.XPub         `json:"root_xpubs"`
	Quorum      int                    `json:"quorum"`
	Alias       string                 `json:"alias"`
	Tags        map[string]interface{} `json:"tags"`
	AccessToken string                 `json:"access_token"`
}

// assetIns is used for asset related request.
type assetIns struct {
	RootXPubs   []chainkd.XPub         `json:"root_xpubs"`
	Quorum      int                    `json:"quorum"`
	Alias       string                 `json:"alias"`
	Tags        map[string]interface{} `json:"tags"`
	Definition  map[string]interface{} `json:"definition"`
	AccessToken string                 `json:"access_token"`
}

type requestQuery struct {
	Filter       string        `json:"filter,omitempty"`
	FilterParams []interface{} `json:"filter_params,omitempty"`
	SumBy        []string      `json:"sum_by,omitempty"`
	PageSize     int           `json:"page_size"`
	AscLongPoll  bool          `json:"ascending_with_long_poll,omitempty"`
	Timeout      json.Duration `json:"timeout"`
	After        string        `json:"after"`
	StartTimeMS  uint64        `json:"start_time,omitempty"`
	EndTimeMS    uint64        `json:"end_time,omitempty"`
	TimestampMS  uint64        `json:"timestamp,omitempty"`
	Type         string        `json:"type"`
	Aliases      []string      `json:"aliases,omitempty"`
}