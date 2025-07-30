import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet';
import { useState } from 'react';


const StepOne = () => {
    const [selectMethod, setSelectMethod] = useState<number>(0);
    const [inputAccNo, setInputAccNo] = useState<string>('');
    const [inputBaht, setInputBaht] = useState<string>('');

    function handleSelectBank(num: number): void {
        if(selectMethod === num){
            return
        }
        setInputAccNo('');
        setInputBaht('');
        setSelectMethod(num);
    }

    function renderInputNo(num: number): JSX.Element {
        switch (num) {
            case 1: {
                return (
                    <div className='mt-5 animate-fadein'>
                        <div>
                            <span>Account No</span>
                            <input value={inputAccNo} onChange={(e) => setInputAccNo(e.target.value)} placeholder='Enter bank Account' className='w-full border-b-1 border-[#e1e1e1] rounded p-3 text-base' type='text' />
                        </div>
                        <div className='mt-5'>
                            <span className=''>Amount</span>
                            <input value={inputBaht} onChange={(e) => setInputBaht(e.target.value)} placeholder='Baht' className='w-full border-b-1 border-[#e1e1e1] rounded p-3 text-base' type='text' />
                        </div>
                    </div>
                )
            }
            case 2: {
                return (
                    <div className='mt-5 animate-fadein'>
                        <div>
                            <span>Account No</span>
                            <input value={inputAccNo} onChange={(e) => setInputAccNo(e.target.value)} placeholder='Enter phone number' className='w-full border-b-1 border-[#e1e1e1] rounded p-3 text-base' type='text' />
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
        <div className="bg-white p-5 rounded-xl">
            <span className="text-xl font-semibold">To</span>
            <div className='flex items-center gap-x-5 mt-5'>
                <button onClick={() => handleSelectBank(1)} className='flex flex-col items-center cursor-pointer'>
                    <AccountBalanceIcon />
                    <span className={selectMethod === 1 ? "underline" : ""}>Bank</span>
                </button>

                <button onClick={() => handleSelectBank(2)} className='flex flex-col items-center cursor-pointer'>
                    <AccountBalanceWalletIcon />
                    <span className={selectMethod === 2 ? "underline" : ""}>PrompPay</span>
                </button>
            </div>
            {renderInputNo(selectMethod)}
        </div>
    )
}

export default StepOne