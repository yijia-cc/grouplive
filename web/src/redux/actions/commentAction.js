import { CREATE_COMMENT } from '../../pages/DiscussBoard/services/DiscussService';
import { ACTION_TYPES } from './actionType';

const createComment = (payload) => {    
    return dispatch => {
        return CREATE_COMMENT(payload).then(response => {                               
            dispatch({ type: ACTION_TYPES.NEW_COMMENT_SUCCESS, data: response });
        })
    }
}

export const commentActions = {
    createComment
}