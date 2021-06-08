import React from "react";
import { connect } from "react-redux";
import { List, Card, Button } from "antd";
import Macy from "macy";
import "./index.less";
class Dashboard extends React.Component {
  constructor(props) {
    super(props);
    this.testRef = React.createRef();
    this.state = {
      eventList: [
        {
          id: 1,
          title: `BBQ`,
          content: `Delicious BBq Event`,
          reserve: false,
          img: `https://images.ctfassets.net/86mn0qn5b7d0/featured-img-of-post-152543/436cf88c49f85eb46d1ab04748cbf8e6/featured-img-of-post-152543.jpg?w=1800&q=50&fm=jpg&fl=progressive`,
        },
        {
          id: 2,
          title: `Basketball game`,
          content: `3 V 3 basketball game`,
          reserve: false,
          img: `https://images.crazygames.com/games/basketball-stars-2019/cover-1583231506155.png?auto=format,compress&q=75&cs=strip`,
        },
        {
          id: 3,
          title: `Hollowean`,
          content: `Halloween evolved from the ancient Celtic holiday of Samhain.`,
          reserve: true,
          img: `https://ichef.bbci.co.uk/news/976/cpsprodpb/753F/production/_115151003_smallergettyimages-1184857940.jpg`,
        },
      ],
    };
  }

  getMacy = () => {
    if (this.state.masonry) {
      this.state.masonry.reInit();
    } else {
      let masonry = new Macy({
        container: ".macy-container",
        trueOrder: false,
        waitForImages: true,
        debug: true,
        margin: { x: 60, y: 30 },
        columns: 2,
      });
      this.setState({ masonry });
    }
  };
  serveAndUnserveHandler = (id) => {
    const { eventList } = this.state;
    const newEventList = eventList.map((ele) => {
      if (id === ele.id) {
        ele.reserve = !ele.reserve;
      }
      return ele;
    });
    this.setState({ eventList: newEventList });
  };
  render() {
    const { eventList } = this.state;
    const reservedEvent = eventList.filter((ele) => ele.reserve);
    return (
      <div className="dashboard-wrapper">
        <div className="macy-container">
          <Card
            hoverable
            className=".macy-item"
            title="Events"
            style={{
              width: 300,
              height: 600,
              overflow: "auto",
            }}
          >
            <List
              itemLayout="vertical"
              size="small"
              dataSource={eventList}
              renderItem={(item) => (
                <List.Item
                  key={item.title}
                  actions={[
                    <Button
                      type="primary"
                      onClick={() => {
                        this.serveAndUnserveHandler(item.id);
                      }}
                      danger={item.reserve ? true : false}
                    >
                      {item.reserve ? "Unreserve" : "Reserve"}
                    </Button>,
                  ]}
                  extra={<img width={272} alt="logo" src={item.img} />}
                >
                  <List.Item.Meta title={<div>{item.title}</div>} />
                  {item.content}
                </List.Item>
              )}
            />
          </Card>
          <Card
            hoverable
            className=".macy-item"
            title="Recommendation Event"
            style={{ width: 300 }}
          >
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
          </Card>
          <Card
            hoverable
            className=".macy-item"
            title="Announcement"
            style={{ width: 300 }}
          >
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
          </Card>
          <Card
            hoverable
            className=".macy-item"
            title="Reserved"
            style={{ width: 300 }}
          >
            <List
              itemLayout="vertical"
              size="small"
              dataSource={reservedEvent}
              renderItem={(item) => (
                <List.Item
                  key={item.title}
                  actions={[
                    <Button
                      type="primary"
                      onClick={() => {
                        this.serveAndUnserveHandler(item.id);
                      }}
                      danger={item.reserve ? true : false}
                    >
                      {item.reserve ? "Unreserve" : "Reserve"}
                    </Button>,
                  ]}
                  extra={<img width={272} alt="logo" src={item.img} />}
                >
                  <List.Item.Meta title={<div>{item.title}</div>} />
                  {item.content}
                </List.Item>
              )}
            />
          </Card>
          <Card
            hoverable
            className=".macy-item"
            title="Policy"
            style={{ width: 300 }}
          >
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
          </Card>
        </div>
      </div>
    );
  }
  componentDidMount() {
    this.getMacy();
  }
}

export default connect(
  (state) => ({
    amenityList: state.testData,
  }),
  {}
)(Dashboard);
