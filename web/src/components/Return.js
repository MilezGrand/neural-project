import s from "./Return.module.css";
import { useNavigate } from "react-router-dom";

const Return = () => {
    const navigate = useNavigate();

  return (
      <div className={s.return}>
        <div className={s.btn} onClick={() => { (window.location.pathname == "/info") ? navigate("/lib") : navigate("/")}}>
            <img src="./img/back.png" width="15" height="15"/>
            <span >Назад</span>
        </div>
        </div>
    
  );
};

export default Return;
