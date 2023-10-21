import { useEffect,useState } from "react";


const useFetch = (url,{props}) => {
    const [ data,setData ] = useState(null);
    const [ err,setError ] = useState(null);

    useEffect(()=>{
        const getData = async() => {
            try{
                const response = await fetch(url,{
                    method: props.method,
                    headers: props.headers,
                    body: props.body
                })
                const response_json = await response.json()
                setData(response_json)
            }catch (error){
                setError(error)
            }
        }
        getData();
    },[url, props])

    return [ data,err ];
}

export default useFetch;