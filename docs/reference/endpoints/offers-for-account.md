---
title: Offers for Account
clientData:
  laboratoryUrl: http://zionc.info/laboratory/#explorer?resource=offers&endpoint=for_account
---

People on the Zion network can make [offers](../resources/offer.md) to buy or sell assets.  This endpoint represents all the offers a particular account makes.


## Request

```
GET /accounts/{account}/offers{?cursor,limit,order}
```

### Arguments

| name | notes | description | example |
| ---- | ----- | ----------- | ------- |
| `account` | required, string | Account ID | `GA2HGBJIJKI6O4XEM7CZWY5PS6GKSXL6D34ERAJYQSPYA6X6AI7HYW36` |
| `?cursor` | optional, any, default _null_ | A paging token, specifying where to start returning records from. | `12884905984` |
| `?order`  | optional, string, default `asc` | The order in which to return rows, "asc" or "desc". | `asc` |
| `?limit`  | optional, number, default: `10` | Maximum number of records to return. | `200` |

### curl Example Request

```sh
curl "https://equator-testnet.zion.org/accounts/GCJ34JYMXNI7N55YREWAACMMZECOMTPIYDTFCQBWPUP7BLJQDDTVGUW4/offers"
```

### JavaScript Example Request

```js
var ZionSdk = require('zion-sdk');
var server = new ZionSdk.Server('https://equator-testnet.zion.org');

server.offers('accounts', 'GCJ34JYMXNI7N55YREWAACMMZECOMTPIYDTFCQBWPUP7BLJQDDTVGUW4')
  .call()
  .then(function (offerResult) {
    console.log(offerResult);
  })
  .catch(function (err) {
    console.error(err);
  })
```

## Response

The list of offers.

### Example Response

```js
{
  "_links": {
    "self": {
      "href": "https://equator-testnet.zion.org/accounts/GCJ34JYMXNI7N55YREWAACMMZECOMTPIYDTFCQBWPUP7BLJQDDTVGUW4/offers?order=asc&limit=10&cursor="
    },
    "next": {
      "href": "https://equator-testnet.zion.org/accounts/GCJ34JYMXNI7N55YREWAACMMZECOMTPIYDTFCQBWPUP7BLJQDDTVGUW4/offers?order=asc&limit=10&cursor=122"
    },
    "prev": {
      "href": "https://equator-testnet.zion.org/accounts/GCJ34JYMXNI7N55YREWAACMMZECOMTPIYDTFCQBWPUP7BLJQDDTVGUW4/offers?order=desc&limit=10&cursor=121"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "self": {
            "href": "https://equator-testnet.zion.org/offers/121"
          },
          "offer_maker": {
            "href": "https://equator-testnet.zion.org/accounts/GCJ34JYMXNI7N55YREWAACMMZECOMTPIYDTFCQBWPUP7BLJQDDTVGUW4"
          }
        },
        "id": 121,
        "paging_token": "121",
        "seller": "GCJ34JYMXNI7N55YREWAACMMZECOMTPIYDTFCQBWPUP7BLJQDDTVGUW4",
        "selling": {
          "asset_type": "credit_alphanum4",
          "asset_code": "BAR",
          "asset_issuer": "GBAUUA74H4XOQYRSOW2RZUA4QL5PB37U3JS5NE3RTB2ELJVMIF5RLMAG"
        },
        "buying": {
          "asset_type": "credit_alphanum4",
          "asset_code": "FOO",
          "asset_issuer": "GBAUUA74H4XOQYRSOW2RZUA4QL5PB37U3JS5NE3RTB2ELJVMIF5RLMAG"
        },
        "amount": "23.6692509",
        "price_r": {
          "n": 387,
          "d": 50
        },
        "price": "7.7400000"
      },
      {
        "_links": {
          "self": {
            "href": "https://equator-testnet.zion.org/offers/122"
          },
          "offer_maker": {
            "href": "https://equator-testnet.zion.org/accounts/GCJ34JYMXNI7N55YREWAACMMZECOMTPIYDTFCQBWPUP7BLJQDDTVGUW4"
          }
        },
        "id": 122,
        "paging_token": "122",
        "seller": "GCJ34JYMXNI7N55YREWAACMMZECOMTPIYDTFCQBWPUP7BLJQDDTVGUW4",
        "selling": {
          "asset_type": "credit_alphanum4",
          "asset_code": "BAR",
          "asset_issuer": "GBAUUA74H4XOQYRSOW2RZUA4QL5PB37U3JS5NE3RTB2ELJVMIF5RLMAG"
        },
        "buying": {
          "asset_type": "credit_alphanum4",
          "asset_code": "FOO",
          "asset_issuer": "GBAUUA74H4XOQYRSOW2RZUA4QL5PB37U3JS5NE3RTB2ELJVMIF5RLMAG"
        },
        "amount": "72.0000000",
        "price_r": {
          "n": 779,
          "d": 100
        },
        "price": "7.7900000"
      }
    ]
  }
}
```

## Possible Errors

- The [standard errors](../errors.md#Standard_Errors).
