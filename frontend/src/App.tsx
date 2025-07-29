import Home from "./pages/Home"
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Transactions from "./pages/Transaction";

function App() {

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/transaction" element={<Transactions />} />
      </Routes>
    </BrowserRouter>
    
  )
}

export default App
