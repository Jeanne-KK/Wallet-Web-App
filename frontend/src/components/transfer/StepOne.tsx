import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet';


interface StepOneProps {
    setSelectMethod: (num: string) => void;
    setInputAccNo: (num: number) => void;
    setInputBaht: (num: string) => void;
    selectMethod: string;
    inputAccNo: number;
    inputBaht: string;
    errAccNo: string;
    errAmount: string;
}

const StepOne = ({errAmount, errAccNo, setSelectMethod, setInputAccNo, setInputBaht, selectMethod, inputAccNo, inputBaht}: StepOneProps) => {
    

    function handleSelectBank(num: string): void {
        if (selectMethod === num) {
            return
        }
        setInputAccNo(0);
        setInputBaht('');
        setSelectMethod(num);
    }

    function renderInputNo(num: string): JSX.Element {
        switch (num) {
            case "1": {
                return (
                    <div className='mt-5 animate-fadein'> 
                        <div>
                            {errAccNo === "" ? (<span>Account No.</span>) : (<span className='text-red-600'>Account No. {errAccNo}</span>)}
                            <input onChange={(e) => setInputAccNo(Number(e.target.value))} placeholder='Enter bank Account' className='w-full border-b-1 border-[#e1e1e1] rounded p-3 text-base' type='number' required/>
                        </div>
                        <div className='mt-5'>
                            {errAmount === "" ? (<span className=''>Amount</span>) : (<span className='text-red-600'>Amount {errAmount}</span>)}
                            <input value={inputBaht} onChange={(e) => setInputBaht(e.target.value)} placeholder='Baht' className='w-full border-b-1 border-[#e1e1e1] rounded p-3 text-base' type='number' required/>
                        </div>
                    </div>
                )
            }
            case "2": {
                return (
                    <div className='mt-5 animate-fadein'>
                        <div>
                            <span>Account No</span>
                            <input onChange={(e) => setInputAccNo(Number(e.target.value))} placeholder='Enter phone number' className='w-full border-b-1 border-[#e1e1e1] rounded p-3 text-base' type='text' />
                        </div>
                        <div className='mt-5'>
                            <span className=''>Amount</span>
                            <input value={inputBaht} onChange={(e) => setInputBaht(e.target.value)} placeholder='Baht' className='w-full border-b-1 border-[#e1e1e1] rounded p-3 text-base' type='text' />
                        </div>
                    </div>
                )
            }
            default: {
                return (
                    <></>
                )
            }
        }

    }

    return (
        <div className="bg-white p-5 rounded-xl ">
            <div className='w-full md:w-150 mx-auto'>
                <span className="text-xl font-semibold">To</span>
                <div className='flex items-center gap-x-5 mt-5'>
                    <button onClick={() => handleSelectBank("1")} className='flex flex-col items-center cursor-pointer'>
                        <AccountBalanceIcon />
                        <span className={selectMethod === "1" ? "underline" : ""}>Bank</span>
                    </button>

                    {/*
                    <button onClick={() => handleSelectBank(2)} className='flex flex-col items-center cursor-pointer'>
                        <AccountBalanceWalletIcon />
                        <span className={selectMethod === 2 ? "underline" : ""}>PrompPay</span>
                    </button>
                    */}
                </div>
                {renderInputNo(selectMethod)}
            </div>

        </div>
    )
}

export default StepOne