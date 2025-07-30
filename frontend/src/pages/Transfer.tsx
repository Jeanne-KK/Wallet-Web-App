import { useState } from "react";
import StepOne from "../components/transfer/StepOne"
import StepTwo from "../components/transfer/StepTwo";
import BackNavbar from "../components/ฺBackNavbar"

import ArrowForwardIcon from '@mui/icons-material/ArrowForward';


const Transfer = () => {
    const [step, setStep] = useState<number>(1);

    return (
        <div className="bg-gradient-to-br from-indigo-400 to-purple-500 min-h-screen">
            <div className="flex flex-col px-3 py-3 md:px-20 md:py-3 gap-y-5 md:gap-y-10 text-sm md:text-base">
                <div><BackNavbar /></div>
                {step === 1 ? (
                    <>
                        <div className="flex-grow"><StepOne /></div>
                        <div className="self-end flex items-center gap-x-2 ">
                            <span className="text-white text-base md:text-xl">confirm</span>
                            <button onClick={() => setStep(2)} className=" bg-white w-10 h-10 md:w-14 md:h-14 rounded-4xl text-3xl md:text-5xl flex items-center justify-center cursor-pointer"><ArrowForwardIcon fontSize="inherit" /></button>
                        </div>
                    </>
                ) : (
                    <>
                        <div className="flex-grow"><StepTwo /></div>
                        <div className="self-end flex items-center gap-x-2 ">
                            <span className="text-white text-base md:text-xl">confirm</span>
                            <button onClick={() => setStep(2)} className=" bg-white w-10 h-10 md:w-14 md:h-14 rounded-4xl text-3xl md:text-5xl flex items-center justify-center cursor-pointer"><ArrowForwardIcon fontSize="inherit" /></button>
                        </div>
                    </>
                )}

            </div>
        </div>
    )
}

export default Transfer