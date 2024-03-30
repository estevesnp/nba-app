"use client";
import { useState } from "react";
import { Player } from "../types/Player";

export default function Page() {
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
