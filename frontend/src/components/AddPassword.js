import { AiFillAppstore,AiFillLock,AiFillFileAdd } from "react-icons/ai";
import { BsFillKeyFill } from "react-icons/bs";
import { useState } from "react";

function AddPassword(){
    const [ app,setApp ] = useState("");
    const [ pwd,setPwd ] = useState("");
    const [ key,setKey ] = useState("");

    const handleCreate = async(e) => {
        e.preventDefault();
        const add_url = "http://localhost:8080/api/pwd"
        try{
            const req = {
                method:"POST",
                headers:{
                    "Content-Type":"application/json",
                    Authorization:`Bearer ${localStorage.getItem("access_token")}`,
                    "Refresh":localStorage.getItem("refresh_token")
                },
                body:JSON.stringify({
                    application:app,
                    password:pwd,
                    key:key
                })
            }

            const response = await fetch(add_url,req)
            if(response.status === 201){
                alert("Password saved")
                const json = await response.json()
                if(json.updated_token !== ""){
                    localStorage.setItem("access_token",json.updated_token)
                }
            }
        }catch (error){
            alert(error);
        }
    }

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
                        value={app}
                        onChange={e=>{setApp(e.target.value)}}
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
                        value={pwd}
                        onChange={e=>{setPwd(e.target.value)}}
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
                        value={key}
                        onChange={e=>{setKey(e.target.value)}}
                        className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                </div>
            </div>
            <AiFillFileAdd size={30} className="ml-[15vh] cursor-pointer"
                           onClick={handleCreate}/>
            </div>
        </div>
    )
}

export default AddPassword;