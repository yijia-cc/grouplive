import React, { useState } from "react";
import { Layout } from "antd";
import TopBar from "../TopBar";
import Main from "../Main";

import { TOKEN_KEY } from "../../constants";

const GroupLive = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(
        localStorage.getItem(TOKEN_KEY) ? true : false
    );

    // When logging out, the browser needs to delete the token from localStorage; no interaction is needed from the server!
    const logout = () => {
        console.log("log out");
        localStorage.removeItem(TOKEN_KEY);
        setIsLoggedIn(false);
    };

    // When logging in, the browser needs to save the token in the localStorage of the brwoser.
    const loggedIn = token => {
        if (token) {
            localStorage.setItem(TOKEN_KEY, token);
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
