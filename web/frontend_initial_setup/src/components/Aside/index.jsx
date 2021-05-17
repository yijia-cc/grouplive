import { NavLink, withRouter } from "react-router-dom";
import { Layout, Menu } from "antd";
import {
  CalendarOutlined,
  DashboardOutlined,
  WechatOutlined,
  FireOutlined,
} from "@ant-design/icons";

const { Sider } = Layout;
function Aside(props) {
  const { pathname } = props.location;
  const keyArray = pathname.split("/");
  const selectedKey = keyArray[1] === undefined ? "dashboard" : keyArray[1];
  return (
    <Sider width={200} className="site-layout-background">
      <Menu
        // theme="dark"
        mode="inline"
        defaultSelectedKeys={["dashboard"]}
        selectedKeys={[selectedKey]}
        style={{ height: "100%", borderRight: 0 }}
      >
        <Menu.Item key="dashboard" icon={<DashboardOutlined />}>
          Dash Board
          <NavLink to="/dashboard" />
        </Menu.Item>
        <Menu.Item key="discussion" icon={<FireOutlined />}>
          Discussion Board
          <NavLink to="/discussboard" />
        </Menu.Item>
        <Menu.Item key="chatRoom" icon={<WechatOutlined />}>
          Chat Room
          <NavLink to="/chatRoom" />
        </Menu.Item>
        <Menu.Item key="Calendar" icon={<CalendarOutlined />}>
          Calendar Schedule
          <NavLink to="/Calendar" />
        </Menu.Item>
      </Menu>
    </Sider>
  );
}

export default withRouter(Aside);
