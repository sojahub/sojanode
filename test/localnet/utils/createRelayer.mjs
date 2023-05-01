import { $, nothrow } from "zx";
import { send } from "../lib/send.mjs";

export async function createRelayer({
  sojaChainProps,
  otherChainProps,
  registryFrom = `/tmp/localnet/config/registry`,
}) {
  const { chain, home } = otherChainProps;
  const relayerHome = `${home}/relayer`;

  await nothrow($`mkdir -p ${relayerHome}`);
  await nothrow(
    $`ibc-setup init --home ${relayerHome} --registry-from ${registryFrom} --src ${sojaChainProps.chain} --dest ${chain}`
  );

  let addresses = await $`ibc-setup keys list --home ${relayerHome}`;
  addresses = addresses.toString().split("\n");

  const sojaChainAddress = addresses
    .find((item) => item.includes(`${sojaChainProps.chain}`))
    .replace(`${sojaChainProps.chain}: `, ``);
  const otherChainAddress = addresses
    .find((item) => item.includes(`${chain}`))
    .replace(`${chain}: `, ``);

  console.log(`sojaChainAddress: ${sojaChainAddress}`);
  console.log(`otherChainAddress: ${otherChainAddress}`);

  await send({
    ...otherChainProps,
    src: `${chain}-source`,
    dst: otherChainAddress,
    amount: 10e10,
    node: `tcp://127.0.0.1:${otherChainProps.rpcPort}`,
  });

  return {
    sojaChainAddress,
    otherChainAddress,
    sojaSendRequest: {
      ...sojaChainProps,
      src: `${sojaChainProps.chain}-source`,
      dst: sojaChainAddress,
      amount: 10e10,
      node: `tcp://127.0.0.1:${sojaChainProps.rpcPort}`,
    },
  };
}
