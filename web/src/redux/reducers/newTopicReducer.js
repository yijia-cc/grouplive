import { ACTION_TYPES } from "../actions/actionType";

const initState = {
    msg: '',
    title: '',
    type: ''  
};

export const newTopicReducer = (state = initState, action) => {
    const { type } = action;
    switch (type) {
        case ACTION_TYPES.NEW_TOPIC_CREATED:
            state = {
                type: 'success',
                title: 'Topic',
                msg: 'New topic has been created'
            }  
            return state;
        case ACTION_TYPES.NEW_TOPIC_FAILED:
            state = {
                type: 'failed',
                title: 'Topic',
                msg: 'New topic create failed'
            }
            return state;
        case ACTION_TYPES.RESET_NEW_TOPIC:            
            return initState;
        default:
            return initState;
    }    
};