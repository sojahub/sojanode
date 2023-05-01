# AuthZ tutorial 
- Authz module allows granting arbitrary privileges from one account (the granter) to another account (the grantee). Authorizations must be granted for a particular Msg service method one by one using an implementation of the Authorization interface.
- The built in types include:`send`,`generic`,`delegate`,`unbond`,`redelegate`
- The `generic` authorization can be used to authorize any address to execute a message on their behalf

## Steps to provide authorization
1. Grant authorization to a particular address
```shell
sojanoded tx authz grant soja1l7hypmqk2yc334vc6vmdwzp5sdefygj2ad93p5 generic --msg-type=/sojanode.clp.v1.MsgCreatePool --from=soja --keyring-backend=test --chain-id=localnet

```
In this case the granter is `soja` . This allows `soja1l7hypmqk2yc334vc6vmdwzp5sdefygj2ad93p5` to perform any TX of type `MsgCreatePool` on their behalf
Query grants
```shell
sojanoded q authz grants soja1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd soja1l7hypmqk2yc334vc6vmdwzp5sdefygj2ad93p5 /sojanode.clp.v1.MsgCreatePool
```
2. Create tx
```shell
sojanoded tx clp create-pool --from soja1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd --symbol ceth --nativeAmount 1000000000000000000 --externalAmount 1000000000000000000  --yes --chain-id=localnet --keyring-backend=test --generate-only > tx.json
```

3. Sign and broadcast
```shell
 sojanoded tx authz exec tx.json --from akasha --keyring-backend=test --chain-id=localnet
```
Logs from exec 
```json lines
    messages:
    - '@type': /cosmos.authz.v1beta1.MsgExec
      grantee: soja1l7hypmqk2yc334vc6vmdwzp5sdefygj2ad93p5
      msgs:
      - '@type': /sojanode.clp.v1.MsgCreatePool
        external_asset:
          symbol: ceth
        external_asset_amount: "1000000000000000000"
        native_asset_amount: "1000000000000000000"
        signer: soja1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd
```
Notes 
- The MsgCreatePool is wrapped inside a MsgExec .
- The signer for MsgCreatePool is `soja` , but the actual signature was done by `akasha`