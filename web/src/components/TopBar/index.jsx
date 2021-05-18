import { NavLink, withRouter } from "react-router-dom";
import { Layout, Menu, Row, Col } from "antd";
import {
  CalendarOutlined,
  DashboardOutlined,
  WechatOutlined,
  FireOutlined,
  UserOutlined,
} from "@ant-design/icons";
import logo from "./images/logo.svg";
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
        <Col span={22} className="header">
          <div>
            <img src={logo} className="logo" alt="logo" />
          </div>

          <Menu
            theme="dark"
            mode="horizontal"
            defaultSelectedKeys={["dashboard"]}
            className="navigation"
            selectedKeys={[selectedKey]}
          >
            <Menu.Item key="dashboard" icon={<DashboardOutlined />}>
              DashBoard
              <NavLink to="/dashboard" />
            </Menu.Item>
            <Menu.Item key="discussion" icon={<FireOutlined />}>
              Discussion Board
              <NavLink to="/discussion" />
            </Menu.Item>
            <Menu.Item key="chat" icon={<WechatOutlined />}>
              Chat Room
              <NavLink to="/chat" />
            </Menu.Item>
            <Menu.Item key="calendar" icon={<CalendarOutlined />}>
              Calendar Schedule
              <NavLink to="/calendar" />
            </Menu.Item>
          </Menu>
        </Col>
      </Row>
    </Header>
  );
}

export default withRouter(TopBar);
