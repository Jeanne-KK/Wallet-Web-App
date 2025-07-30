interface Total {
    amount: string
}

const TotalBalance = ({amount}: Total) => {



    return (
        <div className="flex flex-col bg-gradient-to-br from-indigo-400 to-purple-600 rounded-3xl p-5">
            <span className="text-white">Total Balance</span>
            <span className="text-white text-4xl font-bold mt-2">฿{amount}</span>
            <div className="flex gap-x-1 md:gap-x-5 mt-7 ">
                <button className="text-white cursor-pointer px-3 py-2 md:px-4 md:py-1 bg-white/20 hover:bg-white/35 duration-300 rounded-3xl">Send Money</button>
                <button className="text-white cursor-pointer px-5 py-2 md:px-4 md:py-1 bg-white/20 hover:bg-white/35 duration-300 rounded-3xl">Request</button>
                <button className="text-white cursor-pointer px-5 py-2 md:px-4 md:py-1 bg-white/20 hover:bg-white/35 duration-300 rounded-3xl">Top Up</button>
            </div>
        </div>
    )
}

export default TotalBalance