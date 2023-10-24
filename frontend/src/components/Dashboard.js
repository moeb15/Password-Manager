import Sidebar from "./Sidebar";
import Passwords from "./Passwords";
import AddPassword from "./AddPassword";

function Dashboard(){
    return(
        <div className="w-full h-screen text-3xl text-gray-300 flex flex-row">
            <Sidebar />
            <div className="w-screen">
                <AddPassword/>
                <Passwords />
            </div>
        </div>
    )
}

export default Dashboard;