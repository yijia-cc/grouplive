import { NavLink } from "react-router-dom";
import { Menu } from "antd";

const SubNavBar = (props) => {
  const { location } = props;
  const keyArray = location.pathname.split("/");
  const selectedKey = keyArray[2];
  return (
    <>
      <Menu
        theme="light"
        mode="horizontal"
        defaultSelectedKeys={["amenitylist"]}
        selectedKeys={[selectedKey]}
      >
        <Menu.Item key="amenitylist">
          Amenity List
          <NavLink to="/calendar/amenitylist" />
        </Menu.Item>
        <Menu.Item key="history">
          History
          <NavLink to="/calendar/history" />
        </Menu.Item>
        <Menu.Item key="calendarScheduler">
          Calendar
          <NavLink to="/calendar/calendarScheduler" />
        </Menu.Item>
      </Menu>
    </>
  );
};

export default SubNavBar;
