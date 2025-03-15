import { useState } from 'react'
import { VideoPlayer } from './VidePlayer'
import { Button } from './components/ui/button'
import DirectionButton, { DIRECTIONS } from './mainView/DirectionButton'
import EmojiButton, { EMOJIS } from './mainView/EmojiButton'
import SoundList from './mainView/SoundList'

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

  const handleEmojiClick = async (emoji: number) => {
    try {
      const response = await fetch("/api/emoji", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ emoji }),
      });

      if (!response.ok) {
        throw new Error("Failed to send emoji");
      }

      console.log(`Emoji ${emoji} sent successfully`);
    } catch (error) {
      console.error("Error sending emoji:", error);
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

  const handleSoundClick = async (sound: number) => {
    try {
      const response = await fetch("/api/sound", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ sound }),
      });

      if (!response.ok) {
        throw new Error("Failed to send sound");
      }

      console.log(`Sound ${sound} sent successfully`);
    } catch (error) {
      console.error("Error sending sound:", error);
    }
  };

  //rendering of buttons
  const renderDirectionButtons = () => {
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
    )
  }

  const renderEmojiButtons = () => {
    const smileEmoji = (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="white">
        <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3-3-1.34-3-3 1.34-3 3-3zm0 13.5c-2.33 0-4.32-1.45-5.12-3.5h1.05c.71 1.22 1.83 2 3.07 2 1.24 0 2.36-.78 3.07-2h1.05c-.8 2.05-2.79 3.5-5.12 3.5z" />
      </svg>
    );

    const sadEmoji = (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="white">
        <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3-3-1.34-3-3 1.34-3 3-3zm0 13.5c-2.33 0-4.32-1.45-5.12-3.5h1.05c.71-1.22 1.83-2-3.07-2-1.24 0-2.36.78-3.07 2h1.05c-.8-2.05-2.79-3.5-5.12-3.5zM12 16.5c-2.33 0-4.32-1.45-5.12-3.5h1.05c.71 1.22 1.83 2 3.07 2 1.24 0 2.36-.78 3.07-2h1.05c-.8 2.05-2.79 3.5-5.12 3.5z" transform="rotate(180 12 12)"/>
      </svg>
    );

    const angryEmoji = (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="white">
        <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3-3-1.34-3-3 1.34-3 3-3zm0 13.5c-2.33 0-4.32-1.45-5.12-3.5h1.05c.71 1.22 1.83 2 3.07 2 1.24 0 2.36-.78 3.07-2h1.05c-.8 2.05-2.79 3.5-5.12 3.5zM12 16.5c-2.33 0-4.32-1.45-5.12-3.5h1.05c.71 1.22 1.83 2 3.07 2 1.24 0 2.36-.78 3.07-2h1.05c-.8 2.05-2.79 3.5-5.12 3.5zM15 8l-6 0 0 1 6 0z" transform="rotate(-45 12 12)"/>
      </svg>
    );

    const wowEmoji = (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="white">
        <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3-3-1.34-3-3 1.34-3 3-3zm0 13.5c-2.33 0-4.32-1.45-5.12-3.5h1.05c.71 1.22 1.83 2 3.07 2 1.24 0 2.36-.78 3.07-2h1.05c-.8 2.05-2.79 3.5-5.12 3.5zM12 16.5c-2.33 0-4.32-1.45-5.12-3.5h1.05c.71 1.22 1.83 2 3.07 2 1.24 0 2.36-.78 3.07-2h1.05c-.8 2.05-2.79 3.5-5.12 3.5zM15 8l-6 0 0 1 6 0z" transform="rotate(-45 12 12)"/><path d="M12 11c-1.66 0-3-1.34-3-3s1.34-3 3-3 3 1.34 3 3-1.34 3-3 3zm6 0c-1.66 0-3-1.34-3-3s1.34-3 3-3 3 1.34 3 3-1.34 3-3 3z"/>
      </svg>
    );

    return(
      <div style={{ display: "flex", gap: "5px", justifyContent: "center", marginTop: "30px" }}>
        <div style={{ gridArea: "smile" }}>
          <EmojiButton emoji={EMOJIS.SMILE} onEmojiClick={handleEmojiClick} svg={smileEmoji} />
        </div>
        <div style={{ gridArea: "angry" }}>
          <EmojiButton emoji={EMOJIS.ANGRY} onEmojiClick={handleEmojiClick} svg={angryEmoji} />
        </div>
        <div style={{ gridArea: "sad" }}>
          <EmojiButton emoji={EMOJIS.SAD} onEmojiClick={handleEmojiClick} svg={sadEmoji} />
        </div>
        <div style={{ gridArea: "right" }}>
          <EmojiButton emoji={EMOJIS.WOW} onEmojiClick={handleEmojiClick} svg={wowEmoji} />
        </div>
      </div>
    )
  };

  return (
    <div>
      {renderDirectionButtons()}
      {renderEmojiButtons()}
      <SoundList onSoundClick={handleSoundClick} style={{ marginTop: '10px' }}/> {/* Removed empty curly braces */}
    </div>
  );
}

export default App
