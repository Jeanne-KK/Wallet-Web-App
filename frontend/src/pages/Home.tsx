import Navbar from "../components/Navbar"
import QuickActions from "../components/QuickActions"
import RecentTrans from "../components/RecentTrans"
import ThisMonth from "../components/ThisMonth"
import TotalBalance from "../components/TotalBalance"



const Home = () => {
    return (
        <div className="bg-gradient-to-br from-indigo-400 to-purple-500 min-h-screen">
            <div className="flex flex-col px-3 py-3 md:px-20 md:py-3 gap-y-5 md:gap-y-10 text-sm md:text-base">
                <div className=""><Navbar /></div> 
                <div><TotalBalance amount={"2,324.50"} /></div>
                <div className=""><QuickActions /></div>
                <div className=""><RecentTrans /></div>
                <div className=""><ThisMonth /></div>
            </div>

        </div>
    )
}

export default Home