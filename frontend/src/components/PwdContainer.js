import { AiFillAppstore,AiFillLock,AiFillDelete,AiFillEdit } from "react-icons/ai"
import { BsFillKeyFill } from "react-icons/bs";
import { IoCopySharp } from "react-icons/io5";
import { useState } from "react";

function PwdContainer({props}){
    const [ key,setKey ] = useState("");
    const [ pwdData,setPwdData ] = useState({});

    const handleCopy = async(e) => {
        e.preventDefault();
        const fetchpwd_url = "http://localhost:8080/api/pwd/decrypt"
        const req = {
            method:"POST",
            headers:{
                "Content-Type":"application/json",
                Authorization:`Bearer ${localStorage.getItem("access_token")}`,
                Refresh:localStorage.getItem("refresh_token")
            },
            body:JSON.stringify({
                application:props.application,
                key:key,
            })
        }
        try{
            const response = await fetch(fetchpwd_url,req)
            if(response.status === 302){
                const json = await response.json()
                setPwdData(json.data)
                if (json.updated_token !== ""){
                    localStorage.setItem("access_token",json.updated_token)
                }
                await navigator.clipboard.writeText(pwdData.password);
                alert("Password copied to clipboard");
            }else{
                alert("Invalid key");
            }
        }catch(error){
            alert(error);
        }
    }

    return(
        <div className="bg-[#2a254e] text-lg text-gray-300 w-fit
                        h-fit flex flex-col md:flex-row items-center p-3
                        rounded-md mt-3">
            <div className="flex flex-row mx-3">
                <AiFillAppstore size={25} className="mx-3"/>
                <h3>{props.application}</h3>
            </div> 
            <div className="flex flex-row">
                <AiFillLock size={25} className="mx-3"/>
                <h3>{props.password}</h3>
            </div>
            <div className="h-full pl-[15vh]">
                <div className="flex rounded-md items-center
                bg-slate-700 h-full text-left w-[30vh]">
                    <BsFillKeyFill size={25} className="mx-3"/>
                    <input type="password" 
                        placeholder="Key"
                        value={key}
                        onChange={e=>{setKey(e.target.value)}}
                        className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                </div>
            </div>
            <IoCopySharp size={25} className="ml-5 cursor-pointer"
                         onClick={handleCopy}/>
            <AiFillDelete size={30} className="ml-5 cursor-pointer"/>
            <AiFillEdit size={30} className="ml-4 cursor-pointer"/>
        </div>
    )
}

export default PwdContainer;