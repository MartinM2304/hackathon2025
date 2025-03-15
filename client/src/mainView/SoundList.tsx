import { useState } from "react";
import { Button } from "@/components/ui/button";
import { SERVER_URL } from "@/config";

export const SOUNDS = {
  BARK: 0,
  MEOW_MEOW: 1,
  OPERA: 2,
  DRUM_BEAT: 3,
};

const SOUND_LABELS: Record<number, string> = {
  [SOUNDS.BARK]: "Dog Bark",
  [SOUNDS.MEOW_MEOW]: "Cat Meow",
  [SOUNDS.OPERA]: "Opera Singing",
  [SOUNDS.DRUM_BEAT]: "Drum Beat",
};

  
const SoundList = () => {
    const [isExpanded, setIsExpanded] = useState(false);
    const [isEnabled, setEnabled] = useState(true);
  
    const toggleExpansion = () => {
      setIsExpanded(!isExpanded);
    };
  
    const handleSoundClick = async (sound: number) => {
      if(!isEnabled) return;
      setEnabled(false);

      try {
        const response = await fetch(`${SERVER_URL}/api/sound`, {
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

      setTimeout(() => {
        setEnabled(true);
      }, 5000);
    };

    const renderSoundButtons = () => {
      return (
        <div className="flex flex-col">
          {Object.values(SOUNDS).map((sound) => (
            <Button
              disabled={!isEnabled}
              key={sound}
              onClick={() => handleSoundClick(sound)}
              className="bg-gray-700 my-1 ease-in-out cursor-pointer hover:bg-gray-900 hover:scale-110 active:scale-90"
            >
              {SOUND_LABELS[sound] || "Unknown Sound"}
            </Button>
          ))}
        </div>
      );
    };
  
    return (
      <div>
        <Button onClick={toggleExpansion} className="cursor-pointer" style={{ width: "100%", textAlign: "left" }}>
          Sound
        </Button>
        {isExpanded && renderSoundButtons()}
      </div>
    );
  };
  
  export default SoundList;