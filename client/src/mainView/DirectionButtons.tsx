import DownArrow from "@/svg/DownArrow";
import DirectionButton, { DIRECTIONS } from "./DirectionButton";
import RightArrow from "@/svg/RightArrrow";
import LeftArrow from "@/svg/LeftArrow";
import UpArrow from "@/svg/UpArrow";

export default function DirectionButtons() {
    return (
        <div className="grid grid-cols-3 grid-rows-2 gap-1 justify-center">
        {/* Up Button */}
        <div className="col-start-2 col-end-3 row-start-1 row-end-2 flex justify-center">
          <DirectionButton direction={DIRECTIONS.UP} svg={<UpArrow />} />
        </div>

        {/* Left Button */}
        <div className="col-start-1 col-end-2 row-start-2 row-end-3 flex justify-center">
          <DirectionButton direction={DIRECTIONS.LEFT} svg={<LeftArrow />} />
        </div>

        {/* Right Button */}
        <div className="col-start-3 col-end-4 row-start-2 row-end-3 flex justify-center">
          <DirectionButton direction={DIRECTIONS.RIGHT} svg={<RightArrow />} />
        </div>

        {/* Down Button */}
        <div className="col-start-2 col-end-3 row-start-3 row-end-3 flex justify-center">
          <DirectionButton direction={DIRECTIONS.DOWN} svg={<DownArrow />} />
        </div>
      </div>
    )
}