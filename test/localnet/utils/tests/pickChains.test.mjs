import { pickChains } from "../pickChains.mjs";

test("pick chains", () => {
  const result = pickChains({ chain: "sojanode,cosmos,akash" });

  expect(result).toMatchSnapshot();
});
