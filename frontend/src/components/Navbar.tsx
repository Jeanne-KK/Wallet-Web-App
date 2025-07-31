import { NavLink } from "react-router-dom"
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import { useState } from "react";



const Navbar = () => {
    const [openSetting, SetOpenSetting] = useState<boolean>(false)

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
                                    <button className="md:px-7 py-3 cursor-pointer hover:underline">Profile</button>
                                    <button className="md:px-7 py-3 cursor-pointer hover:underline">Logout</button>
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