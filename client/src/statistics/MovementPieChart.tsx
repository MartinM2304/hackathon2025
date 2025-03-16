import { Pie, PieChart, Sector } from "recharts"
import { PieSectorDataItem } from "recharts/types/polar/Pie"

import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart"
import { useState } from "react"
const chartData = [
    { name: "Up", value: 275, fill: "#3b82f6" },
    { name: "Down", value: 300, fill: "#1e40af" },
    { name: "Right", value: 187, fill: "#2563eb" },
    { name: "Left", value: 173, fill: "#1d4ed8" },
  ];
  

const chartConfig = {
} satisfies ChartConfig

export function MovementPieChart() {

  const [maxIndex, _setMaxIndex] = useState<number>(() => {
    const maxValue = Math.max(...chartData.map(data => data.value));
    return chartData.findIndex(data => data.value === maxValue);
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
              content={<ChartTooltipContent hideLabel className="text-lg w-xsm" />}
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
  )
}
