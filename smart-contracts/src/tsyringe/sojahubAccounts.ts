import {HardhatRuntimeEnvironment} from "hardhat/types";
import {HardhatRuntimeEnvironmentToken,} from "./injectionTokens";
import {SignerWithAddress} from "@nomiclabs/hardhat-ethers/signers";
import {inject, injectable} from "tsyringe";
import {isHardhatRuntimeEnvironment} from "./hardhatSupport";

/**
 * The accounts necessary for testing a sojahub system
 */
export class SojahubAccounts {
    constructor(
        readonly operatorAccount: SignerWithAddress,
        readonly ownerAccount: SignerWithAddress,
        readonly pauserAccount: SignerWithAddress,
        readonly validatatorAccounts: Array<SignerWithAddress>,
        readonly availableAccounts: Array<SignerWithAddress>
    ) {
    }
}

/**
 * Note that the hardhat environment provides accounts as promises, so
 * we need to wrap a SojahubAccounts in a promise.
 */
@injectable()
export class SojahubAccountsPromise {
    accounts: Promise<SojahubAccounts>

    constructor(accounts: Promise<SojahubAccounts>);
    constructor(@inject(HardhatRuntimeEnvironmentToken) hardhatOrAccounts: HardhatRuntimeEnvironment | Promise<SojahubAccounts>) {
        if (isHardhatRuntimeEnvironment(hardhatOrAccounts)) {
            this.accounts = hreToSojahubAccountsAsync(hardhatOrAccounts)
        } else {
            this.accounts = hardhatOrAccounts
        }
    }
}

export async function hreToSojahubAccountsAsync(hardhat: HardhatRuntimeEnvironment): Promise<SojahubAccounts> {
    const accounts = await hardhat.ethers.getSigners()
    const [operatorAccount, ownerAccount, pauserAccount, validator1Account, ...extraAccounts] = accounts
    return new SojahubAccounts(operatorAccount, ownerAccount, pauserAccount, [validator1Account], extraAccounts)
}
