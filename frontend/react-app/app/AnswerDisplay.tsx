import { Player } from "@/types/Player";
import { useState, useEffect } from "react";

type AnswerDisplayProps = {
  displayNextPlayer: Function;
  setRandButtonDisabled: Function;
  setDisplayAnswer: Function;
  displayAnswer: boolean;
  isCorrect: boolean;
  player: Player | null;
};

export default function AnswerDisplay({
  displayNextPlayer,
  setRandButtonDisabled,
  setDisplayAnswer,
  displayAnswer,
  isCorrect,
  player,
}: AnswerDisplayProps) {
  function handleDisplayChange() {
    if (!displayAnswer) {
      return;
    }

    setRandButtonDisabled(true);

    setTimeout(() => {
      setDisplayAnswer(false);
      displayNextPlayer();
    }, 2000);
  }

  useEffect(() => {
    handleDisplayChange();
  }, [displayAnswer]);

  return (
    <div>
      {displayAnswer && (
        <div>
          <p>{isCorrect ? "Correct" : "Incorrect"}</p>
          <p>The answer is {player?.team}</p>
        </div>
      )}
    </div>
  );
}
