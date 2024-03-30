import type { NextApiRequest, NextApiResponse } from "next";
import { Player } from "../../types/Player";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Player>
) {
  const response = await fetch("http://127.0.0.1:8080/random");
  const data = await response.json();

  res.status(200).json(data);
}
