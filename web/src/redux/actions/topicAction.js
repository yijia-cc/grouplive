import { CREATE_TOPIC, GET_ALL_TOPICS } from '../../pages/DiscussBoard/services/DiscussService';
import { ACTION_TYPES } from './actionType';

const createTopic = (payload, history, redirectTo) => {
    return dispatch => {
        return CREATE_TOPIC(payload).then(response => {
            dispatch({ type: ACTION_TYPES.NEW_TOPIC_CREATED });
            history.push(redirectTo);
        }, err => {
            dispatch({ type: ACTION_TYPES.NEW_TOPIC_FAILED });
        });
    } 
}

const getTopicList = () => {
    return dispatch => {
        return GET_ALL_TOPICS().then(response => {
            dispatch({ type: ACTION_TYPES.GET_ALL_TOPICS, data: response });
        });
    }    
}

export const topicActions = {  
    createTopic,  
    getTopicList
};