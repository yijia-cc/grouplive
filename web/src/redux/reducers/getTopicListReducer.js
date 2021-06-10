import { ACTION_TYPES } from '../actions/actionType'

const initState = {
    topicList: [],
    isSuccess: false,
    isLoading: false,
    msg: '',
    type: '',
    title: ''
};

export const getAllTopicReducer = (state = initState, action) => {
    const { data, type } = action;
    let succState;    
    switch (type) {
        case ACTION_TYPES.GET_ALL_TOPIC_SUCCESS:                              
            succState = {
                topicList: data,
                isSuccess: true,
                isLoading: false,
                msg: '',
                type: 'success',
                title: ''            
            }
            return succState;   
        case ACTION_TYPES.GET_UPDATE_ALL_TOPIC_SUCCESS:                              
            succState = {
                topicList: data,
                isSuccess: true,
                isLoading: false,
                msg: 'Your new Topic has been created.',
                type: 'success',
                title: 'Create Topic'            
            }
            return succState;  
        case ACTION_TYPES.GET_ALL_TOPIC_FAILED:                              
            const failedState = {
                topicList: data,
                isSuccess: false,
                isLoading: false,
                msg: 'Opps, Something Went Wrong, Please try again later!!!',
                type: 'error',
                title: 'Create Topic'
            }
            return failedState;  
        default:
            return state;
    }
}