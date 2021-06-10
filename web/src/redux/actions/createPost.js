import { CREATE_POSTS, GET_ALL_POSTS_SUCCESS } from '../../pages/DiscussBoard/services/DiscussService';
import { ACTION_TYPES } from './actionType';

/**
 * used for after create post popup and rerender main discuss screen
 */
const getUpdatePostList = (payload) => {
    return dispatch => {
        return CREATE_POSTS(payload).then(response => {            
            dispatch({ type: ACTION_TYPES.GET_UPDATE_ALL_POSTS_SUCCESS, data: response });
        });
    }    
}  

const getPostList = () => {
    return dispatch => {
        return GET_ALL_POSTS_SUCCESS().then(response => {
            dispatch({ type: ACTION_TYPES.GET_ALL_POSTS_SUCCESS, data: response });
        });
    }    
}

export const postActions = {
    getUpdatePostList,
    getPostList    
};