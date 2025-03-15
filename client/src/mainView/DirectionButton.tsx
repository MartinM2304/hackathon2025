import React from "react";
import { Button } from "@/components/ui/button";

export const DIRECTIONS = {
  UP: 0,
  DOWN: 1,
  LEFT: 2,
  RIGHT: 3,
};

interface DirectionButtonProps {
  svg: React.JSX.Element;
  direction: number;
  onDirectionClick: (direction: number) => void;
}

const DirectionButton: React.FC<DirectionButtonProps> = ({
  svg,
  direction,
  onDirectionClick,
}) => {
  return (
    <Button onClick={() => onDirectionClick(direction)} className="p-4">
      {svg}
    </Button>
  );
};


export default DirectionButton;