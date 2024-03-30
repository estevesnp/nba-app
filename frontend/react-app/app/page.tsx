"use client";
import { useState } from "react";

type Player = {
  id: number;
  name: string;
  position: string;
  team: string;
};

export default function Page() {
  const [player, setPlayer] = useState<Player | null>(null);

  async function handleButtonClick() {
    const response = await fetch("/api/random");
    const data = await response.json();
    setPlayer(data);
    console.log(data);
  }

  return (
    <div>
      <button onClick={handleButtonClick}>Get random player</button>
      {player ? (
        <div>
          <h2>{player.name}</h2>
          <p>{player.position}</p>
          <p>{player.team}</p>
        </div>
      ) : (
        <p>Click the button to get a random player</p>
      )}
    </div>
  );
}
