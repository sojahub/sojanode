module.exports = async (cb) => {
    const sojahubUtilities = require('./sojahubUtilities')
    const contractUtilites = require('./contractUtilities');

    const argv = sojahubUtilities.processArgs(this, {
        ...sojahubUtilities.sharedYargOptions,
        'block_number': {
            type: "number",
            demandOption: true
        },
        'delay': {
            type: "number",
            // ropsten's average block time right now is 14 seconds, that's a fine default
            default: 14 * 1000,
            describe: "how long to wait between queries for the current block number"
        },
    });

    const logging = sojahubUtilities.configureLogging(this);

    const web3 = contractUtilites.buildWeb3(this, argv, logging);

    let waitTime = 2000;
    switch (argv.ethereum_network) {
        case "ropsten":
        case "mainnet":
            waitTime = 60000;
            break;
    }
    for (
        let blockNumber = await web3.eth.getBlockNumber();
        blockNumber < argv.block_number;
        blockNumber = await web3.eth.getBlockNumber()
    ) {
        const remaining = argv.block_number - blockNumber
        logging.debug(`wait for block ${argv.block_number}, current block ${blockNumber}, remaining blocks ${remaining}`);
        await new Promise(resolve => setTimeout(resolve, 14 * 1000));
    }
    return cb();
};
