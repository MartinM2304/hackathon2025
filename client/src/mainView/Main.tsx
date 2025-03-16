import { VideoPlayer } from "@/VidePlayer";
import LogDisplay from "./LogDisplay";
import DirectionButtons from "./DirectionButtons";
import EmojiButtons from "./EmojiButtons";
import SoundList from "./SoundList";
import { SERVER_URL } from "@/config";

export default function Main() {
    return (
      <div className="flex flex-col md:flex-row h-screen bg-gray-900 text-white">
        <VideoPlayer src='https://robot.g8row.xyz/hls/robot.m3u8' />
      
        <div className="flex flex-col w-full gap-6 md:w-96 p-8 bg-gray-800 h-full p-3">
          {/* Log Display */}
            <LogDisplay socketUrl={`${SERVER_URL}/socket.io`} />
          {/* Actions Section */}
  
            <DirectionButtons />
            <EmojiButtons />
            <SoundList />
        </div>
      </div>
    )
}