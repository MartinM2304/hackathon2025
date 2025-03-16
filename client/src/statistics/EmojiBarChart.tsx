import { Bar, BarChart, CartesianGrid, LabelList, XAxis } from "recharts"

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

const chartData = [
  { value: "Happy", freq: 186 },
  { value: "Sad", freq: 305 },
  { value: "Wow", freq: 237 },
  { value: "Angry", freq: 73 },
]

const chartConfig = {
  desktop: {
    label: "Desktop",
  },
} satisfies ChartConfig

export function EmojiBarChart() {
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
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="value"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Bar dataKey="freq" fill="#1e40af" radius={8}>
              <LabelList
                position="top"
                offset={12}
                className="text-white text-lg text-bold"
                fontSize={12}
              />
            </Bar>
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}
