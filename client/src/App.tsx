import DirectionButtons from './mainView/DirectionButtons';
import EmojiButtons from './mainView/EmojiButtons';
import LogDisplay from './mainView/LogDisplay';
import SoundList from './mainView/SoundList'
import { VideoPlayer } from './VidePlayer';

function App() {
  return (
    <div className="flex flex-col md:flex-row h-screen bg-gray-900 text-white">
      <VideoPlayer src='https://robot.g8row.xyz/hls/robot.m3u8' />
    
      <div className="flex flex-col w-full gap-6 md:w-96 p-8 bg-gray-800 h-full p-3">
        {/* Log Display */}
          <LogDisplay socketUrl={"https://server.g8row.xyz/socket.io"} />
        {/* Actions Section */}

          <DirectionButtons />
          <EmojiButtons />
          <SoundList />
      </div>
    </div>
  );
}

export default App
