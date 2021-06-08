import React from 'react';
import { NavLink, withRouter } from "react-router-dom";
import { Menu } from 'antd';
import { HomeOutlined, FormOutlined } from '@ant-design/icons';

class DiscussHeader extends React.Component {    
    selectedKey = 'postList';
    render() {
        const { path } = this.props.match;        
        const curUrl = window.location.href;
        this.selectedKey = curUrl.substring(curUrl.lastIndexOf('/') + 1);
        return (
            <>
                <Menu
                    theme="light"
                    mode="horizontal"
                    defaultSelectedKeys={['postList']}
                    selectedKeys={[this.selectedKey]}>
                    <Menu.Item icon={<HomeOutlined />} key="postList">
                        Post List
                        <NavLink to={`${path}/postList`} />
                    </Menu.Item>
                    <Menu.Item icon={<FormOutlined />} key="create-topic">                        
                        Create Topic
                        <NavLink to={`${path}/create-topic`} />
                    </Menu.Item>
                    <Menu.Item icon={<FormOutlined />} key="create-post">                        
                        Create Post
                        <NavLink to={`${path}/create-post`} />
                    </Menu.Item>
                </Menu>                
            </>
        );
    }
}

export default withRouter(DiscussHeader);