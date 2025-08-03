import WalletIcon from '@mui/icons-material/Wallet';
import axios from 'axios';
import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import validator from 'validator';

const Register = () => {
    const navigate = useNavigate();
    const [email, setEmail] = useState<string>('');
    const [name, setName] = useState<string>('');
    const [surname, setSurname] = useState<string>('');
    const [pass, setPass] = useState<string>('');
    const [phone, setPhone] = useState<string>('');
    const [loading, Setloading] = useState<boolean>(false);

    const handleregister = async() => {
        //      Protect repeat press
        if(loading){
            return
        }else{
            Setloading(true);
        }
        
        console.log("1")
        //      Remove space front and back
        const newMail = validator.trim(email)
        const newPass = validator.trim(pass)
        const newName = validator.trim(name)
        const newSurname = validator.trim(surname)
        const newPhone = validator.trim(phone)

        //      Return when dont have input
        if(newMail === "" || newPass === "" || newName === "" || newSurname === "" || newPhone === ""){
            return
        }
        console.log("2")

        //      validate
        if(!validator.isEmail(newMail)){ 
            return
        }
        if(!validator.isNumeric(newPhone)){
            return
        }
        if(!validator.isAlpha(newName)){
            return
        }
        if(!validator.isAlpha(newSurname)){
            return
        } 
        if(newPass.length < 8){
            return
        }
        console.log(newPhone.length)
        console.log(newPhone.length >= 9)
        console.log(newPhone.length < 11)
        console.log(newPhone.length < 11 && newPhone.length >= 9)
        if(newPhone.length < 9 || newPhone.length > 10){
            return
        }
        if(newName.length > 100){
            return
        }
        if(newSurname.length > 100){
            return
        }
        
        try{
            const res = await axios.post("http://localhost:5000/register", {Mail: newMail, Password: newPass, Name: newName, Surname: newSurname, Phone: newPhone}, {withCredentials: true})
            if(res.data.success){
                navigate("/", {replace: true})
            }

        }catch(err){
            console.error(err)

        }finally{
            Setloading(false)
        }
    }

    return (
        <div className="bg-gradient-to-br from-indigo-400 to-purple-500 min-h-screen flex items-center justify-center">
            <div className="text-base md:text-lg flex flex-col md:w-100 bg-white rounded-xl px-10 py-6 md:py-8 md:px-5">
                <div className='w-12 h-12 md:w-15 md:h-15 bg-gray-800 rounded-lg text-4xl md:text-5xl flex justify-center items-center mx-auto text-white'><WalletIcon fontSize='inherit' /></div>
                <span className="mx-auto text-2xl font-bold">My Wallet</span>
                <span className='font-semibold mt-3'>Email</span>
                <input onChange={(e)=>setEmail(e.currentTarget.value)} type="text" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your email' />
                <span className='font-semibold mt-3'>Name</span>
                <input onChange={(e)=>setName(e.currentTarget.value)} type="text" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your name' />
                <span className='font-semibold mt-3'>Surname</span>
                <input onChange={(e)=>setSurname(e.currentTarget.value)} type="text" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your surname' />
                <span className='font-semibold mt-3'>Phone</span>
                <input onChange={(e)=>setPhone(e.currentTarget.value)} type="text" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your number' />
                <span className='font-semibold mt-3'>password</span>
                <input onChange={(e)=>setPass(e.currentTarget.value)} type="password" className="border-1 border-gray-300 p-2 rounded-lg" placeholder='Enter your password' />
                <button onClick={handleregister} className="cursor-pointer bg-purple-500 hover:bg-purple-600 rounded-lg p-2 text-white font-semibold duration-300 mt-5">Signup</button>
                <span className='mt-3 mx-auto'>Already have an account? <Link className='font-semibold' to="/login">Login</Link></span>
            </div>
        </div>
    )
}

export default Register