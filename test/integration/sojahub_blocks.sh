height=$(sojanoded --home $CHAINDIR/.sojanoded q block | jq -r .block.header.height)
seq $height | parallel -k sojanoded --home $CHAINDIR/.sojanoded q block {}
