import { AiFillAppstore,AiFillLock,AiFillFileAdd } from "react-icons/ai";
import { BsFillKeyFill } from "react-icons/bs";
import { useState } from "react";

function AddPassword(){
    const [ app,setApp ] = useState("");
    const [ pwd,setPwd ] = useState("");
    const [ key,setKey ] = useState("");

    const handleCreate = async(e) => {
        e.preventDefault();
        const add_url = `${`${process.env.REACT_APP_API_URL}/pwd`}`
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
                window.location.reload();
            }
        }catch (error){
            alert(error);
        }
    }

    return(
        <div className="shadow-md shadow-black flex flex-row h-fit lg:h-[20vh] p-6
                        mb-3 w-full justify-center sm:justify-normal">
            <div className="bg-[#2a254e] text-lg text-gray-300 w-fit
                        h-full flex flex-col lg:grid lg:grid-cols-4 items-center p-3
                        rounded-md mt-3">
            <div className="h-[7vh] sm:mx-3 mt-3 lg:mt-0 ">
                <div className="flex flex-row rounded-md items-center
                bg-slate-700 h-full text-left w-[30vh] sm:w-[34vh] lg:w-[40vh]">
                    <AiFillAppstore size={30} className="mx-3"/>
                    <input type="text" 
                        placeholder="Application Name"
                        value={app}
                        onChange={e=>{setApp(e.target.value)}}
                        className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                </div>
            </div>
            <div className="h-[7vh] sm:mx-3 mt-3 lg:mt-0">
                <div className="flex rounded-md items-center
                bg-slate-700 h-full text-left w-[30vh] sm:w-[34vh] lg:w-[40vh]">
                    <AiFillLock size={30} className="mx-3"/>
                    <input type="password" 
                        placeholder="Password"
                        value={pwd}
                        onChange={e=>{setPwd(e.target.value)}}
                        className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                </div>
            </div>
            <div className="h-[7vh] sm:mx-4 mt-3 lg:mt-0">
                <div className="flex rounded-md items-center
                bg-slate-700 h-full text-left w-[30vh] sm:w-[34vh] lg:w-[40vh]">
                    <BsFillKeyFill size={30} className="mx-3"/>
                    <input type="password" 
                        placeholder="Encryption Key"
                        value={key}
                        onChange={e=>{setKey(e.target.value)}}
                        className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                </div>
            </div>
            <AiFillFileAdd size={30} className="lg:ml-[34vh] cursor-pointer mt-3 lg:mt-0"
                           onClick={handleCreate}/>
            </div>
        </div>
    )
}

export default AddPassword;