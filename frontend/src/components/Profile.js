import { AiFillLock,AiFillDelete } from "react-icons/ai";
import { MdEmail } from "react-icons/md";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { BsFillSaveFill } from "react-icons/bs";

function Profile(){
    const [ email,setEmail ] = useState("");
    const [ pwd,setPwd ] = useState("");
    const [ oldPwd,setOld ] = useState("");
    const [ newPwd,setNew ] = useState("");
    const navigate = useNavigate();

    const handleDel = async(e) => {
        e.preventDefault();
        const del_url = `${process.env.REACT_APP_API_URL}/user/remove`
        const req = {
            method:"POST",
            headers:{
                "Content-Type":"application/json",
                Authorization:`Bearer ${localStorage.getItem("access_token")}`,
                "Refresh":localStorage.getItem("refresh_token")
            },
            body:JSON.stringify({
                email:email,
                password:pwd
            })
        }

        try{
            const response = await fetch(del_url, req)
            if(response.status === 404){
                localStorage.clear();
                navigate("/");
                window.location.reload();
                alert("Account deleted");
            }else{
                alert("Invalid credentials")
            }
        }catch (error){
            alert(error);
        }
    }

    const handleUpdate = async(e) => {
        e.preventDefault();
        const update_url = `${process.env.REACT_APP_API_URL}/user/update`
        const req = {
            method:"POST",
            headers:{
                "Content-Type":"application/json",
                Authorization:`Bearer ${localStorage.getItem("access_token")}`,
                "Refresh":localStorage.getItem("refresh_token")
            },
            body:JSON.stringify({
                password:oldPwd,
                new_password:newPwd
            })
        }

        try{
            const response = await fetch(update_url, req)
            if(response.status === 200){
                const json = await response.json()
                if(json.updated_token !== ""){
                    localStorage.setItem("access_token",json.refresh_token)
                }
                alert("Password updated");
            }else{
                alert("Invalid credentials")
            }
        }catch (error){
            alert(error);
        }
    }

    return(
        <div className="w-screen h-screen flex flex-col mt-5 sm:mt-3 ">
            <div className="w-full h-[30vh] sm:h-[25vh] flex flex-col shadow-md shadow-black p-5">
                <h2 className="underline decoration-pink-600 my-2">Delete Account</h2>
                <div className="bg-[#2a254e] w-full sm:w-fit h-[10vh] flex flex-col rounded-md
                                sm:grid sm:grid-cols-[50vh,50vh,10vh] items-center justify-center">

                    <div className="flex rounded-md items-center bg-slate-700 
                                h-[5vh] text-left w-[30vh] sm:w-[34vh] lg:w-[35vh] mx-3 my-2 sm:my-0">
                        <MdEmail size={30} className="mx-3"/>
                        <input type="text"
                               placeholder="Email"
                               value={email}
                               onChange={e=>setEmail(e.target.value)}
                               className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                    </div>

                    <div className="flex rounded-md items-center bg-slate-700 
                                h-[5vh] text-left w-[30vh] sm:w-[34vh] lg:w-[35vh] mx-3 my-2 sm:my-0">
                        <AiFillLock size={30} className="mx-3"/>
                        <input type="password"
                               placeholder="Password"
                               value={pwd}
                               onChange={e=>setPwd(e.target.value)}
                               className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                    </div>
                    <AiFillDelete size={30} className="cursor-pointer"
                                  onClick={handleDel}/>
                </div>
            </div>
            <div className="w-full h-[30vh] sm:h-[25vh] flex flex-col shadow-md shadow-black p-5">
                <h2 className="underline decoration-pink-600 my-2">Change Password</h2>
                <div className="bg-[#2a254e] w-full sm:w-fit h-full flex flex-col rounded-md
                                sm:grid sm:grid-cols-[50vh,50vh,10vh] items-center justify-center">

                    <div className="flex rounded-md items-center bg-slate-700 
                                h-[5vh] text-left w-[30vh] sm:w-[34vh] lg:w-[35vh] mx-3 my-2 sm:my-0">
                        <AiFillLock size={30} className="mx-3"/>
                        <input type="password"
                               placeholder="Password"
                               value={oldPwd}
                               onChange={e=>setOld(e.target.value)}
                               className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                    </div>

                    <div className="flex rounded-md items-center bg-slate-700 
                                h-[5vh] text-left w-[30vh] sm:w-[34vh] lg:w-[35vh] mx-3 my-2 sm:my-0">
                        <AiFillLock size={30} className="mx-3"/>
                        <input type="password"
                               placeholder="New Password"
                               value={newPwd}
                               onChange={e=>setNew(e.target.value)}
                               className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                    </div>
                    <BsFillSaveFill size={25} className="cursor-pointer"
                                  onClick={handleUpdate}/>
                </div>
            </div>
        </div>
    )
}


export default Profile;