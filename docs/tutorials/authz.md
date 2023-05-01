# AuthZ tutorial 
- Authz module allows granting arbitrary privileges from one account (the granter) to another account (the grantee). Authorizations must be granted for a particular Msg service method one by one using an implementation of the Authorization interface.
- The built in types include:`send`,`generic`,`delegate`,`unbond`,`redelegate`
- The `generic` authorization can be used to authorize any address to execute a message on their behalf

## Steps to provide authorization
1. Grant authorization to a particular address
```shell
sojanoded tx authz grant did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z generic --msg-type=/sojanode.clp.v1.MsgCreatePool --from=soja --keyring-backend=test --chain-id=localnet

```
In this case the granter is `soja` . This allows `did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z` to perform any TX of type `MsgCreatePool` on their behalf
Query grants
```shell
sojanoded q authz grants did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z /sojanode.clp.v1.MsgCreatePool
```
2. Create tx
```shell
sojanoded tx clp create-pool --from did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm --symbol ceth --nativeAmount 1000000000000000000 --externalAmount 1000000000000000000  --yes --chain-id=localnet --keyring-backend=test --generate-only > tx.json
```

3. Sign and broadcast
```shell
 sojanoded tx authz exec tx.json --from akasha --keyring-backend=test --chain-id=localnet
```
Logs from exec 
```json lines
    messages:
    - '@type': /cosmos.authz.v1beta1.MsgExec
      grantee: did:fury:s1l7hypmqk2yc334vc6vmdwzp5sdefygj2qt269z
      msgs:
      - '@type': /sojanode.clp.v1.MsgCreatePool
        external_asset:
          symbol: ceth
        external_asset_amount: "1000000000000000000"
        native_asset_amount: "1000000000000000000"
        signer: did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm
```
Notes 
- The MsgCreatePool is wrapped inside a MsgExec .
- The signer for MsgCreatePool is `soja` , but the actual signature was done by `akasha`