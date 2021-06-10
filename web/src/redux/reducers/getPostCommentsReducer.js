import { ACTION_TYPES } from '../actions/actionType'

const initCommentState = {
    data: [],
    isSuccess: false
};

export const getPostCommentsReducer = (state = initCommentState, action) => {
    let { data, type } = action;          
    switch (type) {
        case ACTION_TYPES.GET_POST_COMMENTS:                              
            state = { 
                data,
                isSuccess: true 
            }
            return state;
        default:
            return state;
    }
}