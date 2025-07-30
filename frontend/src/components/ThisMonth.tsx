const ThisMonth = () => {
    return (
        <div className="bg-white p-5 rounded-3xl">
            <span className="text-xl font-semibold">This Month</span>
            <div className="grid grid-cols-2 my-5">
                <div className="flex flex-col items-center">
                    <span className="text-[#667eea] font-bold text-2xl">฿20,200</span>
                    <span className="text-[#666]">Income</span>
                </div>
                <div className="flex flex-col items-center">
                    <span className="text-[#667eea] font-bold text-2xl">฿12,200</span>
                    <span className="text-[#666]">Expenses</span>
                </div> 
            </div>
        </div>
    )
}

export default ThisMonth