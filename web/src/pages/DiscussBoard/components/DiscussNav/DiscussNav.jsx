import React from 'react';
import { Route, Switch, withRouter, Redirect } from "react-router-dom";
import CreatePost from '../CreatePost/CreatePost';
import CreateTopic from '../CreateTopic/CreateTopic';
import PostDefail from '../PostDetail/PostDetail';
import DiscussMain from '../DiscussMain/DiscussMain';
import './DiscussNav.css';


class DiscussNav extends React.Component {    
    render() {
        const { path } = this.props.match;
        return (
            <Switch>
                <Route path={`${path}/postList`} component={DiscussMain} />           
                <Route path={`${path}/create-topic`} component={CreateTopic} />
                <Route path={`${path}/create-post`} component={CreatePost} />
                <Route path={`${path}/posts/:id`} component={PostDefail} />
                <Redirect to={`${path}/postList`} />
            </Switch>            
        );
    }
}

export default withRouter(DiscussNav);