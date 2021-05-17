import { Layout } from "antd";
import TopBar from "../TopBar";
import Aside from "../Aside";
import Main from "../Main";
import "./index.css";

const GroupLive = () => {
  return (
    <Layout>
      <TopBar />
      <Layout>
        <Aside />
        <Main />
      </Layout>
    </Layout>
  );
};

export default GroupLive;
