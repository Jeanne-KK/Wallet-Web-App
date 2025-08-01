import EastIcon from '@mui/icons-material/East';

const StepTwo = () => {
    return (
        <div className="bg-white p-5 rounded-xl flex">
            <div className="w-full md:w-150 mx-auto px-5 py-3">
                <div className="border-b-1 border-gray-200 pb-5">
                    <div className="flex flex-col items-center text-lg">
                        <span className="text-gray-600">Transfer Amount</span>
                        <span className="text-3xl font-bold">฿100.00</span>
                        <span className="text-gray-500">+ ฿0.00 fee</span>
                    </div>
                    <div>
                        <div className="font-semibold text-gray-600">FROM</div>
                        <div className="flex flex-col gap-y-1 mt-3 px-1">
                            <span className="text-lg font-semibold">My Account</span>
                            <div className="text-gray-900">Naranin Pansuk</div>
                            <div className="text-gray-600">100-2-12325-2</div>
                        </div>
                    </div>
                    <div className="flex justify-center">
                        <div className="w-10 h-10 bg-indigo-500 rounded-4xl text-3xl flex items-center justify-center text-white"><EastIcon fontSize='inherit' /></div>
                    </div>
                    <div>
                        <div className="font-semibold text-gray-600">TO</div>
                        <div className="flex flex-col gap-y-1 mt-3 px-1">
                            <span className="text-lg font-semibold">Kakarin Momammy</span>
                            <div className="text-gray-600">100-2-12325-2</div>
                        </div>
                    </div>

                </div>

                <div className="mt-5 border-b-1 border-gray-200 pb-5">
                    <span className="text-lg font-semibold">Details</span>
                    <div className="flex justify-between px-1">
                        <span className="text-gray-600">Amount:</span>
                        <span className="font-semibold">฿100.00</span>
                    </div>
                    <div className="flex justify-between px-1">
                        <span className="text-gray-600">Fee:</span>
                        <span className="text-red-500">฿0.00</span>
                    </div>
                </div>
                <div className="flex justify-between mt-5">
                    <span className="text-lg font-semibold">Total Amount</span>
                    <span className="text-xl font-bold">฿100.00</span>
                </div>
            </div>



        </div>
    )
}

export default StepTwo