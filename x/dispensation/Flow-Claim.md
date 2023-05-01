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
(This step can be done through CLI on friday , or process events throughout the week . Processing events would be the preferred approach)
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
      "locked": false
    },
    {
      "user_address": "did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm",
      "user_claim_type": "3",
      "user_claim_time": "2021-05-02T02:43:10.593125Z",
      "locked": false
    }
  ],
  "height": "7"
}

```
We can also parse events instead of the using this query . This event would be in the same block as the one which has the dispensation/createClaim request
```json
 {"type": "claim_created",
            "attributes": [
              {
                "key": "Y2xhaW1fY3JlYXRvcg==",
                "value": "c2lmMWw3aHlwbXFrMnljMzM0dmM2dm1kd3pwNXNkZWZ5Z2oyYWQ5M3A1"
              },
              {
                "key": "Y2xhaW1fdHlwZQ==",
                "value": "VmFsaWRhdG9yU3Vic2lkeQ=="
              },
              {
                "key": "dXNlckNsYWltX2NyZWF0aW9uVGltZQ==",
                "value": "MjAyMS0wNS0wMlQwMjo0MzoxMC41OTMxMjVa"
              }
            ]
}
```
After parsing should become 
```json
 {"type": "claim_created",
            "attributes": [
              {
                "key": "claim_creator",
                "value": "did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z"
              },
              {
                "key": "claim_type",
                "value": "ValidatorSubsidy"
              },
              {
                "key": "userClaim_creationTime",
                "value": "2021-05-02T02:43:10.593125Z"
              }
            ]
}

```

### This list obtained above is run through the parsing API(Niko) , which should create a file like below  
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

### This file is then used to create a distribution
```shell
sojanoded tx dispensation create mkey ar1 ValidatorSubsidy input.json output.json --gas 200064128 --generate-only >> offlinetx.json
# First user signs
sojanoded tx sign --multisig $(sojanoded keys show mkey -a) --from $(sojanoded keys show soja -a)  offlinetx.json >> sig1.json
# Second user signs
sojanoded tx sign --multisig $(sojanoded keys show mkey -a) --from $(sojanoded keys show akasha -a)  offlinetx.json >> sig2.json
# Multisign created from the above signatures
sojanoded tx multisign offlinetx.json mkey sig1.json sig2.json >> signedtx.json
# transaction broadcast , distribution happens
sojanoded tx broadcast signedtx.json
```
### Post Dispensation
- Suppose we do the dispensation at height X , The block results should contain this event 
```json
  {
            "type": "distribution_started",
            "attributes": [
              {
                "key": "module_account",
                "value": "did:fury:s1zvwfuvy3nh949rn68haw78rg8jxjevgm2c820c"
              }
            ]
          }
```
- This event signifies that the distribution started , and transfers should begin at height X+1 . The transfer events can be used as confirmation to reset account multipliers.
The sender address for the transfers should be the module account in the distribution_started event . 
```json
 [
          {
            "key": "recipient",
            "value": "did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z"
          },
          {
            "key": "sender",
            "value": "did:fury:s1zvwfuvy3nh949rn68haw78rg8jxjevgm2c820c"
          },
          {
            "key": "amount",
            "value": "10000000000000000000fury"
          }
        ]

```
