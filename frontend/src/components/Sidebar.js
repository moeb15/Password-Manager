import { AiFillLock } from "react-icons/ai";
import { BsSearch } from "react-icons/bs";
import { MdSpaceDashboard } from "react-icons/md";
import { FaUserAlt } from "react-icons/fa";
import { BiSolidLogOut } from "react-icons/bi";
import { useNavigate } from "react-router-dom";

function Sidebar({prfl,setPrfl}){
    const navigate = useNavigate();
    const menu = [
        {title:"Dashboard", icon:<MdSpaceDashboard className="mr-1"/>,
                            clickFn: e => {
                                e.preventDefault();
                                navigate("/home");
                                window.location.reload();
                            }},
        {title:"Profile", icon:<FaUserAlt className="mr-1"/>, 
                          clickFn: e => setPrfl(!prfl)},
        {title:"Logout", icon:<BiSolidLogOut className="mr-1"/>,
                        clickFn: e => {
                            e.preventDefault();
                            localStorage.clear();
                            navigate("/");
                            window.location.reload();
                        }}
    ]

    return(
        <div className="bg-[#2a254e] h-screen p-5 pt-8 sm:w-[30vh] relative
                        font-bold text-2xl hidden sm:flex flex-col items-center text-left">
            <div>
                <AiFillLock size={30} color="white"
                        className="block"/>
            </div>
            <div className="flex rounded-md items-center 
            bg-slate-700 mt-6 px-4 py-2 text-left w-[25vh]">
                <BsSearch size={15} className="block float-left cursor-pointer" />
                <input type={"search"} 
                       placeholder="Search"
                       className="text-sm w-full text-gray-330 focus:outline-none
                                   bg-transparent ml-2"/>
            </div>
            <ul className="pt-2 w-[25vh] text-left">
                {menu.map((item,idx)=>(
                    <>
                        <li key={idx} className="text-gray-300 text-sm flex items-center
                                                gap=x=4 cursor-pointer p-2 hover:bg-slate-700 
                                                duration-75 rounded-md mt-2"
                            onClick={item.clickFn}>
                            {item.icon}
                            <span className="block float-left">{item.title}</span>
                        </li>
                    </>
                ))}
            </ul>
        </div>
    )
}

export default Sidebar;