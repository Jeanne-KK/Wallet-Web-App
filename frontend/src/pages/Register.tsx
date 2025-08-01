import WalletIcon from '@mui/icons-material/Wallet';
import { Link } from 'react-router-dom';

const Register = () => {
    return (
        <div className="bg-gradient-to-br from-indigo-400 to-purple-500 min-h-screen flex items-center justify-center">
            <div className="text-base md:text-lg flex flex-col md:w-100 bg-white rounded-xl px-10 py-6 md:py-8 md:px-5">
                <div className='w-12 h-12 md:w-15 md:h-15 bg-gray-800 rounded-lg text-4xl md:text-5xl flex justify-center items-center mx-auto text-white'><WalletIcon fontSize='inherit' /></div>
                <span className="mx-auto text-2xl font-bold">My Wallet</span>
                <span className='font-semibold mt-3'>Email</span>
                <input type="text" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your email' />
                <span className='font-semibold mt-3'>Name</span>
                <input type="text" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your name' />
                <span className='font-semibold mt-3'>Surname</span>
                <input type="text" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your surname' />
                <span className='font-semibold mt-3'>password</span>
                <input type="password" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your password' />
                <button className="cursor-pointer bg-purple-500 hover:bg-purple-600 rounded-lg p-2 text-white font-semibold duration-300 mt-5">Signup</button>
                <span className='mt-3 mx-auto'>Already have an account? <Link className='font-semibold' to="/login">Login</Link></span>
            </div>
        </div>
    )
}

export default Register