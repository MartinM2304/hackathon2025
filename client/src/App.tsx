import DirectionButtons from './mainView/DirectionButtons';
import EmojiButtons from './mainView/EmojiButtons';
import SoundList from './mainView/SoundList'
import { VideoPlayer } from './VidePlayer';

function App() {
  return (
    <div className="flex flex-col md:flex-row h-screen bg-gray-900 text-white">
      <VideoPlayer src='http://robot.g8row.xyz/hls/robot.m3u8' />
    
      <div className="w-full md:w-96 p-8 bg-gray-800 h-full">
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

        {/* Actions Section */}
        <div className="flex flex-col gap-6 w-full">
            <DirectionButtons />
            <EmojiButtons />
            <SoundList />
        </div>
      </div>
    </div>
  );
}

export default App
