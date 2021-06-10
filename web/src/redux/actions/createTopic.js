import { CREATE_TOPIC, GET_ALL_TOPICS_SUCCESS } from '../../pages/DiscussBoard/services/DiscussService';
import { ACTION_TYPES } from './actionType';

/**
 * used for update topic list popup and rerender main discuss screen
 */
const getUpdateTopicList = (payload) => {
    return dispatch => {
        return CREATE_TOPIC(payload).then(response => {
            dispatch({ type: ACTION_TYPES.GET_UPDATE_ALL_TOPIC_SUCCESS, data: response });
        });
    }    
}

const getTopicList = () => {
    return dispatch => {
        return GET_ALL_TOPICS_SUCCESS().then(response => {
            dispatch({ type: ACTION_TYPES.GET_ALL_TOPIC_SUCCESS, data: response });
        });
    }    
}

export const topicActions = {    
    getUpdateTopicList,
    getTopicList
};