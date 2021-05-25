import { Route, Switch, Redirect } from "react-router-dom";
import { Layout, Row, Col } from "antd";
import CalendarSchedule from "../../pages/CalendarSchedule";
import ChatRoom from "../../pages/ChatRoom";
import Dashboard from "../../pages/Dashboard";
import DiscussBoard from "../../pages/DiscussBoard";
import UserInfo from "../../pages/UserInfo";
import Payment from "../../pages/Payment";
import Login from "../../pages/Login/Login";
import Register from "../../pages/Register/Register";
import "./index.css";

const { Content } = Layout;

const Main = props => {
    const { isLoggedIn, handleLoggedIn } = props;

    console.log(props);

    const showLogin = () => {
        return isLoggedIn ? (
            <Redirect to="/login" />
        ) : (
            <Login handleLoggedIn={handleLoggedIn} />
        );
    };

    return (
        <Row justify="center" style={{ padding: "112px 0 24px" }}>
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
                        <Route path="/" exact render={showLogin} />
                        <Route path="/login" render={showLogin} />
                        <Route path="/register" component={Register} />

                        <Route path="/dashboard" component={Dashboard} />
                        <Route path="/discussion" component={DiscussBoard} />
                        <Route path="/chat" component={ChatRoom} />
                        <Route path="/calendar" component={CalendarSchedule} />
                        <Route path="/userinfo" component={UserInfo} />
                        <Route path="/payment" component={Payment} />
                        <Redirect to="/dashboard" />
                    </Switch>
                </Content>
            </Col>
        </Row>
        // </Layout>
    );
};

export default Main;
