import { ACTION_TYPES } from "../actions/actionType";

const initState = {
    msg: '',
    title: '',
    type: ''  
};

export const newPostReducer = (state = initState, action) => {
    const { type } = action;
    switch (type) {
        case ACTION_TYPES.NEW_POST_CREATED:
            state = {
                type: 'success',
                title: 'Post',
                msg: 'New post has been created'
            }  
            return state;
        case ACTION_TYPES.NEW_POST_FAILED:
            state = {
                type: 'failed',
                title: 'Post',
                msg: 'New post create failed'
            }
            return state;
        case ACTION_TYPES.RESET_NEW_POST:            
            return initState;
        default:
            return initState;
    }    
};