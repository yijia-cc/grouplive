import { createStore, applyMiddleware } from "redux";
import { composeWithDevTools } from "redux-devtools-extension"
import thunkmiddleware from 'redux-thunk'
import reducers from "./reducers";
export default createStore(reducers, applyMiddleware(thunkmiddleware), composeWithDevTools());
