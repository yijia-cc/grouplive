import { NavLink, withRouter } from "react-router-dom";
import { Layout, Menu } from "antd";
import UserSetting from "./UserSetting";
import logo1 from "./images/20210521_logo3.svg";
import "./index.css";
const { Header } = Layout;

function TopBar(props) {
  const { pathname } = props.location;
  const keyArray = pathname.split("/");
  const selectedKey = keyArray[1] === undefined ? "dashboard" : keyArray[1];
  return (
    <Header
      className="topBar-wrapper"
      style={{ position: "fixed", zIndex: 1, width: "100%" }}
    >
      <img src={logo1} alt="topBar-logo" className="topBar-wrapper-logo" />
      <span className="topBar-title">GroupLive</span>

      <Menu
        className="topBar-right"
        theme="dark"
        mode="horizontal"
        defaultSelectedKeys={["dashboard"]}
        selectedKeys={[selectedKey]}
      >
        <Menu.Item key="dashboard">
          DashBoard
          <NavLink to="/dashboard" />
        </Menu.Item>
        <Menu.Item key="discussion">
          Discussion Board
          <NavLink to="/discussion" />
        </Menu.Item>
        <Menu.Item key="chat">
          Chat Room
          <NavLink to="/chat" />
        </Menu.Item>
        <Menu.Item key="calendar">
          Calendar Schedule
          <NavLink to="/calendar" />
        </Menu.Item>
      </Menu>
      <UserSetting />
    </Header>
  );
}

export default withRouter(TopBar);
