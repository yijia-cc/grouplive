import { Route, Switch, Redirect } from "react-router-dom";
import { Layout } from "antd";
import Login from "../../pages/Login";
import Calendar from "../../pages/Calendar";
import ChatRoom from "../../pages/ChatRoom";
import Dashboard from "../../pages/Dashboard";
import DiscussBoard from "../../pages/DiscussBoard";
import UserInfo from "../../pages/UserInfo";
import Payment from "../../pages/Payment";
import "./index.css";
const { Content } = Layout;

const Main = (props) => {
  const { isLoggedIn, handleLoggedIn } = props;
  const showLogin = () => {
    return isLoggedIn ? (
      <Redirect to="/dashboard" />
    ) : (
      <Login handleLoggedIn={handleLoggedIn} />
    );
  };
  const showDashBoard = () => {
    return isLoggedIn ? <Dashboard /> : <Redirect to="/login" />;
  };
  const showDiscussBoard = () => {
    return isLoggedIn ? <DiscussBoard /> : <Redirect to="/login" />;
  };
  const showChatRoom = () => {
    return isLoggedIn ? <ChatRoom /> : <Redirect to="/login" />;
  };
  const showCalendar = () => {
    return isLoggedIn ? <Calendar /> : <Redirect to="/login" />;
  };
  const showUserInfo = () => {
    return isLoggedIn ? <UserInfo /> : <Redirect to="/login" />;
  };
  const showPayment = () => {
    return isLoggedIn ? <Payment /> : <Redirect to="/login" />;
  };
  return (
    <Layout style={{ padding: "112px 50px 24px" }}>
      <Content
        className="site-layout-background"
        style={{
          padding: 24,
          margin: 0,
          minHeight: 1200,
        }}
      >
        <Switch>
          <Route path="/" exact render={showLogin} />
          <Route path="/login" render={showLogin} />
          <Route path="/dashboard" render={showDashBoard} />
          <Route path="/discussion" render={showDiscussBoard} />
          <Route path="/chat" render={showChatRoom} />
          <Route path="/calendar" render={showCalendar} />
          <Route path="/userinfo" render={showUserInfo} />
          <Route path="/payment" render={showPayment} />
          <Redirect to="/login" />
        </Switch>
      </Content>
    </Layout>
  );
};

export default Main;
