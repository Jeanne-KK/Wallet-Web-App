import WalletIcon from '@mui/icons-material/Wallet';
import { Link } from 'react-router-dom';
import axios from 'axios';
import { useState } from 'react';
import validator from 'validator';
import { useNavigate } from 'react-router-dom';


const Login = () => {
    const navigate = useNavigate();
    const [mail, setMail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [loading, setLoading] = useState<boolean>(false);

    const handleLogin = async() =>{
        //      Protect repeat press
        if(loading){
            return
        }else{
            setLoading(true);
        }

        //      Remove space front and back
        const newMail = validator.trim(mail)
        const newPass = validator.trim(password)

        //      Return when dont have input
        if(newMail === "" || newPass === ""){
            return
        }
    
        //      Check mail format
        if(!validator.isEmail(newMail)){
            //console.log("invalid email")
            return
        }
        console.log(mail);
        console.log(newPass);
        try{
            const res = await axios.post("http://localhost:5000/login", {Mail: newMail, Password: newPass}, {withCredentials: true})      
            //console.log(res)
            if(res.data.success){
                navigate("/", {replace: true})
            }
        }catch(err){
            console.error(err)
        }finally{
            setLoading(false)
        }
    }

    return (
        <div className="bg-gradient-to-br from-indigo-400 to-purple-500 min-h-screen flex items-center justify-center">
            <div className="text-base md:text-lg flex flex-col md:w-100 bg-white rounded-xl px-10 py-6 md:py-8 md:px-5">
                <div className='w-12 h-12 md:w-15 md:h-15 bg-gray-800 rounded-lg text-4xl md:text-5xl flex justify-center items-center mx-auto text-white'><WalletIcon fontSize='inherit' /></div>
                <span className="mx-auto text-2xl font-bold">My Wallet</span>
                <span className='font-semibold mt-3'>Email</span>
                <input onChange={(e)=>setMail(e.currentTarget.value)} type="text" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your email' required/>
                <span className='font-semibold mt-3'>password</span>
                <input onChange={(e)=>setPassword(e.currentTarget.value)} type="password" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your password' required/>
                <button onClick={handleLogin} className="cursor-pointer bg-purple-500 hover:bg-purple-600 rounded-lg p-2 text-white font-semibold duration-300 mt-5">Login</button>
                <span className='mt-3 mx-auto'>Don't have an account? <Link className='font-semibold' to="/register">Register</Link></span> 
            </div>
        </div>
    )
}

export default Login