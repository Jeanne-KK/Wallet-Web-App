import { NavLink } from "react-router-dom"



const Navbar = () => {
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
            </div>
        </div>
    )
}

export default Navbar