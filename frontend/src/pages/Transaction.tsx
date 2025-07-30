import { useState } from "react"
import Navbar from "../components/Navbar"
import History from "../components/transaction/History";


const Transactions = () => {
    const [filter, setFilter] = useState<number>(0);

    function handleFilter(num: number) {
        setFilter(num);
    }

    return (
        <div className="bg-gradient-to-br from-indigo-400 to-purple-500 min-h-screen">
            <div className="flex flex-col px-3 py-3 md:px-20 md:py-3 gap-y-5 md:gap-y-10 text-sm md:text-base">
                <div className=""><Navbar /></div>
                <div>
                    <div className="flex gap-x-3">
                        <button onClick={() => handleFilter(0)}
                            className={`py-1 px-5 rounded-4xl border-1 font-semibold text-gray-400 cursor-pointer ${filter === 0 ? "bg-[#6C47FF] text-white border-[#6C47FF]" : "bg-white border-gray-400"}`}
                        >
                            All
                        </button>
                        <button onClick={() => handleFilter(1)}
                            className={`py-1 px-5 rounded-4xl border-1 font-semibold text-gray-400 cursor-pointer ${filter === 1 ? "bg-[#6C47FF] text-white border-[#6C47FF]" : "bg-white border-gray-400"}`}
                        >Transfer/Receive
                        </button>
                    </div>
                    <div className="mt-2"><History /></div>
                </div>
            </div>
        </div>
    )
}

export default Transactions