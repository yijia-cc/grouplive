import { Menu, Dropdown, message, Button } from "antd";
import { DownOutlined } from "@ant-design/icons";
import { LeftOutlined, RightOutlined } from "@ant-design/icons";
import "./index.css";
const onClick = ({ key }) => {
  message.info(`Click on item ${key}`);
};

const menu = (
  <Menu onClick={onClick}>
    <Menu.Item key="1">Meeting room</Menu.Item>
    <Menu.Item key="2">Basketball court</Menu.Item>
    <Menu.Item key="3">Golf course</Menu.Item>
  </Menu>
);

const menu2 = (
  <Menu onClick={onClick}>
    <Menu.Item key="1">Room1</Menu.Item>
    <Menu.Item key="2">Room2</Menu.Item>
    <Menu.Item key="3">Room3</Menu.Item>
  </Menu>
);

const CalendarHeader = () => {
  return (
    <div className="claendar-header">
      <div className="claendar-header-left">
        <Dropdown overlay={menu} className="claendar-header-amenity">
          <a className="ant-dropdown-link" onClick={(e) => e.preventDefault()}>
            Meeting room <DownOutlined />
          </a>
        </Dropdown>
        <Dropdown overlay={menu2} className="claendar-header-amenity-room">
          <a className="ant-dropdown-link" onClick={(e) => e.preventDefault()}>
            Room1 <DownOutlined />
          </a>
        </Dropdown>
      </div>

      <div className="claendar-header-right">
        <Button className="claendar-header-right-btn">Today</Button>
        <Button
          className="claendar-header-right-btn"
          shape="circle"
          icon={<LeftOutlined />}
        />
        <Button
          className="claendar-header-right-btn"
          shape="circle"
          icon={<RightOutlined />}
        />
      </div>
    </div>
  );
};

export default CalendarHeader;
