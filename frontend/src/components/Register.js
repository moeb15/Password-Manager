import { useState } from "react";
import { Link } from "react-router-dom";

function Register(){
    const [ user,setUser ] = useState("");
    const [ pwd,setPwd ] = useState("");
    const [ key,setKey ] = useState("");

    const handleSubmit = async(e) => {
        e.preventDefault();
        const login_url = "http://localhost:8080/auth/register";
        let req = {
            method:"POST",
            headers:{
                "Content-Type":"application/json"
            },
            body:JSON.stringify({
                username:user,
                password:pwd,
                masterkey:key
            })
        }
        
        try{
            const response = await fetch(login_url,req);
            if(response.status !== 201){
                alert("Username in use");
            }else{
                alert("Account created");
            }
        }catch (error){
            alert("Internal server error");
        }
    }

    return(
        <div className="w-full h-screen text-center items-center
                        text-gray-300 flex flex-col justify-center">
            <form className="flex flex-col w-[30vh] text-lg"
                  onSubmit={handleSubmit}>
                <label htmlFor="username">Username</label>
                <input type="text"
                       placeholder="Username"
                       value={user}
                       onChange={e=>setUser(e.target.value)}
                       className="bg-black my-3"/>
                
                <label htmlFor="password">Password</label>
                <input type="password"
                       placeholder="Password"
                       value={pwd}
                       onChange={e=>setPwd(e.target.value)}
                       className="bg-black my-3"/>

            <label htmlFor="key">Key</label>
                <input type="password"
                       placeholder="Encryption Key"
                       value={key}
                       onChange={e=>setKey(e.target.value)}
                       className="bg-black my-3"/>

                <button className="bg-black hover:bg-gray-700 duration-100 my-3">
                    Register
                </button>
            </form>
            <Link to="/"
                  className="text-sm underline text-white m-3">
                    Already have an account? Login here
            </Link>
        </div>
    )
}

export default Register;