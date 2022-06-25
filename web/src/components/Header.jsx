import s from "./Header.module.css";
import { NavLink } from "react-router-dom";

const Header = () => {
  return (
    <div className={s.wrapper}>
      <div className={s.header}>
        <div className={s.logo}>
          <img src="../logo.svg" width="120" height="120"></img>
          <div className={s.name}>
            <span>NAGGETSY</span>
            <p>NEURAL NETWORK</p>
          </div>
        </div>
        <p className={s.subtext}>
          Определите спектр эмоций на фотографии online
        </p>
        
      </div>

      <div className={s.navbar}>
            <ul>
                <li className={s.navItem}>
                    <NavLink activeClassName={s.active} to="/" >
                        Сканировать
                    </NavLink>
                </li>
                <li className={s.navItem}>
                    <NavLink activeClassName={s.active} exact to="/lib" >
                        База данных
                    </NavLink>
                </li>
            </ul>
        </div>
    </div>
  );
};

export default Header;
