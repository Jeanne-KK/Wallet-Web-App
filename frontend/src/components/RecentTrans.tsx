const RecentTrans = () => {
    return (
        <div className="bg-white p-5 rounded-3xl">
            <div className="flex justify-between">
                <span className="text-xl font-semibold">Recent Transactions</span>
            </div>
            <div className="flex flex-col my-5 md:my-10 px-2 md:px-5">
                <div className="border-b border-gray-200 py-3">
                    <div className="flex justify-between items-center hover:translate-x-1 duration-300">
                        <div className="flex flex-col ">
                            <span className="font-bold">Salary Payment</span>
                            <span className="text-[#666]">Today, 2:30 PM</span>
                        </div>
                        <div className="">
                            <span className="font-semibold text-[#4ecdc4] text-lg">+฿15,000.00</span>
                        </div>
                    </div>
                </div>
                <div className="border-b border-gray-200 py-3 ">
                    <div className="flex justify-between items-center hover:translate-x-1 duration-300">
                        <div className="flex flex-col">
                        <span className="font-bold">Shopping</span>
                        <span className="text-[#666]">Yesterday, 6:45 PM</span>
                    </div>
                    <div className="">
                        <span className="font-semibold text-[#ff6b6b] text-lg">-฿12,500.00</span>
                    </div>
                    </div>
                    
                </div>
                <div className="border-b border-gray-200 py-3">
                    <div className="flex justify-between items-center hover:translate-x-1 duration-300">
                        <div className="flex flex-col">
                        <span className="font-bold">Electricity Bill</span>
                        <span className="text-[#666]">2 days ago, 10:15 AM</span>
                    </div>
                    <div className="">
                        <span className="font-semibold text-[#ff6b6b] text-lg">-฿2,150.00</span>
                    </div>
                    </div>
                    
                </div>
                <div className="border-b border-gray-200 py-3">
                    <div className="flex justify-between items-center hover:translate-x-1 duration-300">
                        <div className="flex flex-col">
                        <span className="font-bold">Netflix Subscription</span>
                        <span className="text-[#666]">3 days ago, 9:00 AM</span>
                    </div>
                    <div className="">
                        <span className="font-semibold text-[#ff6b6b] text-lg">-฿369.00</span>
                    </div>
                    </div>
                    
                </div>
                <div className="border-b border-gray-200 py-3">
                    <div className="flex justify-between items-center hover:translate-x-1 duration-300">
                        <div className="flex flex-col">
                            <span className="font-bold">Freelance Payment</span>
                            <span className="text-[#666]">5 days ago, 3:20 PM</span>
                        </div>
                        <div className="">
                            <span className="font-semibold text-[#4ecdc4] text-lg">+฿2,500.00</span>
                        </div>
                    </div>
                    
                </div>
            </div>
        </div>
    )
}

export default RecentTrans