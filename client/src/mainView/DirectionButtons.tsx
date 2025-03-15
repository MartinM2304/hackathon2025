import DownArrow from "@/svg/DownArrow";
import RightArrow from "@/svg/RightArrrow";
import LeftArrow from "@/svg/LeftArrow";
import UpArrow from "@/svg/UpArrow";
import { useState } from "react";
import StyledButton from "./StyledButton";

const DIRECTIONS = {
  UP: 0,
  DOWN: 1,
  LEFT: 2,
  RIGHT: 3,
};

export default function DirectionButtons() {

    const [isEnabled, setEnabled] = useState(true);

    const handleDirectionClick = async (direction: number) => {
      if(!isEnabled) return;

      setEnabled(false);
      
      try {
        const apiUrl = import.meta.env.VITE_API_URL;
        const response = await fetch(`${apiUrl}/api/direction`, {
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

      setTimeout(() => {
        setEnabled(true);
      }, 5000);
    };

    return (
      <div className="grid grid-cols-3 grid-rows-2 gap-1 justify-center">
        <div className="col-start-2 col-end-3 row-start-1 row-end-2 flex justify-center">
          <StyledButton svg={<UpArrow />} isEnabled={isEnabled} handler={() => handleDirectionClick(DIRECTIONS.UP)} />
        </div>

        <div className="col-start-1 col-end-2 row-start-2 row-end-3 flex justify-center">
          <StyledButton svg={<LeftArrow />} isEnabled={isEnabled} handler={() => handleDirectionClick(DIRECTIONS.LEFT)}/>
        </div>

        <div className="col-start-3 col-end-4 row-start-2 row-end-3 flex justify-center">
          <StyledButton svg={<RightArrow />} isEnabled={isEnabled} handler={() => handleDirectionClick(DIRECTIONS.RIGHT)}/>
        </div>

        <div className="col-start-2 col-end-3 row-start-3 row-end-3 flex justify-center">
          <StyledButton svg={<DownArrow />} isEnabled={isEnabled} handler={() => handleDirectionClick(DIRECTIONS.DOWN)} />
        </div>
      </div>
    )
}