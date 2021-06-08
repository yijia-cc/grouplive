import React from 'react';
import { connect } from 'react-redux'
import { Layout, List, Avatar, Space, notification } from 'antd';
import { LikeOutlined, LikeFilled, DislikeOutlined, DislikeFilled, MessageOutlined, StarOutlined } from '@ant-design/icons';
import { postActions } from '../../../../redux/actions/postAction';
import { voteActions } from '../../../../redux/actions/voteAction';
import { ACTION_TYPES } from '../../../../redux/actions/actionType';
import './DiscussMain.css';

const { Content } = Layout;
const avatar = 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png';
const altLogo = 'https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png';

const IconText = ({ icon, text, item, that }) => {    
    return (        
        <Space>
            <a onClick={() => that.iconClick(text, item, that, icon)}>
                {React.createElement(icon) }                 
            </a>     
            {text !== 'like' && text !== 'dislike' ? text : ''}                
        </Space>
    );    
};

const VOTE = {
    like: 'UPVOTE',
    dislike: 'DOWNVOTE'
}

const POST_ICON_MAP = {    
    LIKE_ICON: LikeOutlined,
    LIKE_CLICKED_ICON: LikeFilled,        
    DISLIKE_ICON: DislikeOutlined,
    DISLIKE_CLICKED_ICON: DislikeFilled
}

const openNotificationWithIcon = (type, title, msg) => {
    notification[type]({
        message: title,
        description: msg        
    });
};

class DiscussMain extends React.Component {
    state = {
        selectedItem: null
    };
    
    componentDidMount() {
        if (this.props.newPost.title == 'Post' && this.props.newPost.msg !== '') {
            openNotificationWithIcon(this.props.newPost.type, this.props.newPost.title, this.props.newPost.msg);                        
        } else if (this.props.newTopic.title === 'Topic' && this.props.newTopic.msg != '') {
            openNotificationWithIcon(this.props.newTopic.type, this.props.newTopic.title, this.props.newTopic.msg);            
        }
        this.props.getAllPost();  
    }    

    componentDidUpdate() {
        if (this.props.vote.isSuccess && 
            this.state.selectedItem !== null && 
            this.state.selectedItem.id === this.props.vote.item.id && 
            this.state.selectedItem.voteCount != this.props.vote.item.voteCount) {
            this.setState({ selectedItem: this.props.vote.item });
        }
    }

    componentWillUnmount() {
        this.props.dispatch({ type: ACTION_TYPES.RESET_NEW_POST });
        this.props.dispatch({ type: ACTION_TYPES.RESET_NEW_TOPIC });
    }

    iconClick = (text, item, that) => {
        this.setState({ selectedItem: item });
        if (text === 'like' || text === 'dislike') {
            const payload = {
                voteType: VOTE[text],
                postId: item.id
            };        
            that.props.voting(payload, item, text);
        }
    }
   
    render() {        
        return (
            <Layout className="discuss-main-layout">
                <Content className="site-layout-background discuss-main-content">
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
                                key={item.id}
                                actions={[
                                    <IconText icon={StarOutlined} text={item.voteCount} item={item} that={this} key="voteCnt" />,
                                    <IconText icon={item.upVote ? POST_ICON_MAP.LIKE_CLICKED_ICON : POST_ICON_MAP.LIKE_ICON} text="like" item={item} that={this} key='like' />,                                    
                                    <IconText icon={item.downVote ? POST_ICON_MAP.DISLIKE_CLICKED_ICON : POST_ICON_MAP.DISLIKE_ICON} text="dislike" item={item} that={this} key='dislike' />,
                                    <IconText icon={MessageOutlined} text={item.commentCount} item={item} that={this} key="commentsCnt" />,
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
                                    title={<a href={`/discussion/posts/${item.id}`}>{item.subredditName}</a>}
                                    description={item.postName}/>
                                    { item.description }
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
        topic: state.getAllTopicReducer,
        newPost: state.newPostReducer,
        newTopic: state.newTopicReducer,
        vote: state.voteReducer
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        getAllPost: () => dispatch(postActions.getPostList()),        
        voting: (vote, item, option) => dispatch(voteActions.voting(vote, item, option)),
        dispatch
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(DiscussMain);