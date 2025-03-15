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
      const apiUrl = import.meta.env.VITE_API_URL;
      const response = await fetch(`${apiUrl}/direction`, {
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
    <Button onClick={() => handleDirectionClick(direction)} className="w-12 h-12 cursor-pointer transition-transform duration-200 ease-in-out hover:scale-110 active:scale-90">
      {svg}
    </Button>
  );
};


export default DirectionButton;