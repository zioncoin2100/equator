package resource

import (
	"github.com/zion/go/amount"
	"github.com/zion/equator/assets"
	"github.com/zion/equator/db2/core"
	"github.com/zion/equator/httpx"
	"github.com/zion/equator/render/hal"
	"golang.org/x/net/context"
)

func (this *Offer) Populate(ctx context.Context, row core.Offer) {
	this.ID = row.OfferID
	this.PT = row.PagingToken()
	this.Seller = row.SellerID
	this.Amount = amount.String(row.Amount)
	this.PriceR.N = row.Pricen
	this.PriceR.D = row.Priced
	this.Price = row.PriceAsString()
	this.Buying = Asset{
		Type:   assets.MustString(row.BuyingAssetType),
		Code:   row.BuyingAssetCode.String,
		Issuer: row.BuyingIssuer.String,
	}
	this.Selling = Asset{
		Type:   assets.MustString(row.SellingAssetType),
		Code:   row.SellingAssetCode.String,
		Issuer: row.SellingIssuer.String,
	}

	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	this.Links.Self = lb.Linkf("/offers/%d", row.OfferID)
	this.Links.OfferMaker = lb.Linkf("/accounts/%s", row.SellerID)
	return
}

func (this Offer) PagingToken() string {
	return this.PT
}
