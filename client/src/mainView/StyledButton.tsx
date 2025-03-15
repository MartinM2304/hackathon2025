import React from "react";
import { Button } from "@/components/ui/button";

interface EmojiButtonProps {
  svg: React.JSX.Element;
  handler: () => void;
  isEnabled: boolean;
}

const StyledButton: React.FC<EmojiButtonProps> = ({
  svg,
  handler,
  isEnabled
}) => {

  return (
    <Button onClick={handler} disabled={!isEnabled} className="w-14 h-14 cursor-pointer transition-transform duration-200 ease-in-out hover:scale-110 active:scale-90">
      {svg}
    </Button>
  );
};


export default StyledButton;