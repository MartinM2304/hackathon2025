// SoundList.tsx
import React, { useState } from "react";
import { Button } from "@/components/ui/button";

export const SOUNDS = {
  BARK: 0,
  MEOW_MEOW: 1,
  OPERA: 2,
  DRUM_BEAT: 3,
};
interface SoundButtonProps {
    sound: number;
  }
  
  interface SoundListProps {
    style?: React.CSSProperties;
  }
  
  const SoundList: React.FC<SoundListProps> = () => {
    const [isExpanded, setIsExpanded] = useState(false);
  
    const toggleExpansion = () => {
      setIsExpanded(!isExpanded);
    };
  
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

    const renderSoundButtons = () => {
      return (
        <div style={{ display: "flex", flexDirection: "column" }}>
          {Object.values(SOUNDS).map((sound) => (
            <Button key={sound} onClick={() => handleSoundClick(sound)} className="p-2">
              {Object.keys(SOUNDS).find((key) => SOUNDS[key as keyof typeof SOUNDS] === sound)}
            </Button>
          ))}
        </div>
      );
    };
  
    return (
      <div>
        <Button onClick={toggleExpansion} style={{ width: "100%", textAlign: "left" }}>
          Sound
        </Button>
        {isExpanded && renderSoundButtons()}
      </div>
    );
  };
  
  export default SoundList;