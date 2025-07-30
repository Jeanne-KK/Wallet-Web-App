import StepOne from "../components/transfer/StepOne"
import BackNavbar from "../components/ฺBackNavbar"

import ArrowForwardIcon from '@mui/icons-material/ArrowForward';


const Transfer = () =>{
    

    return (
        <div className="bg-gradient-to-br from-indigo-400 to-purple-500 min-h-screen">
            <div className="flex flex-col h-screen px-3 py-3 md:px-20 md:py-3 gap-y-5 md:gap-y-10 text-sm md:text-base">
                <div><BackNavbar /></div>
                <div className="flex-grow"><StepOne /></div> 
                <div className="self-end flex items-center gap-x-2 cursor-pointer">
                    <span className="text-white text-base">confirm</span>
                    <div className=" bg-white rounded-3xl"><ArrowForwardIcon fontSize="large" /></div>
                </div>
            </div>
        </div>
    )
}

export default Transfer