


# Multisig Key - It is a key composed of two or more keys (N) , with a signing threshold (K) ,such that the transaction needs K out of N votes to go through.
sojanoded tx dispensation claim ValidatorSubsidy --from akasha --keyring-backend test --yes --chain-id localnet -o json
sojanoded tx dispensation claim ValidatorSubsidy --from soja --keyring-backend test --yes --chain-id localnet -o json
# create airdrop
# mkey = multisig key
# ar1 = name for airdrop , needs to be unique for every airdrop . If not the tx gets rejected
# input.json list of funding addresses  -  Input address must be part of the multisig key
# output.json list of airdrop receivers.
sleep 8
sojanoded q dispensation claims-by-type ValidatorSubsidy -o json
sleep 8
sojanoded tx dispensation create ValidatorSubsidy output.json did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm --gas 200064128 --from did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm --keyring-backend test --fees 100000fury --yes --chain-id sojahub-devnet-042 --node tcp://rpc-devnet-042.sojahub.finance:80

sleep 8
sojanoded q dispensation distributions-all --chain-id localnet -o json
#sojanoded q dispensation records-by-name-all ar1 >> all.json
#sojanoded q dispensation records-by-name-pending ar1 >> pending.json
#sojanoded q dispensation records-by-name-completed ar1 >> completed.json
#sojanoded q dispensation records-by-addr did:fury:s1cp23ye3h49nl5ty35vewrtvsgwnuczt03jwg00


