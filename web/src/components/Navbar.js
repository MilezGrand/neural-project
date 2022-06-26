import s from "./Navbar.module.css";
import { NavLink } from "react-router-dom";

const Navbar = () => {
  return (
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
    
  );
};

export default Navbar;
