"use client";

import { Button } from "@/components/ui/button";
import React, { useEffect, useState } from "react";
import { Input } from "@/components/ui/input";
import { useGameWebsocket } from "@/context/game-socket";
import { ResponseMessage, STATE } from "@/type/ws";

export default function Home() {
  const [username, setUsername] = useState("");
  const [word, setWord] = useState("");
  const [test, setTest] = useState<ResponseMessage | undefined>();

  const [gameState, setGameState] = useState<STATE>(STATE.NONE);

  const { isConnected, testJoin, addMessageListener, closeConnection,sendMessage } =
    useGameWebsocket();
  const handleStart = () => {
    testJoin(username, "duel");
  };

  useEffect(() => {
    if (isConnected) {
      addMessageListener((message: ResponseMessage) => {
        setGameState(message.state)
        setTest(message);
      });
    }
  }, [isConnected]);

  const handleSendMessage = () => {
    sendMessage(word);
  };

  const handleSomethingElse = (options: "quit" | "play_again") => {
    sendMessage(options);
    if (options == "quit"){
      closeConnection()
      setGameState(STATE.NONE)
    }
    setWord("");
  };

  console.log("gameState",gameState)

  // Should be state from Backend for everyscreen to easy travel
  const components = {
    // Have an input and button
    // input: name
    // button : play => go to queue
    [STATE.NONE]: (
      <>
        <p className="text-center text-gray-600 mb-6">
          Chào mừng bạn đến với trò chơi Nối từ! Hãy nhập tên của bạn và bắt đầu
          chơi.
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
    ),
    // Queueing
    // if match => go to playing
    [STATE.QUEUEING]: (
      <>
        <div>chờ chút</div>
      </>
    ),

    [STATE.WAITING]: (
      <>
        <div>your score: {test?.score}</div>
        <p className="text-center text-gray-600 mb-6">
          {test?.msg} {test?.currentWord}
        </p>
        <div className="space-y-4">
          chờ xíu
        </div>
      </>
    ),

    // should add score
    [STATE.PLAYING]: (
      <>
        <div>your score: {test?.score}</div>
        <p className="text-center text-gray-600 mb-6">
          {test?.msg} {test?.currentWord}
        </p>
        <div className="space-y-4">
          <Input
            type="text"
            placeholder="Nhập tên của bạn"
            value={word}
            onChange={(e) => setWord(e.target.value)}
            className="w-full"
            onKeyDown={(e) => {
              if (e.key === 'Enter') {
                handleSendMessage();
              }
            }}
          />
          <Button
            onClick={handleSendMessage}
            className="w-full bg-gradient-to-r from-green-400 to-blue-500 hover:from-green-500 hover:to-blue-600 text-white font-bold py-2 px-4 rounded"
          >
            Nhập
          </Button>
        </div>
      </>
    ),
    // Ending:
    // Will ask: Play again or quit
    // will have state like isQuitting in golang
    // if click playAgain => send play back to server
    // if click Quit => remove this user from websocket =>
    [STATE.ENDING]: (
      <>
        END GAME
        <div>
          <Button onClick={() => handleSomethingElse("play_again")}>
            Play again?
          </Button>
          <Button onClick={() => handleSomethingElse("quit")}>Quit</Button>
        </div>
      </>
    ),
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-r from-green-400 to-blue-500">
      <div className="bg-white p-8 rounded-lg shadow-2xl w-full max-w-md">
        <h1 className="text-4xl font-bold text-center text-gray-800 mb-6">
          Nối từ
        </h1>
        {components[gameState]}
      </div>
    </div>
  );
}
