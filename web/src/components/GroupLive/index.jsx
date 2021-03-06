import { useState } from "react";
import { Layout } from "antd";
import { TOKEN_KEY } from "../../constants";
import TopBar from "../TopBar";
import Main from "../Main";

const GroupLive = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(
    localStorage.getItem(TOKEN_KEY) ? true : false
  );
  const logout = () => {
    console.log("log out");
    localStorage.removeItem(TOKEN_KEY);
    setIsLoggedIn(false);
  };

  const loggedIn = (token) => {
    console.log(typeof token);    
    if (token) {
      localStorage.setItem(TOKEN_KEY, JSON.stringify(token));
      console.log(localStorage.getItem(TOKEN_KEY))
      setIsLoggedIn(true);
    }
  };
  return (
    <Layout>
      <TopBar isLoggedIn={isLoggedIn} handleLogout={logout} />
      <Main isLoggedIn={isLoggedIn} handleLoggedIn={loggedIn} />
    </Layout>
  );
};

export default GroupLive;
