import axios from "axios";
import { BASE_URL, TOKEN_KEY } from "../../../constants";

const ROOT_LOCAL_URL = 'http://localhost:8080';
const ROOT_URL = ROOT_LOCAL_URL;
const CREATE_POST_URL = `${ROOT_URL}/api/posts/`;
const GET_ALL_POSTS_SUCCESS_URL = `${ROOT_URL}/api/posts/`;
const GET_SINGLE_POST_URL = `${ROOT_URL}/api/posts/`;
const CREATE_TOPIC_URL = `${ROOT_URL}/api/subreddit`;
const GET_ALL_TOPICS_URL = `${ROOT_URL}/api/subreddit`;
const VOTE_URL = `${ROOT_URL}/api/votes/`;
const GET_POST_COMMENTS_URL = `${ROOT_URL}/api/comments/by-post/`;
const CREATE_COMMENT_URL = `${ROOT_URL}/api/comments/`;

const TOKEN = localStorage.getItem(TOKEN_KEY);
const axios_instance = axios.create();
axios_instance.defaults.headers.common['Authorization'] = TOKEN;

export const CREATE_POST = (payload) => {    
    return axios_instance.post(CREATE_POST_URL, payload)
            .then(res => res.data)
            .catch(error => Promise.reject(error));
}

export const GET_ALL_POSTS = () => {
    return axios_instance.get(GET_ALL_POSTS_SUCCESS_URL)
            .then(res => res.data)
            .catch(error => Promise.reject(error));;
}

export const GET_SINGLE_POST = (postId) => {
    return axios_instance.get(GET_SINGLE_POST_URL + postId)
            .then(res => res.data)
            .catch(error => Promise.reject(error));;
}

export const CREATE_TOPIC = (payload) => {
    return axios_instance.post(CREATE_TOPIC_URL, payload)
            .then(res => res.data)
            .catch(error => Promise.reject(error));
}

export const GET_ALL_TOPICS = () => {
    return axios_instance.get(GET_ALL_TOPICS_URL)
            .then(res => res.data)
            .catch(error => Promise.reject(error));
}

export const VOTE = (payload) => {
    return axios_instance.post(VOTE_URL, payload)
            .then(res => res.status)
            .catch(error => Promise.reject(error));
}

export const GET_POST_COMMENTS = (postId) => {
    return axios_instance.get(GET_POST_COMMENTS_URL + postId)
            .then(res => res.data)
            .catch(error => Promise.reject(error));
}

export const CREATE_COMMENT = (payload) => {
    return axios_instance.post(CREATE_COMMENT_URL, payload)
            .then(res => res.data)
            .catch(error => Promise.reject(error));
}