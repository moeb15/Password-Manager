import Login from "./components/Login.js";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Register from "./components/Register.js";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" Component={Login}/>
        <Route path="/register" Component={Register}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
