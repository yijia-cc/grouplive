import { Route, Switch, Redirect } from "react-router-dom";
import { Layout, Row, Col } from "antd";
import Calendar from "../../pages/Calendar";
import ChatRoom from "../../pages/ChatRoom";
import Dashboard from "../../pages/Dashboard";
import DiscussBoard from "../../pages/DiscussBoard";
import UserInfo from "../../pages/UserInfo";
import Payment from "../../pages/Payment";
import "./index.css";
const { Content } = Layout;

const Main = (props) => {
  return (
    <Content style={{ padding: "112px 50px 24px" }}>
      {/* <Col span={22}> */}
      <Content
        className="site-layout-background"
        style={{
          padding: 24,
          margin: 0,
          minHeight: 1000,
        }}
      >
        <Switch>
          <Route path="/dashboard" component={Dashboard} />
          <Route path="/discussion" component={DiscussBoard} />
          <Route path="/chat" component={ChatRoom} />
          <Route path="/calendar" component={Calendar} />
          <Route path="/userinfo" component={UserInfo} />
          <Route path="/payment" component={Payment} />
          <Redirect to="/dashboard" />
        </Switch>
      </Content>
      {/* </Col> */}
    </Content>
  );
};

export default Main;
