import { ACTION_TYPES } from '../actions/actionType'

const initState = {
    postList: [],
    isSuccess: false,
    isLoading: false,
    msg: '',
    type: '',
    title: ''
};

export const getAllPostReducer = (state = initState, action) => {
    const { data, type } = action;
    if (data && data instanceof Array) {
        data.reverse();
    }
    let succState;
    switch (type) {
        case ACTION_TYPES.GET_ALL_POSTS_SUCCESS:                              
            succState = {
                postList: data,
                isSuccess: true,
                isLoading: false,
                msg: '',
                type: 'success',
                title: ''            
            }
            return succState;   
        case ACTION_TYPES.GET_UPDATE_ALL_POSTS_SUCCESS:                              
            succState = {
                postList: data,
                isSuccess: true,
                isLoading: false,
                msg: 'Your new Post has been created.',
                type: 'success',
                title: 'Create Post'            
            }
            return succState;  
        case ACTION_TYPES.GET_ALL_POSTS_FAILED:                              
            const failedState = {
                postList: data,
                isSuccess: false,
                isLoading: false,
                msg: 'Opps, Something Went Wrong, Please try again later!!!',
                type: 'error',
                title: 'Create Post'
            }
            return failedState;  
        default:
            return state;
    }
}