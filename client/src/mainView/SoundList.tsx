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
    onSoundClick: (sound: number) => void;
  }
  
  interface SoundListProps {
    onSoundClick: (sound: number) => void;
    style?: React.CSSProperties;
  }
  
  const SoundList: React.FC<SoundListProps> = ({ onSoundClick }) => {
    const [isExpanded, setIsExpanded] = useState(false);
  
    const toggleExpansion = () => {
      setIsExpanded(!isExpanded);
    };
  
    const renderSoundButtons = () => {
      return (
        <div style={{ display: "flex", flexDirection: "column" }}>
          {Object.values(SOUNDS).map((sound) => (
            <Button key={sound} onClick={() => onSoundClick(sound)} className="p-2">
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