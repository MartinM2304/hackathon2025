import {
  Bar,
  BarChart,
  CartesianGrid,
  LabelList,
  XAxis,
  YAxis,
} from "recharts";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { useMemo } from "react";
import { SoundData } from "./MainStats";

const chartConfig = {
  desktop: {
    label: "Desktop",
  },
} satisfies ChartConfig;

// Mapping IDs to emoji names
const SOUND_LABELS: Record<number, string> = {
  0: "Bark",
  1: "Meow Meow",
  2: "Hello",
  3: "Bye",
};

interface SoundBarChartProps {
  data: SoundData[];
}

export function SoundsBarChart({ data }: SoundBarChartProps) {
  const chartData = useMemo(() => {
    return data.map((item) => ({
      value: SOUND_LABELS[item.Id] || `Sound ${item.Id}`,
      freq: item.Count,
    }));
  }, [data]);

  const maxValue = Math.max(...chartData.map((d) => d.freq), 10); // Ensure a minimum scale

  return (
    <Card className="bg-slate-950 w-full max-w-2xl border-0">
      <CardHeader>
        <CardTitle className="text-white">Reactions Frequency</CardTitle>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <BarChart
            accessibilityLayer
            data={chartData}
            margin={{
              top: 20,
              right: 20,
              left: 10,
              bottom: 10,
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="value"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
            />
            <YAxis domain={[0, maxValue]} hide />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Bar dataKey="freq" fill="#1e40af" radius={8} barSize={30}>
              <LabelList
                position="top"
                offset={8}
                className="text-white text-sm font-bold"
              />
            </Bar>
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
