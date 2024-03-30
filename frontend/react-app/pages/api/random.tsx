import type { NextApiRequest, NextApiResponse } from "next";

type ResponseData = {
  id: number;
  name: string;
  position: string;
  team: string;
};

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<ResponseData>
) {
  const response = await fetch("http://127.0.0.1:8080/random");
  const data = await response.json();

  res.status(200).json(data);
}
