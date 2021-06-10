import { VOTE } from '../../pages/DiscussBoard/services/DiscussService';
import { ACTION_TYPES } from './actionType';

const voting = (payload, item, option) => {    
    return dispatch => {
        return VOTE(payload).then(statusCd => {  
            if (statusCd === 200) {
                if (option === 'like') {                    
                    if (item.upVote) {
                        // case1: duplicate like voteType
                        item.voteUpCount -= 1;
                        item.upVote = !item.upVote;
                    } else {
                        // case2: not like before
                        item.voteUpCount += 1;
                        item.upVote = !item.upVote;
                        if (item.downVote) { // case2.1: previous dislike
                            item.downVote = !item.downVote;
                            item.voteDownCount -= 1;
                        }
                    }       
                } else {
                    if (item.downVote) {
                        // case1: duplicate dislike voteType
                        item.voteDownCount -= 1;
                        item.downVote = !item.downVote;
                    } else {
                        // case2: not like before
                        item.voteDownCount += 1;
                        item.downVote = !item.downVote;
                        if (item.upVote) { // case2.1: previous dislike
                            item.upVote = !item.upVote;
                            item.voteUpCount -= 1;
                        }
                    }  
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
