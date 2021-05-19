import { NavLink, withRouter } from "react-router-dom";
import { Layout, Menu, Row, Col } from "antd";

import UserSetting from "./UserSetting";
import logo from "./images/GroupLive.svg";
import "./index.css";
const { Header } = Layout;

function TopBar(props) {
  const { pathname } = props.location;
  const keyArray = pathname.split("/");
  const selectedKey = keyArray[1] === undefined ? "dashboard" : keyArray[1];
  return (
    <Header
      className="header-wrtapper"
      style={{ position: "fixed", zIndex: 1, width: "100%" }}
    >
      <Row justify="center">
        <Col span={23} className="header">
          <div className="logo">
            <img src={logo} alt="logo" />
          </div>

          <div className="rightBar">
            <Menu
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
          </div>
        </Col>
      </Row>
    </Header>
  );
}

export default withRouter(TopBar);
