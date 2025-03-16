import { CollaborationChart } from "./CollaborationChart";
import { EmojiBarChart } from "./EmojiBarChart";
import { MovementPieChart } from "./MovementPieChart";

export default function MainStats() {

    return (
        <div className="flex flex-wrap flex-col h-screen w-screen h-full bg-slate-900 gap-2 p-2 overflow-y-auto">
            <CollaborationChart />
            <MovementPieChart />
            <EmojiBarChart />
            <MovementPieChart />
        </div>
    )
}