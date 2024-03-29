// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
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

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgRejectGame: (data: MsgRejectGame): EncodeObject => ({ typeUrl: "/avalkov.cosmoscheckersgame.cosmoscheckersgame.MsgRejectGame", value: data }),
    msgPlayMove: (data: MsgPlayMove): EncodeObject => ({ typeUrl: "/avalkov.cosmoscheckersgame.cosmoscheckersgame.MsgPlayMove", value: data }),
    msgCreateGame: (data: MsgCreateGame): EncodeObject => ({ typeUrl: "/avalkov.cosmoscheckersgame.cosmoscheckersgame.MsgCreateGame", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
