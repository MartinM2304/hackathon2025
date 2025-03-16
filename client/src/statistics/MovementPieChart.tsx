import { Pie, PieChart, Sector } from "recharts";
import { PieSectorDataItem } from "recharts/types/polar/Pie";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { useState, useMemo } from "react";
import { DirectionData } from "./MainStats";

const chartConfig = {} satisfies ChartConfig;

// Mapping IDs to movement names
const DIRECTION_LABELS: Record<number, string> = {
  0: "Up",
  1: "Down",
  2: "Left",
  3: "Right",
};

const COLOR_MAP: Record<number, string> = {
  0: "#3b82f6", // Up
  1: "#1e40af", // Down
  2: "#2563eb", // Right
  3: "#1d4ed8", // Left
};

interface MovementPieChartProps {
  data: DirectionData[];
}

export function MovementPieChart({ data }: MovementPieChartProps) {
  const chartData = useMemo(() => {
    return data.map((item) => ({
      name: DIRECTION_LABELS[item.Id] || `Direction ${item.Id}`,
      value: item.Count,
      fill: COLOR_MAP[item.Id] || "#8884d8", // Default color if missing
    }));
  }, [data]);

  const [maxIndex, _setMaxIndex] = useState<number>(() => {
    const maxValue = Math.max(...chartData.map((d) => d.value));
    return chartData.findIndex((d) => d.value === maxValue);
  });

  const [activeIndex, setActiveIndex] = useState<number>(maxIndex);

  const onPieEnter = (_: any, index: number) => setActiveIndex(index);
  const onPieLeave = () => setActiveIndex(maxIndex);

  return (
    <Card className="bg-slate-950 w-full max-w-md border-0">
      <CardHeader className="items-center pb-0">
        <CardTitle className="text-white">Robot Movements</CardTitle>
      </CardHeader>
      <CardContent className="flex-1 pb-0">
        <ChartContainer
          config={chartConfig}
          className="mx-auto aspect-square max-h-[250px]"
        >
          <PieChart>
            <ChartTooltip
              cursor={false}
              content={
                <ChartTooltipContent hideLabel className="text-lg w-xsm" />
              }
            />
            <Pie
              data={chartData}
              dataKey="value"
              nameKey="name"
              innerRadius={60}
              strokeWidth={5}
              onMouseEnter={onPieEnter}
              onMouseLeave={onPieLeave}
              activeIndex={activeIndex}
              activeShape={({
                outerRadius = 0,
                ...props
              }: PieSectorDataItem) => (
                <Sector {...props} outerRadius={outerRadius + 10} />
              )}
            />
          </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
