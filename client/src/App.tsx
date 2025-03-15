import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { VideoPlayer } from './VidePlayer'
import { Button } from './components/ui/button'
import DirectionButton, { DIRECTIONS } from './mainView/DirectionButton'

function App() {
  const [count, setCount] = useState(0)

  const handleDirectionClick = async (direction: number) => {
    try {
      const response = await fetch("/api/direction", {
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

  const upArrow = (
    <svg data-name="1-Arrow Up" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="white">
      <path d="m26.71 10.29-10-10a1 1 0 0 0-1.41 0l-10 10 1.41 1.41L15 3.41V32h2V3.41l8.29 8.29z" />
    </svg>
  );

  const downArrow = (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="white">
      <path d="M26.29 20.29 18 28.59V0h-2v28.59l-8.29-8.3-1.42 1.42 10 10a1 1 0 0 0 1.41 0l10-10z" data-name="2-Arrow Down"/>
    </svg>
  );

  const leftArrow = (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="white">
      <path d="M32 15H3.41l8.29-8.29-1.41-1.42-10 10a1 1 0 0 0 0 1.41l10 10 1.41-1.41L3.41 17H32z" data-name="4-Arrow Left"/>
    </svg>
  );

    // const leftArrow = (
    //   <svg xmlns="http://www.w3.org/2000/svg" width="655.359" height="655.359" viewBox="0 0 6.827 6.827">
    //     <path d="m4.014 1.11-.726.725.726.725-.151.151-.8-.8-.076-.076.075-.076.801-.8z"/>
    //   </svg>
    // );

  const rightArrow = (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="white">
      <path d="m31.71 15.29-10-10-1.42 1.42 8.3 8.29H0v2h28.59l-8.29 8.29 1.41 1.41 10-10a1 1 0 0 0 0-1.41z" data-name="3-Arrow Right"/>
    </svg>
  );

  return (
    <div style={{ display: "grid", gridTemplateAreas: `'up up up' 'left . right' 'down down down'`, justifyContent: 'center', gap: '1px'}}>
      <div style={{ gridArea: "up" }}>
        <DirectionButton direction={DIRECTIONS.UP} onDirectionClick={handleDirectionClick} svg={upArrow} />
      </div>
      <div style={{ gridArea: "down" }}>
        <DirectionButton direction={DIRECTIONS.DOWN} onDirectionClick={handleDirectionClick} svg={downArrow} />
      </div>
      <div style={{ gridArea: "left" }}>
        <DirectionButton direction={DIRECTIONS.LEFT} onDirectionClick={handleDirectionClick} svg={leftArrow} />
      </div>
      <div style={{ gridArea: "right" }}>
        <DirectionButton direction={DIRECTIONS.RIGHT} onDirectionClick={handleDirectionClick} svg={rightArrow} />
      </div>
    </div>
  );
}

export default App
