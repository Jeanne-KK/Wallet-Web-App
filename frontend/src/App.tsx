import Home from "./pages/Home"
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Transactions from "./pages/Transaction";
import Transfer from "./pages/Transfer";
import Login from "./pages/Login";
import Register from "./pages/Register";

function App() {

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/transaction" element={<Transactions />} />
        <Route path="/transfer" element={<Transfer />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
      </Routes>
    </BrowserRouter>
    
  )
}

export default App
