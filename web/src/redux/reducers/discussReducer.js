import { ACTION_TYPES } from '../actions/actionType'
const postDefaultState = {
    title: {
        isValid: false,
        isRequired: true,     
        value: ''
    },
    url: {
        isValid: false,
        isRequired: false,
        value: ''
    },
    topic: {
        isValid: false,
        isRequired: true,
        value: ''
    },
    description: {
        isValid: false,
        isRequired: false, 
        value: ''
    },
    payload: {
        postName: '',
        url: '',
        subredditName: '',
        description: ''
    }, 
    postList: []    
}


export const postReducer = (state = postDefaultState, action) => {
    const { data, type } = action;
    switch (type) {
        case ACTION_TYPES.CREATE_POST:
            return data;        
        default:
            return postDefaultState;
    }
}

const topicDefaultState = {
    title: {
        isValid: false,
        isRequired: true,     
        value: ''
    },    
    description: {
        isValid: false,
        isRequired: false, 
        value: ''
    },
    payload: {
        name: '',
        description: ''
    }    
}

export const topicReducer = (state = topicDefaultState, action) => {
    const { data, type } = action;
    switch (type) {        
        case ACTION_TYPES.CREATE_TOPIC:
            return data;
        default:
            return topicDefaultState;
    }
}