import { NavLink, useNavigate } from "react-router-dom"
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import { useState } from "react";
import axios from "axios";


interface NavbarProps{
    name: string;
}

const Navbar = ({name}: NavbarProps) => {
    const navigate = useNavigate();
    const [openSetting, SetOpenSetting] = useState<boolean>(false)

    const Logout = async () => {
        try{
            const res = await axios.post("http://localhost:5000/logout",{} ,{withCredentials: true})
            if (res.data.success){
                navigate("/login", {replace: true})
            }

        }catch(err){
            console.error("Logout Fail", err)
            navigate("/login", {replace: true})
        }
    }

    return (
        <div className="flex w-auto bg-white p-5 rounded-xl md:rounded-3xl items-center">
            <div className="text-xl md:text-3xl font-bold text-[#667eea]">My Wallet</div>
            <div className="flex ml-auto gap-x-1 md:gap-x-10">
                <NavLink
                    to="/"
                    className={({ isActive }) =>
                        `cursor-pointer rounded-xl hover:bg-indigo-100 py-1 px-1 md:py-2 md:px-5 duration-300 ${isActive && "underline"}`
                    }
                >
                    Dashboard
                </NavLink>
                <NavLink to="/transaction" className={({ isActive }) =>
                    `cursor-pointer rounded-xl hover:bg-indigo-100 py-1 px-1 md:py-2 md:px-5 duration-300 ${isActive && "underline"}`
                }
                >
                    Transactions
                </NavLink>
                <div className="flex items-center">
                    <button onClick={()=>SetOpenSetting(!openSetting)} className="cursor-pointer"><AccountCircleIcon /></button>
                    {openSetting && 
                        <div className="md:relative">
                            <div className='fixed inset-0 z-40 bg-black/30 md:bg-transparent' onClick={() => SetOpenSetting(false)}></div>
                            <div className="absolute right-0 top-15 md:top-11 bg-white shadow-lg z-50 rounded">
                                <div className="text-lg flex flex-col w-screen md:w-auto">
                                    <span className="md:px-7 py-3 whitespace-nowrap text-center">Hi, {name}</span>
                                    <button className="md:px-7 py-3 cursor-pointer hover:underline">Profile</button>
                                    <button onClick={Logout} className="md:px-7 py-3 cursor-pointer hover:underline">Logout</button>
                                </div>
                                
                            </div>
                        </div>
                    }
                </div>  
            </div>
            
        </div>
    )
}

export default Navbar