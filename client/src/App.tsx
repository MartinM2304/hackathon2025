import Main from "./mainView/Main";
import {BrowserRouter, Route, Routes} from 'react-router'
import MainStats from "./statistics/MainStats";

function App() {
  return (
    <BrowserRouter>
    <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/statistics" element={<MainStats />} />
    </Routes>
    </BrowserRouter>
  );
}

export default App
