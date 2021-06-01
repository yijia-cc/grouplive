import React from "react";
import { List, Button } from "antd";

const AmenityList = (props) => {
  const { amenityList, history } = props;
  const handleCickReserve = (type) => {
    history.push({
      pathname: "/calendar/calendarScheduler",
      state: {
        type,
      },
    });
  };
  return (
    <div>
      <List
        itemLayout="vertical"
        size="large"
        pagination={{
          onChange: (page) => {
            console.log(page);
          },
          pageSize: 3,
        }}
        dataSource={amenityList}
        renderItem={(item) => (
          <List.Item
            key={item.title}
            actions={[
              <Button
                type="primary"
                onClick={() => {
                  handleCickReserve(item.title);
                }}
              >
                Reserve
              </Button>,
            ]}
            extra={<img width={272} alt="logo" src={item.img} />}
          >
            <List.Item.Meta title={<div>{item.title}</div>} />
            {item.content}
          </List.Item>
        )}
      />
    </div>
  );
};

export default AmenityList;
