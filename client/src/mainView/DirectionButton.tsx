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
}

const DirectionButton: React.FC<DirectionButtonProps> = ({
  svg,
  direction,
}) => {

  const handleDirectionClick = async (direction: number) => {
    try {
      const response = await fetch("http://localhost:3000/api/direction", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ direction }),
      });

      if (!response.ok) {
        throw new Error("Failed to send direction");
      }

      console.log("Direction sent successfully");
    } catch (error) {
      console.error("Error sending direction:", error);
    }
  };

  return (
    <Button onClick={() => handleDirectionClick(direction)} className="w-12 h-12">
      {svg}
    </Button>
  );
};


export default DirectionButton;