import { AiFillAppstore,AiFillDelete,AiFillEdit,AiFillLock } from "react-icons/ai"
import { BsFillKeyFill,BsSaveFill } from "react-icons/bs";
import { BiSolidUser } from "react-icons/bi";
import { IoCopySharp } from "react-icons/io5";
import { useState } from "react";

function PwdContainer({props}){
    const [ key,setKey ] = useState("");
    const [ new_pwd,setNewPwd ] = useState("");
    const [ pwdData,setPwdData ] = useState({});
    const [ edit,setEdit ] = useState(false);

    const handleEdit = () => setEdit(!edit);
 
    const handleCopy = async(e) => {
        e.preventDefault();
        const fetchpwd_url = `${process.env.REACT_APP_API_URL}/pwd/decrypt`
        const req = {
            method:"POST",
            headers:{
                "Content-Type":"application/json",
                Authorization:`Bearer ${localStorage.getItem("access_token")}`,
                Refresh:localStorage.getItem("refresh_token")
            },
            body:JSON.stringify({
                application:props.application,
                username:props.username,
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
                navigator.clipboard.writeText(pwdData.password)
                .then(alert("Password copied to clipboard"))
            }else{
                alert("Invalid key");
            }
        }catch(error){
            alert(error);
        }
    }

    const handleDel = async(e) => {
        e.preventDefault();
        const del_url = `${process.env.REACT_APP_API_URL}/pwd?app=${props.application}`
        try{
            const req = {
                method:"DELETE",
                headers:{
                    "Content-Type":"application/json",
                    Authorization:`Bearer ${localStorage.getItem("access_token")}`,
                    Refresh:localStorage.getItem("refresh_token")
                },
            }
            const response = await fetch(del_url,req)
            if(response.status === 404){
                alert("Password deleted");
                window.location.reload();
            }
        }catch(error){
            alert(error);
        }
    }

    const submitEdit = async(e) => {
        setEdit(!edit)
        e.preventDefault();
        const del_url = `${process.env.REACT_APP_API_URL}/pwd`
        try{
            const req = {
                method:"PATCH",
                headers:{
                    "Content-Type":"application/json",
                    Authorization:`Bearer ${localStorage.getItem("access_token")}`,
                    Refresh:localStorage.getItem("refresh_token")
                },
                body:JSON.stringify({
                    application:props.application,
                    username:props.username,
                    password:new_pwd,
                    key:key
                })
            }
            const response = await fetch(del_url,req)
            if(response.status === 200){
                const json = await response.json()
                if (json.updated_token !== ""){
                    localStorage.setItem("access_token",json.updated_token)
                }
                alert("Password updated")
            }
        }catch(error){
            alert(error);
        }
    }

    return(
        <div className="bg-[#2a254e] text-lg text-gray-300 sm:w-fit
                        h-fit flex flex-col md:flex-row items-center p-3
                        rounded-md mt-3 justify-center">
            <div className="flex flex-col items-center w-fit
                            lg:grid lg:grid-cols-[34vh,30vh,40vh,60vh]">
                <div className="flex flex-row">
                    <AiFillAppstore size={25} className="mx-3"/>
                    <h3>{props.application}</h3>
                </div> 
                <div className="flex flex-row my-3 sm:my-0">
                    <BiSolidUser size={25} className="mx-3"/>
                    <h3>{props.username}</h3>
                </div>

                <div className="flex flex-col items-center sm:flex-row sm:my-0 my-3">
                    <div className="h-full pl-[1vh] lg:pl-[10vh]">
                        <div className="flex rounded-md items-center
                        bg-slate-700 h-full text-left sm:w-[30vh]">
                            <AiFillLock size={25} className="mx-3"/>
                            <input type="password" 
                                placeholder="New Password"
                                value={new_pwd}
                                onChange={e=>{setNewPwd(e.target.value)}}
                                className="text-sm w-full text-gray-330 focus:outline-none
                                            bg-transparent ml-2"
                                readOnly={!edit}/>
                        </div>
                    </div>
                </div>

                <div className="flex flex-col items-center sm:flex-row ">
                    <div className="h-full pl-[1vh] lg:pl-[10vh]">
                        <div className="flex rounded-md items-center
                        bg-slate-700 h-full text-left sm:w-[30vh]">
                            <BsFillKeyFill size={25} className="mx-3"/>
                            <input type="password" 
                                placeholder="Key"
                                value={key}
                                onChange={e=>{setKey(e.target.value)}}
                                className="text-sm w-full text-gray-330 focus:outline-none
                                            bg-transparent ml-2"/>
                        </div>
                    </div>
                    {!edit ? 
                    <div className="flex flex-row mt-5 sm:mt-0">
                        <IoCopySharp size={25} className="ml-5 cursor-pointer"
                                    onClick={handleCopy}/>
                        <AiFillDelete size={30} className="ml-5 cursor-pointer"
                                    onClick={handleDel}/>
                        <AiFillEdit size={30} className="ml-4 cursor-pointer"
                                    onClick={handleEdit}/>
                    </div>
                    :
                    <div className="flex flex-row mt-5 sm:mt-0">
                        <BsSaveFill size={30} className="ml-[9vh] cursor-pointer"
                                    onClick={submitEdit}/>
                        <AiFillEdit size={30} className="ml-[2vh] cursor-pointer"
                                    onClick={handleEdit}/>
                    </div>}
                </div>
            </div>
        </div>
    )
}

export default PwdContainer;