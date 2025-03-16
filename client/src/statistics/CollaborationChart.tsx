import { Area, AreaChart, CartesianGrid, XAxis, YAxis } from "recharts";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { EntropyData } from "./MainStats";
import { useEffect, useState } from "react";

const chartConfig = {
  desktop: {
    label: "Entropy",
  },
} satisfies ChartConfig;

interface CollaborationChartProps {
  data: EntropyData[];
}

export function CollaborationChart({ data }: CollaborationChartProps) {
  const [sortedDataByTurn, setSortedDataByTurn] = useState<EntropyData[]>([]);
  useEffect(() => {
    setSortedDataByTurn(data.slice().sort((a, b) => a.turn - b.turn));
  }, [data]);
  return (
    <Card className="bg-slate-950 w-full max-w-3xl border-0">
      <CardHeader>
        <CardTitle className="text-white">
          Temporal Dynamics of Collaborative Control
        </CardTitle>
        <CardDescription>
          Showing the temporal dynamics of collaborative control in ZafiraBot
        </CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <AreaChart data={sortedDataByTurn} margin={{ left: 12, right: 12 }}>
            <CartesianGrid vertical={false} strokeDasharray="3 3" />
            <XAxis
              dataKey="turn"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
            />
            <YAxis />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dot" hideLabel />}
            />
            <Area
              dataKey="entropy"
              type="monotone"
              stroke="#8884d8"
              fillOpacity={0.4}
              fill="#8884d8"
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
