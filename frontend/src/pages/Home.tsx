import { useEffect, useState } from "react"
import Navbar from "../components/Navbar"
import QuickActions from "../components/QuickActions"
import RecentTrans from "../components/RecentTrans"
import ThisMonth from "../components/ThisMonth"
import TotalBalance from "../components/TotalBalance"
import axios from "axios"


interface User {
  name: string;
  surname: string;
  mail: string;
  phone: string;
}


const Home = () => {
    const [info, SetInfo] = useState<User>({name:'',surname:'',mail:'',phone:''})

    useEffect(() => {
        const getUserInfo = async() =>{
            try{
                const res = await axios.post("http://localhost:5000/getUserInfo",{} ,{withCredentials: true})
                if(res.data.success){
                    SetInfo(res.data.Data)
                }
            }catch(err){
                console.error(err)
            }
        }
        getUserInfo()
    },[])
    return (
        <div className="bg-gradient-to-br from-indigo-400 to-purple-500 min-h-screen">
            <div className="flex flex-col px-3 py-3 md:px-20 md:py-3 gap-y-5 md:gap-y-10 text-sm md:text-base">
                <div className=""><Navbar name={info.name} /></div>       
                <div><TotalBalance amount={"2,324.50"} /></div>
                <div className=""><QuickActions /></div>
                <div className=""><RecentTrans /></div>
                <div className=""><ThisMonth /></div>
            </div>
        </div>
    )
}

export default Home