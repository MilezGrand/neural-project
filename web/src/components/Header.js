import s from "./Header.module.css";

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
          Определите спектр эмоций online
        </p>
      </div>
    </div>
  );
};

export default Header;
