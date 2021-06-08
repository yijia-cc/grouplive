import { ACTION_TYPES } from '../actions/actionType'

const initPostState = {
    data: {}
};

export const getSinglePostReducer = (state = initPostState, action) => {
    let { data, type } = action;            
    switch (type) {
        case ACTION_TYPES.GET_SINGLE_POST:                              
            state = {
                data        
            }
            return state;        
        default:
            return state;
    }
}