import { useEffect, useState } from "react";
import { CollaborationChart } from "./CollaborationChart";
import { EmojiBarChart } from "./EmojiBarChart";
import { MovementPieChart } from "./MovementPieChart";
import { SERVER_URL } from "@/config";

export interface EntropyData {
  turn: number;
  entropy: number;
}

interface EmojiData {
  id: number;
  count: number;
}

export interface DirectionData {
  Id: number;
  Count: number;
}

interface SoundData {
  id: number;
  count: number;
}

interface StatsData {
  entropy: EntropyData[];
  emoji: EmojiData[];
  direction: DirectionData[];
  sound: SoundData[];
}

export default function MainStats() {
  const [statsData, setStatsData] = useState<StatsData>();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(`${SERVER_URL}/api/stats`);
        const data: StatsData = await response.json();
        setStatsData(data);
      } catch (error) {
        console.error("Error fetching stats:", error);
      }
    };

    fetchData();
  }, []);

  if (!statsData) {
    return (
      <div className="h-screen w-screen flex items-center justify-center text-white">
        Loading...
      </div>
    );
  }

  return (
    <div className="h-screen w-screen bg-slate-900 p-4 flex items-center justify-center overflow-y-auto">
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 xl:grid-cols-3 gap-4 max-w-6xl w-full">
        <div className="flex justify-center">
          <CollaborationChart data={statsData.entropy} />
        </div>
        <div className="flex justify-center">
          <MovementPieChart data={statsData.direction} />
        </div>
        <div className="flex justify-center">
          <EmojiBarChart />
        </div>
      </div>
    </div>
  );
}
