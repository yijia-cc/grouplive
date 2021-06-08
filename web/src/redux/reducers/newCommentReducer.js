import { ACTION_TYPES } from '../actions/actionType'

const initCommentState = {
    data: [],
    notificationFlag: false    
};

export const newCommentsReducer = (state = initCommentState, action) => {
    const { data, type } = action;          
    switch (type) {
        case ACTION_TYPES.NEW_COMMENT_SUCCESS:                              
            state = { 
                data,
                notificationFlag: false,
                type: 'success',
                title: 'Comment',
                msg: 'New comment has been created'
            }
            return state;
        case ACTION_TYPES.CREATE_COMMENT_NOTIFICATION:
            state = {
                data: [],
                notificationFlag: true
            }
            return state;
        case ACTION_TYPES.RESET_NEW_COMMENT:
            return state;
        default:
            return state;
    }
}