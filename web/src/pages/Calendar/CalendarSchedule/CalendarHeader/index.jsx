import { useState } from "react";
import { Menu, Dropdown } from "antd";
import { DownOutlined } from "@ant-design/icons";
import "./index.css";

const CalendarHeader = (props) => {
  let type = props.state ? props.state.type : "Meeting room";
  const [amenity, setAmenity] = useState(type);
  const [room, setRoom] = useState("Room1");
  const amenityHandler = ({ key }) => {
    setAmenity(key);
  };
  const roomHandler = ({ key }) => {
    setRoom(key);
  };
  const menu = (
    <Menu selectable={true} onClick={amenityHandler}>
      <Menu.Item key="Meeting room">Meeting room</Menu.Item>
      <Menu.Item key="Basketball court">Basketball court</Menu.Item>
      <Menu.Item key="Golf course">Golf course</Menu.Item>
    </Menu>
  );
  const menu2 = (
    <Menu onClick={roomHandler}>
      <Menu.Item key="Room1">Room1</Menu.Item>
      <Menu.Item key="Room2">Room2</Menu.Item>
      <Menu.Item key="Room3">Room3</Menu.Item>
    </Menu>
  );
  return (
    <div className="calendar-header">
      <div className="calendar-header-left">
        <Dropdown overlay={menu} className="calendar-header-amenity">
          <a className="ant-dropdown-link" onClick={(e) => e.preventDefault()}>
            {amenity} <DownOutlined />
          </a>
        </Dropdown>
        <Dropdown overlay={menu2} className="calendar-header-amenity-room">
          <a className="ant-dropdown-link" onClick={(e) => e.preventDefault()}>
            {room} <DownOutlined />
          </a>
        </Dropdown>
      </div>
    </div>
  );
};

export default CalendarHeader;
