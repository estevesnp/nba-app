import { useEffect } from "react";
import { Player } from "@/types/Player";
import teamNameMapData from "@/resources/teamNameMap.json";

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
  const teamNameMap: { [key: string]: string } = teamNameMapData;
  function handleDisplayChange() {
    if (!displayAnswer) {
      return;
    }

    setRandButtonDisabled(true);

    setTimeout(() => {
      setDisplayAnswer(false);
      displayNextPlayer();
    }, 1500);
  }

  useEffect(() => {
    handleDisplayChange();
  }, [displayAnswer]);

  return (
    <div>
      {displayAnswer && (
        <>
          {isCorrect ? (
            <h2 className="correctAnswer">Correct!</h2>
          ) : (
            <h2 className="incorrectAnswer">Incorrect!</h2>
          )}
          <p>
            The answer was the{" "}
            {player ? teamNameMap[player.team] : "unknown team"}
          </p>
        </>
      )}
    </div>
  );
}
