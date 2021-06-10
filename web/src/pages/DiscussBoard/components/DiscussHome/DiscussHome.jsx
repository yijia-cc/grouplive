import React from 'react';
import { Route, Switch, withRouter, Redirect } from "react-router-dom";
import CreatePost from '../CreatePost/CreatePost';
import CreateTopic from '../CreateTopic/CreateTopic';
import DiscussMain from '../DiscussMain/DiscussMain';
import './DiscussHome.css';


class DiscussHome extends React.Component {    
    render() {
        const { path } = this.props.match;
        return (
            <Switch>
                <Route path={`${path}/home`} component={DiscussMain} />           
                <Route path={`${path}/posts`} component={CreatePost} />
                <Route path={`${path}/topics`} component={CreateTopic} />
                <Redirect to={`${path}/home`} />
            </Switch>            
        );
    }
}

export default withRouter(DiscussHome);