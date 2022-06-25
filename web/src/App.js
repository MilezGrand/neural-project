import Header from './components/Header'
import Scan from './components/Scan'
import Lib from './components/Lib'
import Info from './components/Info'
import Results from './components/Results'

import { BrowserRouter, Link, Route, Routes } from "react-router-dom";

const App = () => {
  return (
    <BrowserRouter>
      <Header/>
      
      <Routes>
        <Route path="/" element={<Scan/>}></Route>
        <Route path="/lib" element={<Lib/>}></Route>
        <Route path="/info" element={<Info/>}></Route>
        <Route path="/results" element={<Results/>}></Route>
      </Routes>

      </BrowserRouter>
  );
}

export default App;

