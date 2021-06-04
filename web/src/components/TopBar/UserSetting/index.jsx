import { NavLink } from "react-router-dom";
import { Menu, Dropdown, Avatar } from "antd";
import { DownOutlined } from "@ant-design/icons";
import avator from "../images/testAvatar.JPG";
import "./index.css";
const User = (props) => {
  const { handleLogout } = props;
  const pathname = props.pathname.split("/")[1];
  let selectedKey = "";
  if (pathname === "userinfo") {
    selectedKey = "userinfo";
  } else if (pathname === "payment") {
    selectedKey = "payment";
  }
  const menu = (
    <Menu theme="light" selectedKeys={[selectedKey]}>
      <Menu.Item key="userinfo">
        <NavLink to="/userinfo">User Info</NavLink>
      </Menu.Item>
      <Menu.Item key="payment">
        <NavLink to="/payment">Payment</NavLink>
      </Menu.Item>
      <Menu.Item key="logout" danger onClick={handleLogout}>
        Log out
      </Menu.Item>
    </Menu>
  );
  return (
    <Dropdown
      overlay={menu}
      className="user-wrapper"
      placement="bottomCenter"
      arrow
    >
      <a className="ant-dropdown-link" onClick={(e) => e.preventDefault()}>
        <Avatar src={avator} size={42} /> <DownOutlined />
      </a>
    </Dropdown>
  );
};

export default User;
