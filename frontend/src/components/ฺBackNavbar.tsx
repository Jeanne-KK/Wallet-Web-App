import ArrowBackIosNewIcon from '@mui/icons-material/ArrowBackIosNew';
import { Link } from 'react-router-dom';

const BackNavbar = () => {
    return (
        <div className="flex w-auto bg-white p-5 rounded-xl md:rounded-3xl items-center justify-between">
            <Link to="/"><ArrowBackIosNewIcon /></Link>
            <div className="text-xl md:text-3xl font-bold text-[#667eea]">My Wallet</div>
            
        </div>
    )
}

export default BackNavbar