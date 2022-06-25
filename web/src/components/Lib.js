import React from 'react';
import s from './Lib.module.css';
import axios from 'axios';
import { useNavigate  } from "react-router-dom";

const Lib = () => {
    const [persons, setPersons] = React.useState([]);
    const navigate = useNavigate();

    React.useEffect(() => {
        const fetchData = async () => {
            var personsResponse = await axios.get(
                "http://localhost:49812/database/persons"
            );
            setPersons(personsResponse.data);
        }
        fetchData();
    }, [])

    const handleClick = async (index) => {
        navigate("/info",{state: {index}});
    }

    return(
        <div className={s.dataBase}>
            <ul>
            {persons.map((o) => 
                <li key={o.Id} onClick={() => handleClick(o.Id)}>{o.Name}</li>
            )}
            </ul>
        </div>
    )
}

export default Lib;