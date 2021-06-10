import { combineReducers } from "redux";
import testReducer from "./testReducer";
import { postReducer, topicReducer } from './discussReducer';
import { getAllPostReducer } from './getPostListReducer';
import { getAllTopicReducer } from './getTopicListReducer';

export default combineReducers({
  testData: testReducer,
  postReducer,
  topicReducer,
  getAllPostReducer,
  getAllTopicReducer
});
