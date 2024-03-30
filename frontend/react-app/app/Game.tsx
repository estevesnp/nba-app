"use client";
import { useState } from "react";
import { Player } from "../types/Player";
import PlayerCard from "./PlayerCard";

export default function Game() {
  const [player, setPlayer] = useState<Player | null>(null);
  const [error, setError] = useState<string | null>(null);

  async function handleButtonClick() {
    const response = await fetch("/api/random");
    if (!response.ok) {
      setError("Failed to get random player, try again later");
      return;
    }

    const data = await response.json();
    setPlayer(data);
    setError(null);
  }

  return (
    <div>
      <button onClick={handleButtonClick}>Get random player</button>

      {error && <p className="errorMessage">{error}</p>}

      <PlayerCard player={player} />
    </div>
  );
}
