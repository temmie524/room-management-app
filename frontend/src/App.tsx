import { useEffect } from 'react';
import { Auth } from './components/Auth';
import { Rooms } from './components/Rooms'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import axios from 'axios'
import { CsrfToken } from './types';

function App() {
    useEffect(() => {
        axios.defaults.withCredentials = true
        const getCsrfToken = async () => {
            const  { data } = await axios.get<CsrfToken>(
                `${process.env.REACT_APP_API_URL}/csrf`
            )
            axios.defaults.headers['X-CSRF-TOKEN'] = data.csrf_token
        }
        getCsrfToken()
    },[])
  return (
    <BrowserRouter>
        <Routes>
            <Route path="/" element={<Auth />} />
            <Route path="/rooms" element={<Rooms />} />
        </Routes>
    </BrowserRouter>
  );
}

export default App;
