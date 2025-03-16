import { CollaborationChart } from "./CollaborationChart";
import { MovementPieChart } from "./MovementPieChart";

export default function MainStats() {

    return (
        <div className="flex flex-wrap h-screen w-screen flex-col h-full gap-4 bg-slate-900 gap-2 p-2 overflow-y-auto">
            <CollaborationChart />
            <MovementPieChart />
        </div>
    )
}