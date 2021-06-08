import { VOTE } from '../../pages/DiscussBoard/services/DiscussService';
import { ACTION_TYPES } from './actionType';

const voting = (payload, item, option) => {    
    return dispatch => {
        return VOTE(payload).then(statusCd => {  
            if (statusCd === 200) {
                if (option === 'like') {
                    item.voteCount += 1;
                    item.upVote = true;
                    item.downVote = false;
                } else {
                    item.voteCount -= 1;
                    item.downVote = true;
                    item.upVote = false;
                }
                const data = {
                    isSuccess: true,
                    item
                }
                dispatch({ type: ACTION_TYPES.VOTE_SUCCESS, data });
            }                  
        }, (err) => Promise.reject(err));
    }
}

export const voteActions = { voting };
