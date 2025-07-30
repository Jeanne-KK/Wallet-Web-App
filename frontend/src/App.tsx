import Home from "./pages/Home"
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Transactions from "./pages/Transaction";
import Transfer from "./pages/Transfer";

function App() {

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/transaction" element={<Transactions />} />
        <Route path="/transfer" element={<Transfer />} />
      </Routes>
    </BrowserRouter>
    
  )
}

export default App
