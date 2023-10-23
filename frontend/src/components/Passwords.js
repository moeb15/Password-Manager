import PwdContainer from "./PwdContainer";
import { useEffect, useState } from "react";

function Passwords(){
    const getpwds_url = "http://localhost:8080/api/pwd";
    const [ data,setData ] = useState([]);
    const [ token,setToken ] = useState("");
    useEffect(()=>{
            const getData = async() =>{
            try{
                const response = await fetch(getpwds_url,{
                    method:"GET",
                    headers:{
                        "Content-Type":"application/json",
                        Authorization:`Bearer ${localStorage.getItem("access_token")}`,
                        "Refresh":localStorage.getItem("refresh_token"),
                    }
                });
                const json = await response.json();
                setData(json.data);
                setToken(json.updated_token);
                if(token !== "" && token !== localStorage.getItem("access_token")){
                    localStorage.setItem("access_token",token);
                }
            }catch(error){
                alert(error);
            }
        }
        getData();
    },[data.userid,token])
    return(
        <div className="shadow-md shadow-black flex flex-col h-[67vh] p-6 w-screen
                        overflow-y-scroll">
            {Array.from(data).map((pwd,idx)=>(
                <PwdContainer key={idx} props={pwd}/>
            ))}
        </div>
    );
}

export default Passwords;