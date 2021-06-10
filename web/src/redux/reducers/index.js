import { combineReducers } from "redux";
import testReducer from "./testReducer";
import { getAllPostReducer } from './getPostListReducer';
import { getAllTopicReducer } from './getTopicListReducer';
import { getSinglePostReducer } from './getSinglePostReducer';
import { getPostCommentsReducer } from './getPostCommentsReducer';

import { newPostReducer } from './newPostReducer';
import { newTopicReducer } from './newTopicReducer';
import { newCommentsReducer } from './newCommentReducer';
import { voteReducer } from './voteReducer';

export default combineReducers({
  testData: testReducer,  
  getAllPostReducer,
  getAllTopicReducer,
  getSinglePostReducer,
  getPostCommentsReducer,
  newPostReducer,
  newTopicReducer,
  newCommentsReducer,
  voteReducer
});
