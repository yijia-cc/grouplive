import { Route, Switch, Redirect } from "react-router-dom";
import { Layout, Row, Col } from "antd";
import CalendarSchedule from "../../pages/CalendarSchedule";
import ChatRoom from "../../pages/ChatRoom";
import Dashboard from "../../pages/Dashboard";
import DiscussBoard from "../../pages/DiscussBoard";
import User from "../../pages/User";
import "./index.css";
const { Content } = Layout;

const Main = (props) => {
  return (
    // style={{ padding: "100px 24px 24px" }}
    <Layout style={{ padding: "100px 36px 24px" }}>
      <Row justify="center">
        <Col span={22}>
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
              <Route path="/calendar" component={CalendarSchedule} />
              <Route path="/user" component={User}></Route>
              <Redirect to="/dashboard" />
            </Switch>
          </Content>
        </Col>
      </Row>
    </Layout>
  );
};

export default Main;
