import { useState } from "react";

function Login(){
    const [ user,setUser ] = useState("");
    const [ pwd,setPwd ] = useState("");

    const handleSubmit = e =>{
        e.preventDefault();
        alert(`${user} ${pwd}`)
    }

    return(
        <div className="w-full h-screen text-center items-center
                        text-gray-300 flex flex-col justify-center">
            <form className="flex flex-col w-[30vh]"
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
        </div>
    )
}

export default Login;