import type { NextApiRequest, NextApiResponse } from "next";
import { Player } from "../../types/Player";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Player | string>
) {
  const controller = new AbortController();
  const timeoutId = setTimeout(() => controller.abort(), 2000);
  const host = process.env.BACKEND_HOST || "localhost";
  const port = process.env.BACKEND_PORT || "9000";

  try {
    const response = await fetch(`http://${host}:${port}/random`, {
      signal: controller.signal,
    });
    if (!response.ok) {
      res.status(500).send("Error fetching player data");
    }
    const data = await response.json();
    res.status(200).json(data);
  } catch (err) {
    if (err instanceof Error) {
      res.status(500).send(err.message);
    } else {
      res.status(500).send("An error occurred");
    }
  } finally {
    clearTimeout(timeoutId);
  }
}
