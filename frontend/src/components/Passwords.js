import PwdContainer from "./PwdContainer";
import { useEffect, useState } from "react";

function Passwords(){
    const getpwds_url = `${process.env.REACT_APP_API_URL}/pwd`;
    const [ data,setData ] = useState(null);
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
                if(data !== json.data){
                    setData(json.data);
                }
                if(token !== json.updated_token){
                    setToken(json.updated_token);
                }
                if(token !== "" && token !== localStorage.getItem("access_token")){
                    localStorage.setItem("access_token",token);
                }
            }catch(error){
                alert(error);
            }
        }
        getData();
        
    },[])
    return(
        <div className="shadow-md shadow-black flex flex-col h-[77vh] p-6 w-screen
                        overflow-y-scroll">
            {data !== null?
            Array.from(data).map((pwd,idx)=>(
                <PwdContainer key={idx} props={pwd}/>
            )):
            <h1>No Saved Passwords</h1>}
        </div>
    );
}

export default Passwords;