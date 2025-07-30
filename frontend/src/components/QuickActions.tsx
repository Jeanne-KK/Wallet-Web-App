import { Link } from "react-router-dom"

Link

const QuickActions = () => {
    return (
        <div className="bg-white p-5 rounded-3xl">
            <span className="text-xl font-semibold">Quick Actions</span>
            <div className="grid grid-cols-2 md:grid-cols-4 my-5 md:my-10 gap-y-5">
                <Link to="/transfer" className="flex flex-col items-center cursor-pointer animate-fadein hover:-translate-y-1 duration-300">
                    <div className="bg-gradient-to-br from-rose-400 to-orange-300 rounded text-2xl p-1">💸</div>
                    <span>Transfer</span>
                </Link>
                <div className="flex flex-col items-center cursor-pointer animate-fadein hover:-translate-y-1 duration-300">
                    <div className="bg-gradient-to-br from-teal-300 to-emerald-500 rounded text-2xl p-1">💰</div> 
                    <span>Request</span>
                </div>
                <div className="flex flex-col items-center cursor-pointer animate-fadein hover:-translate-y-1 duration-300">
                    <div className="bg-gradient-to-br from-indigo-400 to-purple-600 rounded text-2xl p-1">📱</div>
                    <span>Scan QR</span>
                </div>
                <div className="flex flex-col items-center cursor-pointer animate-fadein hover:-translate-y-1 duration-300">
                    <div className="bg-gradient-to-br from-yellow-300 to-pink-300 rounded text-2xl p-1">🧾</div>
                    <span>Pay Bills</span>
                </div>
            </div>
        </div>
    )
}

export default QuickActions