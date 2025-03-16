import { useEffect, useState } from "react";
import { CollaborationChart } from "./CollaborationChart";
import { EmojiBarChart } from "./EmojiBarChart";
import { MovementPieChart } from "./MovementPieChart";
import { SERVER_URL } from "@/config";
import { SoundsBarChart } from "./SoundsChart";

export interface EntropyData {
  turn: number;
  entropy: number;
}

export interface EmojiData {
  Id: number;
  Count: number;
}

export interface DirectionData {
  Id: number;
  Count: number;
}

export interface SoundData {
  Id: number;
  Count: number;
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

    const interval = setInterval(() => {
      fetchData();
    }, 5000);

    return () => clearInterval(interval);
  }, []);

  if (!statsData) {
    return (
      <div className="h-screen w-screen flex items-center justify-center text-black">
        Loading...
      </div>
    );
  }

  return (
    <div className="h-screen w-screen bg-slate-900 p-4 flex items-center justify-center overflow-y-auto">
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 xl:grid-cols-3 gap-4 max-w-6xl w-full">
        <div className="flex justify-center">
          {statsData.entropy && <CollaborationChart data={statsData.entropy} />}
        </div>
        <div className="flex justify-center">
          <MovementPieChart data={statsData.direction} />
        </div>
        <div className="flex justify-center">
          <EmojiBarChart data={statsData.emoji} />
        </div>
        <div className="flex justify-center">
          <SoundsBarChart data={statsData.sound} />
        </div>
      </div>
    </div>
  );
}
