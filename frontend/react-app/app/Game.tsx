"use client";
import { useState } from "react";
import { Player } from "@/types/Player";
import PlayerCard from "./PlayerCard";
import GameButtons from "./GameButtons";
import AnswerDisplay from "./AnswerDisplay";

export default function Game() {
  const [player, setPlayer] = useState<Player | null>(null);
  const [score, setScore] = useState<number>(0);
  const [buttonDisabled, setButtonDisabled] = useState<boolean>(false);
  const [displayAnswer, setDisplayAnswer] = useState<boolean>(false);
  const [isCorrect, setIsCorrect] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  function incrementScore() {
    setScore(score + 1);
  }

  async function handleButtonClick() {
    setButtonDisabled(true);
    const response = await fetch("/api/random");
    if (!response.ok) {
      setError("Failed to get random player, try again later");
      setButtonDisabled(false);
      return;
    }

    const data = await response.json();
    setPlayer(data);
    setError(null);
    setButtonDisabled(false);
  }

  return (
    <div>
      <div>Player Score: {score}</div>
      <button onClick={handleButtonClick} disabled={buttonDisabled}>
        Get random player
      </button>

      {error && <p className="errorMessage">{error}</p>}

      <PlayerCard player={player} />
      <GameButtons
        player={player}
        incrementScore={incrementScore}
        setDisplayAnswer={setDisplayAnswer}
        setIsCorrect={setIsCorrect}
      />
      <AnswerDisplay
        displayNextPlayer={handleButtonClick}
        setRandButtonDisabled={setButtonDisabled}
        setDisplayAnswer={setDisplayAnswer}
        displayAnswer={displayAnswer}
        isCorrect={isCorrect}
        player={player}
      />
    </div>
  );
}
