# Distributes funds to these test wallets from FURY_SOURCE

cp $BASEDIR/smart-contracts/deployments/$DEPLOYMENT_NAME/BridgeBank.json $BASEDIR/smart-contracts/build/contracts

test_wallets="did:fury:s1fpq67nw66thzmf2a5ng64cd8p8nxa5vl9d3cm4
did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777exhphm
did:fury:s1hjkgsq0wcmwdh8pr3snhswx5xyy4zpgs833akh
did:fury:s1ypc5qcq5ha562xlak4xw3g6v352k39t6868jhx
did:fury:s1u7cp5e5kty8xwuu7k234ah4jsknvkzazqagvl6
did:fury:s1lj3rsayj4xtrhp2e3elv4nf7lazxty272zqegr
did:fury:s1cffgyxgvw80rr6n9pcwpzrm6v8cd6dax8x32f5
did:fury:s1dlse3w2pxlmuvsj5eda344zp99fegual958qyr
did:fury:s1m7257566ehx7ya4ypeq7lj4y2h075z6u2xu79v
did:fury:s1qrxylp97p25wcqn4cs9nd02v672073ynpkt4yr
did:fury:s13rysrrdlhtmuc2pzve7jk0t4pwytwyxhaqcqcn
did:fury:s1shywxv2g8gvjcqknvkxu4p6lkqhfclwwj2qk6h
did:fury:s1gqm44p5ax4kgk6hksxgv4vuh2adue2acxvg542
did:fury:s1zwgc9frcfpt3hhkqfu9u7up94ag5rp30kwrwrj"

for i in $test_wallets
do
  DESTINATION_ACCOUNT=$i python3 -m pytest --color=yes -x -olog_cli=true -olog_level=DEBUG -v -olog_file=vagrant/data/pytest.log src/py/token_distribution.py
done
