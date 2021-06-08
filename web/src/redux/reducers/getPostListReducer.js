import { ACTION_TYPES } from '../actions/actionType'

const initPostState = {
    postList: [],    
    msg: '',
    type: '',
    title: ''
};

export const getAllPostReducer = (state = initPostState, action) => {
    const { data, type } = action;            
    switch (type) {
        case ACTION_TYPES.GET_ALL_POSTS:                              
            state = {
                postList: data,                
                msg: '',
                type: 'success',
                title: ''            
            }
            return state;   
        default:
            return state;
    }
}