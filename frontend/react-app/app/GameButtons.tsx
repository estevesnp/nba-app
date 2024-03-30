import { useState, useEffect } from "react";
import { Player } from "../types/Player";
import teamNameMap from "../resources/teamNameMap.json";

interface GameButtonsProps {
  player: Player | null;
}

type GameButton = {
  text: string;
  isCorrect: boolean;
};

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

function handleButtonClick(
  setButtonsDisabled: Function,
  team: string,
  player: Player | null
) {
  if (!player) {
    return;
  }

  console.log(team === player.team ? "Correct" : "Incorrect");
  setButtonsDisabled(true);
}

export default function GameButtons({ player }: GameButtonsProps) {
  const [buttons, setButtons] = useState<string[]>([]);
  const [buttonsDisabled, setButtonsDisabled] = useState(false);

  useEffect(() => {
    createButtons(setButtons, player);
    setButtonsDisabled(false);
  }, [player]);

  return (
    <div>
      {buttons.map((team, index) => (
        <button
          onClick={() => handleButtonClick(setButtonsDisabled, team, player)}
          key={index}
          disabled={buttonsDisabled}
        >
          {team}
        </button>
      ))}
    </div>
  );
}
