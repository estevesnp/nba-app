"use client";
import { useState } from "react";
import { Player } from "../types/Player";

export default function Game() {
  const [player, setPlayer] = useState<Player | null>(null);
  const [error, setError] = useState<string | null>(null);

  async function handleButtonClick() {
    setError(null);

    const response = await fetch("/api/random");
    if (!response.ok) {
      setError("Failed to get random player, try again later");
      return;
    }

    const data = await response.json();
    setPlayer(data);
  }

  return (
    <div>
      <button onClick={handleButtonClick}>Get random player</button>

      <p className="errorMessage">{error}</p>

      {player ? (
        <div>
          <img
            src={`https://cdn.nba.com/headshots/nba/latest/260x190/${player.id}.png`}
            alt="Player Picture"
            width={260}
            height={190}
          />
          <h2>{player.name}</h2>
          <p>Position: {player.position}</p>
          <p>Team: {player.team}</p>
        </div>
      ) : (
        <p>Click the button to get a random player</p>
      )}
    </div>
  );
}
