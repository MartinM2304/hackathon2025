import DirectionButtons from './mainView/DirectionButtons';
import EmojiButtons from './mainView/EmojiButtons';
import LogDisplay from './mainView/LogDisplay';
import SoundList from './mainView/SoundList'
import { VideoPlayer } from './VidePlayer';

function App() {
  return (
    <div className="flex h-screen bg-gray-900 text-white">
      <VideoPlayer src='http://robot.g8row.xyz/hls/robot.m3u8' />
    
      <div className="w-96 bg-gray-800 p-4">
        {/* Chat Section */}
        {/* <div className="mb-6">
          <h2 className="mb-2 text-lg font-semibold">Chat</h2>
          <div className="h-64 overflow-y-auto rounded-lg bg-gray-700 p-3">
            <div className="text-sm">
              <p className="text-gray-300">User1: Hello!</p>
              <p className="text-gray-300">User2: Hi there!</p>
            </div>
          </div>
        </div> */}
        {/* Log Display */}
          <LogDisplay socketUrl="ws://localhost:3000/socket.io/"/>
        {/* Actions Section */}
        <div className="space-y-4">
          <div>
            <DirectionButtons />
            <EmojiButtons />
            <SoundList />
          </div>
        </div>
      </div>
    </div>
  );
}

export default App
