import Login from "./components/Login.js";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Register from "./components/Register.js";
import Dashboard from "./components/Dashboard.js";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" Component={Login}/>
        <Route path="/register" Component={Register}/>
        <Route path="/home" Component={Dashboard}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
