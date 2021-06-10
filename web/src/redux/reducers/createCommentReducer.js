import { ACTION_TYPES } from '../actions/actionType'

const initCommentState = {
    data: []
};

export const createCommentsReducer = (state = initCommentState, action) => {
    let { data, type } = action;    
    let succState;    
    switch (type) {
        case ACTION_TYPES.CREATE_COMMENT:                              
            succState = { data }
            return succState;        
        default:
            return state;
    }
}