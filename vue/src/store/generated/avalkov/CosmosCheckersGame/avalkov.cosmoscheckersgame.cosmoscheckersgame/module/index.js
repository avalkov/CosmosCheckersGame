// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRejectGame } from "./types/cosmoscheckersgame/tx";
import { MsgPlayMove } from "./types/cosmoscheckersgame/tx";
import { MsgCreateGame } from "./types/cosmoscheckersgame/tx";
const types = [
    ["/avalkov.cosmoscheckersgame.cosmoscheckersgame.MsgRejectGame", MsgRejectGame],
    ["/avalkov.cosmoscheckersgame.cosmoscheckersgame.MsgPlayMove", MsgPlayMove],
    ["/avalkov.cosmoscheckersgame.cosmoscheckersgame.MsgCreateGame", MsgCreateGame],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgRejectGame: (data) => ({ typeUrl: "/avalkov.cosmoscheckersgame.cosmoscheckersgame.MsgRejectGame", value: data }),
        msgPlayMove: (data) => ({ typeUrl: "/avalkov.cosmoscheckersgame.cosmoscheckersgame.MsgPlayMove", value: data }),
        msgCreateGame: (data) => ({ typeUrl: "/avalkov.cosmoscheckersgame.cosmoscheckersgame.MsgCreateGame", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
