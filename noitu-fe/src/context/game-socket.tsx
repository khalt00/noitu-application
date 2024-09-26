"use client";
import { ResponseMessage } from "@/type/ws";
import React from "react";
import {
  useContext,
  Dispatch,
  ReactNode,
  SetStateAction,
  createContext,
  useMemo,
  useState,
} from "react";

interface GameWebsocketContextType {
  ws?: WebSocket;
  setWs: Dispatch<SetStateAction<WebSocket | undefined>>;
}

const GameWebsocketContext = createContext<GameWebsocketContextType>({
  ws: undefined,
  setWs: () => {},
});

export const GameWebsocketProvider = ({
  children,
}: {
  children: ReactNode;
}) => {
  const [ws, setWs] = useState<WebSocket>();

  const value = useMemo(
    () => ({
      ws,
      setWs,
    }),
    [ws, setWs]
  );

  return (
    <GameWebsocketContext.Provider value={value}>
      {children}
    </GameWebsocketContext.Provider>
  );
};

export const useGameWebsocket = () => {
  const { ws, setWs } = useContext(GameWebsocketContext);
  const isConnected = ws !== undefined;

  const testJoin = (username: string, category: string) => {
    const url = `ws://localhost:8081/ws?username=${username}&category=${category}`;
    const ws = new WebSocket(url);

    ws.addEventListener("close", () => {
      setWs(undefined);
    });

    setWs(ws);
  };

  const sendMessage = (word: string) => {
    if (ws) {
      const message = {
        word,
      };
      ws.send(JSON.stringify(message));
    }
  };

  const addMessageListener = (
    onMessage: (message: ResponseMessage) => void
  ) => {
    if (ws) {
      console.log("go here?")
      ws.addEventListener("message", (ev) => {
        try {
          onMessage(JSON.parse(ev.data));
        } catch (e) {}
      });
    }
  };

  const addCloseListener = (onClose: () => void) => {
    if (ws) {
      ws.addEventListener("close", () => {
        onClose();
      });
    }
  };

  const closeConnection = () => {
    if (ws) {
      ws.close();
    }
  };

  return {
    isConnected,
    testJoin,
    sendMessage,
    //   createRoom,
    //   joinRoom,
    //   startGame,
    //   skipQuestion,
    //   nextQuestion,
    //   showScoreboard,
    //   playAgain,
    addMessageListener,
    addCloseListener,
    closeConnection,
  };
};
