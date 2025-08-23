import { useState } from "react";
import StepOne from "../components/transfer/StepOne"
import StepTwo from "../components/transfer/StepTwo";
import BackNavbar from "../components/ฺBackNavbar"
import validator from 'validator';
import axios from "axios";
import { useLocation, useNavigate } from "react-router-dom";

import ArrowForwardIcon from '@mui/icons-material/ArrowForward';


const Transfer = () => {
    const navigate = useNavigate();
    const [step, setStep] = useState<number>(1);
    const [selectMethod, setSelectMethod] = useState<string>("");
    const [inputAccNo, setInputAccNo] = useState<number>(0);
    const [inputBaht, setInputBaht] = useState<string>('');
    const [userName, setUserName] = useState<string>('');
    const [desName, setDesName] = useState<string>('');
    const [userAccNo, setUserAccNo] = useState<number>(0);
    const [errAccNo, setErrAccNo] = useState<string>('');
    const [errAmount, setErrAmount] = useState<string>('');
    const location = useLocation();
    
    const confirmTransfer = async () => {
        console.log(userAccNo)
        console.log(selectMethod)
        console.log(inputBaht)
        console.log(inputAccNo)
        try{
            const res = await axios.post("http://localhost:5000/transfer", {w_id: userAccNo, type: selectMethod, amount: inputBaht, to_w_id: inputAccNo}, {withCredentials: true})
            if(res.data.success){
                navigate("/", {replace: true})
            }
        }catch(err){
            console.log(err)
            return
        }
    }

    const getUserAndDesInfo = async (w_id: number, type: string, to_w_id: number) => {
        try{
            const res = await axios.post("http://localhost:5000/beforetransfer", {w_id: w_id, type: type, to_w_id: to_w_id}, {withCredentials: true})
            setUserName(res.data.Data.u_name)
            setDesName(res.data.Data.d_name)
            setStep(2);
        }catch(err){
            console.log(err)
            return
        }
    }

    const nextStep = async () => {
        setErrAccNo("")
        setErrAmount("")
        //      get acc id from /home
        const acc = location.state?.data;
        setUserAccNo(Number(acc))
        const balance = location.state?.balance;

        //      show err when data not true
        if(inputAccNo === 0){
            setErrAccNo("*")
            return
        }
        if(inputAccNo === Number(acc)){
            setErrAccNo("can't transfer to same wallet")
            return
        }
        setErrAccNo("")
        if(inputBaht === ""){
            setErrAmount("*")
            return
        }
        const newBaht = validator.trim(inputBaht);
 
        if(!validator.isFloat(newBaht, {min: 1})){
            setErrAmount("minimum transfer of at least 1 THB")
            return
        }
        const decimalPart = newBaht.split(".")[1];
        if(decimalPart && decimalPart.length > 2){
            setErrAmount("cannot have more than 3 decimal")
            return
        }
        if(parseFloat(inputBaht) > parseFloat(balance)){
            setErrAmount("not enough balance")
            return
        }
        setErrAmount("")
        setInputBaht(newBaht);
        await getUserAndDesInfo(Number(acc), selectMethod, inputAccNo) 
    }

    return (
        <div className="bg-gradient-to-br from-indigo-400 to-purple-500 min-h-screen">
            <div className="flex flex-col px-3 py-3 md:px-20 md:py-3 gap-y-5 md:gap-y-10 text-sm md:text-base">
                <div><BackNavbar /></div>
                {step === 1 ? (
                    <>
                        <div className="flex-grow"><StepOne errAmount={errAmount} errAccNo={errAccNo} setSelectMethod={setSelectMethod} setInputAccNo={setInputAccNo} setInputBaht={setInputBaht} selectMethod={selectMethod} inputAccNo={inputAccNo} inputBaht={inputBaht} /></div>
                        {selectMethod != "0" &&
                            <div className="self-end flex items-center gap-x-2 ">
                                <span className="text-white text-base md:text-xl">next</span>
                                <button onClick={nextStep} className=" bg-white w-10 h-10 md:w-14 md:h-14 rounded-4xl text-3xl md:text-5xl flex items-center justify-center cursor-pointer"><ArrowForwardIcon fontSize="inherit" /></button>
                            </div>
                        }

                    </>
                ) : (
                    <>
                        <div className="flex-grow"><StepTwo selectMethod={selectMethod} userAccNo={userAccNo} inputAccNo={inputAccNo} inputBaht={inputBaht} Uname={userName} Dname={desName} /></div>
                        <div className="self-end flex items-center gap-x-2 ">
                            <span className="text-white text-base md:text-xl">confirm</span>
                            <button onClick={confirmTransfer} className=" bg-white w-10 h-10 md:w-14 md:h-14 rounded-4xl text-3xl md:text-5xl flex items-center justify-center cursor-pointer"><ArrowForwardIcon fontSize="inherit" /></button>
                        </div>
                    </>
                )}

            </div>
        </div>
    )
}

export default Transfer