1. Initialize and start the chain. From the repo root run the following:

```
make init
make run
```

2. Create pool - add one million dollars of usdt and price fury at 10c:

```
sojanoded tx clp create-pool \
  --from soja \
  --keyring-backend test \
  --symbol cusdt \
  --nativeAmount 10000000000000000000000000 \
  --externalAmount 1000000000000 \
  --fees 100000000000000000fury \
  --chain-id localnet \
  --broadcast-mode block \
  -y
```

3. Query the lppd parameters:

```
sojanoded q clp lppd-params
```

4. Set a new lppd policy:

```
sojanoded tx clp set-lppd-params \
   --from soja \
   --keyring-backend test \
   --chain-id localnet \
   --broadcast-mode block \
   -y \
   --path <( echo '[
    {
        "distribution_period_block_rate": "0.01",
        "distribution_period_start_block": 1,
        "distribution_period_mod": 1,
        "distribution_period_end_block": 433000
    }
]' )
```

5. Query block results and observe the lppd success event:

```
curl http://localhost:26657/block_results
```
