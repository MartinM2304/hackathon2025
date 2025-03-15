import React from "react";
import { Button } from "@/components/ui/button";

export const EMOJIS = {
  SMILE: 0,
  SAD: 1,
  ANGRY: 2,
  WOW: 3,
};

interface EmojiButtonProps {
  svg: React.JSX.Element;
  emoji: number;
}

const EmojiButton: React.FC<EmojiButtonProps> = ({
  svg,
  emoji,
}) => {

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

  return (
    <Button onClick={() => handleEmojiClick(emoji)} className="w-12 h-12">
      {svg}
    </Button>
  );
};


export default EmojiButton;