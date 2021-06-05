import Chat from "./chat";
import { RecoilRoot } from "recoil";
import recoilPersist from "recoil-persist";
import "./index.css";

const ChatRoom = () => {
  return ( <div className="chatRoom-wrapper">
    <RecoilRoot initializeState={updateState}>
      <RecoilPersist />
      <Chat />
    </RecoilRoot>
  </div>)
 
};

const { RecoilPersist, updateState } = recoilPersist([], {
  key: "recoil-persist",
  storage: sessionStorage,
});

export default ChatRoom;
