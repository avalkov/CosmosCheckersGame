/* eslint-disable */
import { StoredGame } from '../cosmoscheckersgame/stored_game';
import { NextGame } from '../cosmoscheckersgame/next_game';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'avalkov.cosmoscheckersgame.cosmoscheckersgame';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.storedGameList) {
            StoredGame.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.nextGame !== undefined) {
            NextGame.encode(message.nextGame, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.storedGameList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 2:
                    message.storedGameList.push(StoredGame.decode(reader, reader.uint32()));
                    break;
                case 1:
                    message.nextGame = NextGame.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.storedGameList = [];
        if (object.storedGameList !== undefined && object.storedGameList !== null) {
            for (const e of object.storedGameList) {
                message.storedGameList.push(StoredGame.fromJSON(e));
            }
        }
        if (object.nextGame !== undefined && object.nextGame !== null) {
            message.nextGame = NextGame.fromJSON(object.nextGame);
        }
        else {
            message.nextGame = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.storedGameList) {
            obj.storedGameList = message.storedGameList.map((e) => (e ? StoredGame.toJSON(e) : undefined));
        }
        else {
            obj.storedGameList = [];
        }
        message.nextGame !== undefined && (obj.nextGame = message.nextGame ? NextGame.toJSON(message.nextGame) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.storedGameList = [];
        if (object.storedGameList !== undefined && object.storedGameList !== null) {
            for (const e of object.storedGameList) {
                message.storedGameList.push(StoredGame.fromPartial(e));
            }
        }
        if (object.nextGame !== undefined && object.nextGame !== null) {
            message.nextGame = NextGame.fromPartial(object.nextGame);
        }
        else {
            message.nextGame = undefined;
        }
        return message;
    }
};
