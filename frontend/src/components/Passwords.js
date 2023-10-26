import PwdContainer from "./PwdContainer";
import { useEffect, useState } from "react";

function Passwords({count,setCount}){
    const [ data,setData ] = useState(null);
    
    useEffect(()=>{
            const getpwds_url = `${process.env.REACT_APP_API_URL}/pwd`;
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
                setData(json.data)
                if(json.updated_token !== ""){
                    localStorage.setItem("access_token",json.updated_token);
                }
            }catch(error){
                alert(error);
            }
        }
        getData();
        
    },[count])

    return(
        <div className="shadow-md shadow-black flex flex-col h-[77vh] p-6 w-screen
                        overflow-y-scroll">
            {data !== null?
            Array.from(data).map((pwd,idx)=>(
                <PwdContainer key={idx} props={pwd} setCount={setCount}/>
            )):
            <h1 className="underline decoration-pink-600">No Saved Passwords</h1>}
        </div>
    );
}

export default Passwords;