import { CollaborationChart } from "./CollaborationChart";
import { EmojiBarChart } from "./EmojiBarChart";
import { MovementPieChart } from "./MovementPieChart";

export default function MainStats() {
  return (
    <div className="h-screen w-screen bg-slate-900 p-4 flex items-center justify-center overflow-y-auto">
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 xl:grid-cols-3 gap-4 max-w-6xl w-full">
        <div className="flex justify-center">
          <CollaborationChart />
        </div>
        <div className="flex justify-center">
          <MovementPieChart />
        </div>
        <div className="flex justify-center">
          <EmojiBarChart />
        </div>
        <div className="flex justify-center">
          <MovementPieChart />
        </div>
      </div>
    </div>
  );
}
