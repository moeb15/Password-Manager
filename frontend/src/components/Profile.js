import { AiFillLock,AiFillDelete } from "react-icons/ai";
import { MdEmail } from "react-icons/md";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

function Profile(){
    const [ email,setEmail ] = useState("");
    const [ pwd,setPwd ] = useState("");
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

    return(
        <div className="w-screen h-screen flex flex-col mt-5 sm:mt-3">
            <div className="w-full h-[21vh] flex flex-col shadow-md shadow-black p-5">
                <h2 className="underline decoration-pink-600 my-2">Delete Account</h2>
                <div className="bg-[#2a254e] w-fit h-full flex flex-col rounded-md
                                sm:grid sm:grid-cols-[50vh,50vh,10vh] items-center justify-center">

                    <div className="flex rounded-md items-center bg-slate-700 
                                h-[5vh] text-left w-[30vh] sm:w-[34vh] lg:w-[35vh] mx-3">
                        <MdEmail size={30} className="mx-3"/>
                        <input type="text"
                               placeholder="Email"
                               value={email}
                               onChange={e=>setEmail(e.target.value)}
                               className="text-sm w-full text-gray-330 focus:outline-none
                                    bg-transparent ml-2"/>
                    </div>

                    <div className="flex rounded-md items-center bg-slate-700 
                                h-[5vh] text-left w-[30vh] sm:w-[34vh] lg:w-[35vh] mx-3">
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
        </div>
    )
}


export default Profile;