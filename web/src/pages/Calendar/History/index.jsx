import { Descriptions, Badge } from "antd";
import "./index.css";
const History = () => {
  const demoData = [
    {
      id: 1,
      amenity: "Meeting room",
      reservation: "2021-05-24",
      startTime: "18:00:00",
      endTime: "20:00:00",
      status: "processing",
    },
    {
      id: 2,
      amenity: "Basketball court",
      reservation: "2021-04-20",
      startTime: "14:00:00",
      endTime: "16:00:00",
      status: "success",
    },
    {
      id: 3,
      amenity: "Golf course",
      reservation: "2020-09-23",
      startTime: "10:00:00",
      endTime: "17:00:00",
      status: "success",
    },
  ];
  return (
    <>
      {demoData.map((ele) => (
        <Descriptions bordered key={ele.id} className="calendar-history-item">
          <Descriptions.Item label="Amenity" span={2}>
            {ele.amenity}
          </Descriptions.Item>
          <Descriptions.Item label="Reservation Date" span={2}>
            {ele.reservation}
          </Descriptions.Item>
          <Descriptions.Item label="Start time" span={2}>
            {ele.startTime}
          </Descriptions.Item>
          <Descriptions.Item label="End time" span={2}>
            {ele.endTime}
          </Descriptions.Item>
          <Descriptions.Item label="Status" span={2}>
            <Badge status={ele.status} text={ele.status} />
          </Descriptions.Item>
        </Descriptions>
      ))}
    </>
  );
};

export default History;
