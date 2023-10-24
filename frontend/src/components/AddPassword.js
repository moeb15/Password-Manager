import { AiFillAppstore,AiFillLock,AiFillFileAdd } from "react-icons/ai";
import { BsFillKeyFill } from "react-icons/bs";

function AddPassword(){
    return(
        <div className="shadow-md shadow-black flex flex-row h-[20vh] p-6
                        mb-3">
            <div className="bg-[#2a254e] text-lg text-gray-300 w-[184vh]
                        h-full flex flex-col md:flex-row items-center p-3
                        rounded-md mt-3">
            <div className="h-[7vh] mx-3">
                <div className="flex rounded-md items-center
                bg-slate-700 h-full text-left w-[40vh]">
                    <AiFillAppstore size={30} className="mx-3"/>
                    <input type="text" 
                        placeholder="Application Name"
                        className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                </div>
            </div>
            <div className="h-[7vh] pl-[12vh] mx-3">
                <div className="flex rounded-md items-center
                bg-slate-700 h-full text-left w-[40vh]">
                    <AiFillLock size={30} className="mx-3"/>
                    <input type="password" 
                        placeholder="Password"
                        className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                </div>
            </div>
            <div className="h-[7vh] pl-[12vh] mx-3">
                <div className="flex rounded-md items-center
                bg-slate-700 h-full text-left w-[40vh]">
                    <BsFillKeyFill size={30} className="mx-3"/>
                    <input type="password" 
                        placeholder="Encryption Key"
                        className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                </div>
            </div>
            <AiFillFileAdd size={30} className="ml-[15vh] cursor-pointer"/>
            </div>
        </div>
    )
}

export default AddPassword;