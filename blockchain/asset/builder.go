package asset

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/bytom/blockchain/signers"
	"github.com/bytom/blockchain/txbuilder"
	chainjson "github.com/bytom/encoding/json"
	"github.com/bytom/protocol/bc"
	"github.com/bytom/protocol/bc/legacy"
)

//NewIssueAction create a new asset issue action
func (reg *Registry) NewIssueAction(assetAmount bc.AssetAmount, referenceData chainjson.Map) txbuilder.Action {
	return &issueAction{
		assets:        reg,
		AssetAmount:   assetAmount,
		ReferenceData: referenceData,
	}
}

//DecodeIssueAction unmarshal JSON-encoded data of asset issue action
func (reg *Registry) DecodeIssueAction(data []byte) (txbuilder.Action, error) {
	a := &issueAction{assets: reg}
	err := json.Unmarshal(data, a)
	return a, err
}

type issueAction struct {
	assets *Registry
	bc.AssetAmount
	ReferenceData chainjson.Map `json:"reference_data"`
}

func (a *issueAction) Build(ctx context.Context, builder *txbuilder.TemplateBuilder) error {
	if a.AssetId.IsZero() {
		return txbuilder.MissingFieldsError("asset_id")
	}

	asset, err := a.assets.findByID(ctx, *a.AssetId)
	if err != nil {
		return err
	}

	var nonce [8]byte
	_, err = rand.Read(nonce[:])
	if err != nil {
		return err
	}

	assetdef := asset.RawDefinition()

	txin := legacy.NewIssuanceInput(nonce[:], a.Amount, a.ReferenceData, asset.InitialBlockHash, asset.IssuanceProgram, nil, assetdef)

	tplIn := &txbuilder.SigningInstruction{}
	path := signers.Path(asset.Signer, signers.AssetKeySpace)
	tplIn.AddWitnessKeys(asset.Signer.XPubs, path, asset.Signer.Quorum)

	log.WithFields(log.Fields{"txin": txin, "tplIn": tplIn}).Info("Issue action build")
	builder.RestrictMinTime(time.Now())
	return builder.AddInput(txin, tplIn)
}
