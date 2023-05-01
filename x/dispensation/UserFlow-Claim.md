# Claims Flow

### Users will create claims throughout the week . 
The wallet can use 
```/dispensation/createClaim```  
This expects the following input 
```go
package main

type DistributionType int64
const Airdrop DistributionType = 1
const LiquidityMining DistributionType = 2
const ValidatorSubsidy DistributionType = 3

type CreateClaimReq struct {
	BaseReq   rest.BaseReq     `json:"base_req"`
	Signer    string           `json:"signer"`
	ClaimType DistributionType `json:"claim_type"`   
}
```

### On friday we get a list of all the claims (After Cut-off time) for the week. Any claims submitted after cutoff would be processed next week.
This query through the cli would look like
```shell
sojanoded q dispensation claims-by-type ValidatorSubsidy --chain-id sojahub --node tcp://rpc.sojahub.finance:80
```
Which returns 
```json
{
  "claims": [
    {
      "user_address": "did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z",
      "user_claim_type": "3",
      "user_claim_time": "2021-05-02T02:43:10.593125Z",
    },
    {
      "user_address": "did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm",
      "user_claim_type": "3",
      "user_claim_time": "2021-05-02T02:43:10.593125Z",
    }
  ],
  "height": "7"
}

```
The relevant event would be in the same block as the one which has the dispensation/createClaim request
```json
 {"type": "userClaim_new",
            "attributes": [
              {
                "key": "userClaim_creator",
                "value": "did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z"
              },
              {
                "key": "userClaim_type",
                "value": "ValidatorSubsidy"
              },
              {
                "key": "userClaim_creationTime",
                "value": "2021-05-02T02:43:10.593125Z"
              }
            ]
}

```

### This list obtained above is run through the parsing API ,to get output list
```json
{
 "Output": [
  {
   "address": "did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z",
   "coins": [
    {
     "denom": "fury",
     "amount": "10000000000000000000"
    }
   ]
  },
  {
   "address": "did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm",
   "coins": [
    {
     "denom": "fury",
     "amount": "10000000000000000000"
    }
   ]
  }
 ]
}
```

### This file is then used to create and run a dispensation
Create
```shell
sojanoded tx dispensation create ValidatorSubsidy output.json did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z --from did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm --yes --gas auto --gas-adjustment=1.5 --gas-prices 1.0fury
```
Run
```shell
sojanoded tx dispensation run 2_did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm ValidatorSubsidy --from did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z --yes --gas auto --gas-adjustment=1.5 --gas-prices 1.0fury
```
