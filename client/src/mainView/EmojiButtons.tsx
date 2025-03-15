import SmileEmoji from "@/svg/SmileEmoji";
import EmojiButton, { EMOJIS } from "./EmojiButton";
import AngryEmoji from "@/svg/AngryEmoji";
import SadEmoji from "@/svg/SadEmoji";
import WowEmoji from "@/svg/WowEmoji";

export default function EmojiButtons() {
    
    return (
      <div style={{ display: "flex", gap: "5px", justifyContent: "center", marginTop: "30px" }}>
        <div style={{ gridArea: "smile" }}>
          <EmojiButton emoji={EMOJIS.SMILE} svg={<SmileEmoji/>} />
        </div>
        <div style={{ gridArea: "angry" }}>
          <EmojiButton emoji={EMOJIS.ANGRY} svg={<AngryEmoji/>} />
        </div>
        <div style={{ gridArea: "sad" }}>
          <EmojiButton emoji={EMOJIS.SAD} svg={<SadEmoji />} />
        </div>
        <div style={{ gridArea: "right" }}>
          <EmojiButton emoji={EMOJIS.WOW} svg={<WowEmoji />} />
        </div>
      </div>
    )
}