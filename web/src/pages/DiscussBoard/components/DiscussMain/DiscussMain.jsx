import React from 'react';
import { Layout, List, Avatar, Space, notification} from 'antd';
import { MessageOutlined, LikeOutlined, DislikeOutlined, StarOutlined } from '@ant-design/icons';
import * as $u from '../../services/DiscussService';
import { connect } from 'react-redux'
import { postActions } from '../../../../redux/actions/createPost';

const { Content } = Layout;
const avatar = 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png';
const altLogo = 'https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png';

const IconText = ({ icon, text }) => (
  <Space>
    <a onClick={() => iconClick(text)}>{React.createElement(icon)}</a>
    {text != 'like' && text != 'dislike' ? text : ''}
  </Space>
);

const VOTE = {
    like: 'UPVOTE',
    dislike: 'DOWNVOTE'
}

const iconClick = (text, id) => {        
    const payload = {
        voteType: VOTE[text],
        postId: id
    };
    return;
    $u.VOTE(payload).then(response => {
        // this.setState({  });
    });
}

const openNotificationWithIcon = (type, title, msg) => {
    notification[type]({
        message: title,
        description: msg 
    });
};

class DiscussMain extends React.Component {
    state = { postList: [] }
    componentDidMount() {
        this.props.getAllPost();
    }    

    componentDidUpdate() {
        // TODO...
        if (this.props.post.msg !== '') {
            openNotificationWithIcon(this.props.post.type, this.props.post.title, this.props.post.msg);            
        }
        if (this.props.topic.msg != '') {
            openNotificationWithIcon(this.props.topic.type, this.props.topic.title, this.props.topic.msg);
        }
    }

    render() {        
        return (
            <Layout style={{ padding: '24px 24px' }}>
                <Content
                    className="site-layout-background"
                    style={{
                    padding: 24,
                    margin: 0,
                    minHeight: 280,
                    }}>
                    <List
                        itemLayout="vertical"
                        size="large"
                        pagination={{
                            onChange: page => {
                                console.log(page);
                            },
                            pageSize: 5,
                        }}
                        dataSource={this.props.post.postList}                    
                        renderItem={item => (
                            <List.Item
                                key={item.name}
                                actions={[
                                    <IconText icon={StarOutlined} text={item.voteCount} key="list-vertical-star-o" />,
                                    <IconText icon={LikeOutlined} text="like" key="list-vertical-like-o" onClick={() => this.iconClick('like')}/>,
                                    <IconText icon={DislikeOutlined} text="dislike" key="list-vertical-dislike-o" onClick={() => this.iconClick('dislike')} />,
                                    <IconText icon={MessageOutlined} text={item.commentCount} key="list-vertical-message" />,
                                ]}
                                extra={
                                    <img
                                        width={272}
                                        alt="logo"
                                        style={{maxHeight: 200}}
                                        src={item.url}/>
                                }>
                                <List.Item.Meta
                                    avatar={<Avatar src={avatar} />}
                                    title={<a href={item.href}>{item.subredditName}</a>}
                                    description={item.postName}/>
                                    {item.description}
                            </List.Item>
                        )}
                    />
                </Content>
            </Layout>
        );
    }
}

const mapStateToProps = (state) => {
    return {
        post: state.getAllPostReducer,
        topic: state.getAllTopicReducer
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        getAllPost: () => dispatch(postActions.getPostList())
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(DiscussMain);