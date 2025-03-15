import { SmileEmoji } from "@/svg/SmileEmoji";
import { AngryEmoji } from "@/svg/AngryEmoji";
import { SadEmoji } from "@/svg/SadEmoji";
import { WowEmoji } from "@/svg/WowEmoji";
import { useState } from "react";
import StyledButton from "./StyledButton";
import { SERVER_URL } from "@/config";

const EMOJIS = {
  SMILE: 0,
  SAD: 1,
  ANGRY: 2,
  WOW: 3,
};

export default function EmojiButtons() {
  const [isEnabled, setEnabled] = useState(true);

  const handleEmojiClick = async (emoji: number) => {
    if (!isEnabled) return;

    setEnabled(false);

    try {
      const response = await fetch(`${SERVER_URL}/api/emoji`, {
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

    setTimeout(() => {
      setEnabled(true);
    }, 5000);
  };

  return (
    <div
      style={{
        display: "flex",
        gap: "5px",
        justifyContent: "center",
        marginTop: "30px",
      }}
    >
      <div style={{ gridArea: "smile" }}>
        <StyledButton
          isEnabled={isEnabled}
          svg={<SmileEmoji />}
          handler={() => {
            handleEmojiClick(EMOJIS.SMILE);
          }}
        />
      </div>
      <div style={{ gridArea: "angry" }}>
        <StyledButton
          isEnabled={isEnabled}
          svg={<AngryEmoji />}
          handler={() => {
            handleEmojiClick(EMOJIS.ANGRY);
          }}
        />
      </div>
      <div style={{ gridArea: "sad" }}>
        <StyledButton
          isEnabled={isEnabled}
          svg={<SadEmoji />}
          handler={() => {
            handleEmojiClick(EMOJIS.SAD);
          }}
        />
      </div>
      <div style={{ gridArea: "right" }}>
        <StyledButton
          isEnabled={isEnabled}
          svg={<WowEmoji />}
          handler={() => {
            handleEmojiClick(EMOJIS.WOW);
          }}
        />
      </div>
    </div>
  );
}
