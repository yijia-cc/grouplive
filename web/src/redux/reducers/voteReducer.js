import { ACTION_TYPES } from "../actions/actionType";

const initVoteState = {
    isSuccess: false,
    postId: -1
};

export const voteReducer = (state = initVoteState, action) => {
    const { type, data } = action;
    switch (type) {
        case ACTION_TYPES.VOTE_SUCCESS:            
            state = {
                isSuccess: data.isSuccess,
                item: data.item
            }
            break;
        case ACTION_TYPES.VOTE_FAIL:
            state = {
                isSuccess: data.isSuccess,
                item: data.item
            }
            break;
        default:
            return state;    
    }    
    return state;
}