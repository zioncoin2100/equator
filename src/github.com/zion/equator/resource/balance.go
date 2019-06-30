package resource

import (
	"github.com/zion/go/amount"
	"github.com/zion/go/xdr"
	"github.com/zion/equator/assets"
	"github.com/zion/equator/db2/core"
	"golang.org/x/net/context"
)

func (this *Balance) Populate(ctx context.Context, row core.Trustline) (err error) {
	this.Type, err = assets.String(row.Assettype)
	if err != nil {
		return
	}

	this.Balance = amount.String(row.Balance)
	this.Limit = amount.String(row.Tlimit)
	this.Issuer = row.Issuer
	this.Code = row.Assetcode
	return
}

func (this *Balance) PopulateNative(stroops xdr.Int64) (err error) {
	this.Type, err = assets.String(xdr.AssetTypeAssetTypeNative)
	if err != nil {
		return
	}

	this.Balance = amount.String(stroops)
	this.Limit = ""
	this.Issuer = ""
	this.Code = ""
	return
}
