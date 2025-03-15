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
  onEmojiClick: (direction: number) => void;
}

const EmojiButton: React.FC<EmojiButtonProps> = ({
  svg,
  emoji,
  onEmojiClick,
}) => {
  return (
    <Button onClick={() => onEmojiClick(emoji)} className="p-4">
      {svg}
    </Button>
  );
};


export default EmojiButton;