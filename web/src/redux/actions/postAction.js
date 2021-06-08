import { CREATE_POST, GET_ALL_POSTS, GET_SINGLE_POST, GET_POST_COMMENTS } from '../../pages/DiscussBoard/services/DiscussService';
import { ACTION_TYPES } from './actionType';

const createPost = (payload, history, redirectTo) => {
    return dispatch => {
        return CREATE_POST(payload).then(response => {
            dispatch({ type: ACTION_TYPES.NEW_POST_CREATED, data: response });
            history.push(redirectTo);
        }, err => {
            return Promise.reject(err, 'New post create failed.');
        });
    }  
}

const getPostList = () => {
    return dispatch => {
        return GET_ALL_POSTS().then(response => {
            dispatch({ type: ACTION_TYPES.GET_ALL_POSTS, data: response });
        });
    }    
}

const getSinglePost = (postId) => {
    return dispatch => {
        return GET_SINGLE_POST(postId).then(response => {
            dispatch({ type: ACTION_TYPES.GET_SINGLE_POST, data: response });
        });
    }
}

const getPostComments = (postId) => {
    return dispatch => {
        return GET_POST_COMMENTS(postId).then(response => {
            dispatch({ type: ACTION_TYPES.GET_POST_COMMENTS, data: response });
        });
    }
}

export const postActions = {
    createPost,    
    getPostList,
    getSinglePost,    
    getPostComments
};