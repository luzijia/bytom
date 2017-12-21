package account

import (
	"encoding/json"

	"github.com/bytom/blockchain/query"
	"github.com/bytom/blockchain/signers"
	chainjson "github.com/bytom/encoding/json"
	"github.com/bytom/protocol/bc"
)

const (
	//UTXOPreFix is AccountUTXOKey prefix
	UTXOPreFix = "ACU:"
)

//UTXOKey makes a account unspent outputs key to store
func UTXOKey(id bc.Hash) []byte {
	name := string(id.Bytes())
	return []byte(UTXOPreFix + name)
}

//UTXO is a structure about account unspent outputs
type UTXO struct {
	OutputID     []byte
	AssetID      []byte
	Amount       uint64
	AccountID    string
	ProgramIndex uint64
	Program      []byte
	SourceID     []byte
	SourcePos    uint64
	RefData      []byte
	Change       bool
}

var emptyJSONObject = json.RawMessage(`{}`)

//Annotated init an annotated account object
func Annotated(a *Account) (*query.AnnotatedAccount, error) {
	aa := &query.AnnotatedAccount{
		ID:     a.ID,
		Alias:  a.Alias,
		Quorum: a.Quorum,
		Tags:   &emptyJSONObject,
	}

	tags, err := json.Marshal(a.Tags)
	if err != nil {
		return nil, err
	}
	if len(tags) > 0 {
		rawTags := json.RawMessage(tags)
		aa.Tags = &rawTags
	}

	path := signers.Path(a.Signer, signers.AccountKeySpace)
	var jsonPath []chainjson.HexBytes
	for _, p := range path {
		jsonPath = append(jsonPath, p)
	}
	for _, xpub := range a.XPubs {
		aa.Keys = append(aa.Keys, &query.AccountKey{
			RootXPub:              xpub,
			AccountXPub:           xpub.Derive(path),
			AccountDerivationPath: jsonPath,
		})
	}
	return aa, nil
}
