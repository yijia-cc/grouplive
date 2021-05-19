import { NavLink } from "react-router-dom";
import { Avatar } from "antd";
import avator from "../images/testAvatar.JPG";
import "./index.css";
const User = () => {
  return (
    <div className="user-wrapper">
      <Avatar src={avator} size={42} />
      <div className="display-user-info">
        <NavLink to="/userinfo">User Info</NavLink>
        <hr />
        <NavLink to="/payment">Payment</NavLink>
      </div>
    </div>
  );
};

export default User;
