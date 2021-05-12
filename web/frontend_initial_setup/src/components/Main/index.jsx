import { Route, Switch, Redirect, withRouter } from "react-router-dom";
import { Layout, Breadcrumb } from "antd";
import CalendarSch from "../../pages/CalendarSch";
import ChatRoom from "../../pages/ChatRoom";
import Dashboard from "../../pages/Dashboard";
import DiscussBoard from "../../pages/DiscussBoard";

const { Content } = Layout;
const routeMap = {
  dashboard: "Dashboard",
  discussboard: "Discussion board",
  chatRoom: "Char room",
  Calendar: "Calendar schedule",
};
const Main = (props) => {
  const { pathname } = props.location;
  const keyArray = pathname.split("/");
  const breadcrumbValue =
    keyArray[1] === undefined ? routeMap.dashboard : routeMap[keyArray[1]];
  return (
    <Layout style={{ padding: "0 24px 24px" }}>
      <Breadcrumb style={{ margin: "16px 0" }}>
        <Breadcrumb.Item>{breadcrumbValue}</Breadcrumb.Item>
      </Breadcrumb>
      <Content
        className="site-layout-background"
        style={{
          padding: 24,
          margin: 0,
          minHeight: 280,
        }}
      >
        <Switch>
          <Route path="/dashboard" component={Dashboard} />
          <Route path="/discussboard" component={DiscussBoard} />
          <Route path="/ChatRoom" component={ChatRoom} />
          <Route path="/Calendar" component={CalendarSch} />
          <Redirect to="/dashboard" />
        </Switch>
      </Content>
    </Layout>
  );
};

export default withRouter(Main);
