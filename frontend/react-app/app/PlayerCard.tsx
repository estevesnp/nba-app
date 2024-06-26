import { useState, useEffect } from "react";
import { Player } from "@/types/Player";
import Image from "next/image";

interface PlayerCardProps {
  player: Player | null;
}

export default function PlayerCard({ player }: PlayerCardProps) {
  const [imageLoaded, setImageLoaded] = useState(true);

  useEffect(() => {
    setImageLoaded(true);
  }, [player]);

  return (
    <div>
      {player ? (
        <div>
          {imageLoaded ? (
            <Image
              src={`https://cdn.nba.com/headshots/nba/latest/260x190/${player.id}.png`}
              alt="Player Picture"
              width={260}
              height={190}
              onError={() => setImageLoaded(false)}
            />
          ) : (
            <Image
              src="/default-image.jpg"
              alt="Player Picture"
              width={260}
              height={190}
            />
          )}
          <h2>{player.name}</h2>
        </div>
      ) : (
        <p>Click the button to get a random player</p>
      )}
    </div>
  );
}
