import axios from "axios";

const ROOT_URL = 'http://localhost:8080';
const CREATE_POSTS_URL = `${ROOT_URL}/api/posts/`;
const GET_ALL_POSTS_SUCCESS_URL = `${ROOT_URL}/api/posts/`;
const CREATE_TOPIC_URL = `${ROOT_URL}/api/subreddit`;
const GET_ALL_TOPICS_URL = `${ROOT_URL}/api/subreddit`;
const VOTE_URL = `${ROOT_URL}/api/votes`;

export const CREATE_POSTS = (payload) => {
    return axios.post(CREATE_POSTS_URL, payload)
            .then(res => res.data)
            .catch(error => Promise.reject(error));
}

export const GET_ALL_POSTS_SUCCESS = () => {
    return axios.get(GET_ALL_POSTS_SUCCESS_URL)
            .then(res => res.data)
            .catch(error => Promise.reject(error));;
}

export const CREATE_TOPIC = (payload) => {
    return axios.post(CREATE_TOPIC_URL, payload)
            .then(res => res.data)
            .catch(error => Promise.reject(error));
}

export const GET_ALL_TOPICS_SUCCESS = () => {
    return axios.get(GET_ALL_TOPICS_URL)
            .then(res => res.data)
            .catch(error => Promise.reject(error));;
}

export const VOTE = (payload) => {
    return axios.post(VOTE_URL, payload);
}