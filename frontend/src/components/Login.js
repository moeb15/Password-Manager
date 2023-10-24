import { useState } from "react";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

function Login(){
    const [ user,setUser ] = useState("");
    const [ pwd,setPwd ] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async(e) => {
        e.preventDefault();
        const login_url = "http://localhost:8080/auth/login";
        let req = {
            method:"POST",
            headers:{
                "Content-Type":"application/json"
            },
            body:JSON.stringify({
                username:user,
                password:pwd
            })
        }
        
        try{
            const response = await fetch(login_url,req);
            const json = await response.json();
            if(response.status/100 === 2){
                localStorage.setItem("access_token",json.access_token);
                localStorage.setItem("refresh_token",json.refresh_token);
                navigate("/home");
            }else{
                alert("invalid credentials");
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

                <button className="bg-black hover:bg-gray-700 duration-100 my-3">
                    Login
                </button>
            </form>
            <Link to="/register"
                  className="text-sm underline text-white m-3">
                    Don't have an account? Register here
            </Link>
        </div>
    )
}

export default Login;