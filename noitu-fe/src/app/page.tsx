"use client";

import { Button } from "@/components/ui/button";
import React, { useEffect, useState } from "react";
import { Input } from "@/components/ui/input";
import { useGameWebsocket } from "@/context/game-socket";
import { ResponseMessage } from "@/type/ws";

enum GameState {
  NONE,
  QUEUEING,
  PLAYING,
  ENDING,
}

export default function Home() {
  const [username, setUsername] = useState("");
  const [word, setWord] = useState("");
  const [test, setTest] = useState<ResponseMessage | undefined>();

  const [gameState, setGameState] = useState<GameState>(GameState.NONE);

  const { isConnected, testJoin, addMessageListener, sendMessage } =
    useGameWebsocket();
  const handleStart = () => {
    testJoin(username, "duel");
    console.log(`Starting game for user: ${username}`);
  };

  useEffect(() => {
    if (isConnected) {
      addMessageListener((message: ResponseMessage) => {
        switch (message) {
          case message.data:
          default:
        }

        console.log("inuse effect", message);
        setTest(message);
      });
    }
  }, [isConnected]);

  const handleSendMessage = () => {
    sendMessage(word);
    setGameState(GameState.ENDING);
  };

  const components = {
    [GameState.NONE]: <></>,
    [GameState.QUEUEING]: (
      <>
        <div>chờ chút</div>
      </>
    ),
    [GameState.PLAYING]: (
      <>
        <div> từ của bạn là : {test?.msg}</div>
      </>
    ),
    [GameState.ENDING]: <></>,
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-r from-green-400 to-blue-500">
      <div className="bg-white p-8 rounded-lg shadow-2xl w-full max-w-md">
        <h1 className="text-4xl font-bold text-center text-gray-800 mb-6">
          Nối từ
        </h1>
        {/* {components[gameState]} */}
        {test?.isPlaying ? (
          <>
            <p className="text-center text-gray-600 mb-6">
              Từ của bạn là {test?.msg}
            </p>
            <div className="space-y-4">
              <Input
                type="text"
                placeholder="Nhập tên của bạn"
                value={word}
                onChange={(e) => setWord(e.target.value)}
                className="w-full"
              />
              <Button
                onClick={handleSendMessage}
                className="w-full bg-gradient-to-r from-green-400 to-blue-500 hover:from-green-500 hover:to-blue-600 text-white font-bold py-2 px-4 rounded"
              >
                Nhập
              </Button>
            </div>
          </>
        ) : (
          <>
            <p className="text-center text-gray-600 mb-6">
              Chào mừng bạn đến với trò chơi Nối từ! Hãy nhập tên của bạn và bắt
              đầu chơi.
            </p>
            <div className="space-y-4">
              <Input
                type="text"
                placeholder="Nhập tên của bạn"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="w-full"
              />
              <Button
                onClick={handleStart}
                className="w-full bg-gradient-to-r from-green-400 to-blue-500 hover:from-green-500 hover:to-blue-600 text-white font-bold py-2 px-4 rounded"
                disabled={!username.trim()}
              >
                Bắt đầu chơi
              </Button>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
