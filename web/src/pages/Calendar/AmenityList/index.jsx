import React from "react";
import { List, Button } from "antd";

const AmenityList = () => {
  const listData = [
    {
      title: `Meeting room`,
      content: `Large high-level meeting room to improve work efficiency`,
      img: `https://spacestor.com/media/scaled_images/insights/amenity_rich/insights_spacestor_021219-072343_medium.jpg`,
    },
    {
      title: `Basketball court`,
      content: `Professional basketball court, instantly improve the level of competition`,
      img: `https://www.downtownmagazinenyc.com/wp-content/uploads/2018/03/SA13-Basketball_revG.jpg`,
    },
    {
      title: `Golf course`,
      content: `Upscale golf course, take you to experience high-end life`,
      img: `https://www.ciderridgegolf.com/images/slideshow/slide4.jpg`,
    },
  ];

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
        dataSource={listData}
        renderItem={(item) => (
          <List.Item
            key={item.title}
            actions={[<Button type="primary">Reserve</Button>]}
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
