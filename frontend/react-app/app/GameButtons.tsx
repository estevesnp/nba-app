import { useState, useEffect } from "react";
import { Player } from "../types/Player";
import teamNameMap from "../resources/teamNameMap.json";

interface GameButtonsProps {
  player: Player | null;
  incrementScore: Function;
  setDisplayAnswer: Function;
  setIsCorrect: Function;
}

function createButtons(setButtons: Function, player: Player | null) {
  if (!player) {
    setButtons([]);
    return;
  }

  const teams = teamNameMap.teams;

  const teamButtons: string[] = [];

  while (teamButtons.length < 3) {
    const randomTeam = teams[Math.floor(Math.random() * teams.length)];

    if (randomTeam !== player.team && !teamButtons.includes(randomTeam)) {
      teamButtons.push(randomTeam);
    }
  }

  const randomIndex = Math.floor(Math.random() * (teamButtons.length + 1));
  teamButtons.splice(randomIndex, 0, player.team);

  setButtons(teamButtons);
}

export default function GameButtons({
  player,
  incrementScore,
  setDisplayAnswer,
  setIsCorrect,
}: GameButtonsProps) {
  const [buttons, setButtons] = useState<string[]>([]);
  const [buttonsDisabled, setButtonsDisabled] = useState(false);

  useEffect(() => {
    createButtons(setButtons, player);
    setButtonsDisabled(false);
  }, [player]);

  function handleButtonClick(team: string) {
    if (!player) {
      return;
    }

    if (team === player.team) {
      incrementScore();
      setIsCorrect(true);
      console.log("Correct");
    } else {
      setIsCorrect(false);
      console.log("Incorrect");
    }

    setDisplayAnswer(true);
    setButtonsDisabled(true);
  }

  return (
    <div>
      {buttons.map((team, index) => (
        <button
          onClick={() => handleButtonClick(team)}
          key={index}
          disabled={buttonsDisabled}
        >
          {team}
        </button>
      ))}
    </div>
  );
}
