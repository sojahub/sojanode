For reference when playing with cosmos cain

# Create liquidity pool catk:fury

```
sojanoded tx clp create-pool \
 --from akasha \
 --symbol catk \
 --nativeAmount 500 \
 --externalAmount 500
```

# Create liquidity pool cbtk:fury

```
sojanoded tx clp create-pool \
 --from akasha \
 --symbol cbtk \
 --nativeAmount 500 \
 --externalAmount 500
```

# Verify pool created

```
sojanoded query clp pools
```

# Execute swap

```
sojanoded tx clp swap \
 --from shadowfiend \
 --sentSymbol catk \
 --receivedSymbol cbtk \
 --sentAmount 20
```
