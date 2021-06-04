import React from "react";
import { Card } from "antd";
import { EditOutlined, SettingOutlined } from "@ant-design/icons";
import Macy from "macy";
import "./index.less";
class Dashboard extends React.Component {
  constructor(props) {
    super(props);
    this.testRef = React.createRef();
    this.state = {};
  }

  getMacy = () => {
    if (this.state.masonry) {
      this.state.masonry.reInit();
    } else {
      let masonry = new Macy({
        container: ".macy-container",
        trueOrder: false,
        debug: true,
        margin: { x: 60, y: 30 },
        columns: 2,
      });
      this.setState({ masonry });
    }
  };
  render() {
    return (
      <div className="dashboard-wrapper">
        <div className="macy-container">
          <Card
            hoverable
            className=".macy-item"
            title="Events"
            style={{
              width: 300,
            }}
          >
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
          </Card>
          <Card
            hoverable
            className=".macy-item"
            title="Recommendation news feed"
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

export default Dashboard;
