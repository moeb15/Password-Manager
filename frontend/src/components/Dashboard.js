import Sidebar from "./Sidebar";
import Passwords from "./Passwords";
import AddPassword from "./AddPassword";
import { FaTimes,FaBars } from "react-icons/fa";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import Profile from "./Profile";

function Dashboard(){
    const [ menu,setMenu ] = useState(false);
    const [ prfl,showPrfl ] = useState(false);
    const [ count,setCount ] = useState(0);

    const handleClick = () => setMenu(!menu)
    const navigate = useNavigate();

    const handleLogout = e => {
        setMenu(false);
        e.preventDefault();
        localStorage.clear();
        navigate("/");
        window.location.reload();
    }

    const handleDashboard = e => {
        setMenu(false);
        e.preventDefault();
        navigate("/home");
        window.location.reload();
    }


    return(
        <div className="w-full h-screen text-3xl text-gray-300 flex flex-row">
            <Sidebar prfl={prfl} setPrfl={showPrfl} />
            <div className="fixed sm:hidden z-10">
                {!menu? 
                <FaBars className="m-2 cursor-pointer"
                               onClick={handleClick}/>:
                <FaTimes className="m-2 cursor-pointer"
                               onClick={handleClick}/>}
            </div>
            <ul className={menu ? "absolute top-0 left-0 w-full h-screen bg-[#2a254e] flex flex-col justify-center items-center" : 
                                "hidden"}>
                <li className="py-6 text-4xl cursor-pointer hover:text-black duration-100"
                    onClick={handleDashboard}>
                    Dashboard
                </li>
                <li className="py-6 text-4xl cursor-pointer hover:text-black duration-100">
                    Profile
                </li>
                <li className="py-6 text-4xl cursor-pointer hover:text-black duration-100"
                    onClick={handleLogout}>
                    Logout
                </li>
            </ul>
            {!prfl ?   
            <div className="w-screen h-screen items-center justify-center flex flex-col">
                <AddPassword count={count} setCount={setCount}/>
                <Passwords count={count} setCount={setCount}/>
            </div>
            :
            <Profile />}
        </div>
    )
}

export default Dashboard;