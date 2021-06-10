import React from 'react';
import { withRouter } from 'react-router';
import { connect } from 'react-redux'
import { Form, Input, Select, Button } from 'antd';
import { topicActions } from '../../../../redux/actions/topicAction';
import { postActions } from '../../../../redux/actions/postAction';
import './CreatePost.css';

const layout = {
    labelCol: { span: 4 },
    wrapperCol: { span: 16 },
};

const tailLayout = {    
    wrapperCol: { offset: 10, span: 16 },
}

const validateMessages = {
    required: '${label} is required!'    
};

const PAYLOAD = {
    title: 'postName',
    topic: 'subredditName',
    url: 'url',
    description: 'description'
};

class CreatePost extends React.Component { 
    formRef = React.createRef();
    state = {
        topic: '',
        postTitle: '',
        uploadImgUrl: '',
        postContent: ''
    };    

    componentDidMount() {                
        this.props.getAllTopics();        
    }    

    onReset = () => {
        this.formRef.current.resetFields();     
    }

    onFormChange(val, type) {
        this.setState({ [type]: val});        
    }

    submitPost() {                      
        if (this.isValid()) {
            const redirectTo = '/discussion/postList';
            const payload = {
                [PAYLOAD.topic]: this.state.topic,
                [PAYLOAD.title]: this.state.postTitle,
                [PAYLOAD.url]: this.state.uploadImgUrl,
                [PAYLOAD.description]: this.state.postContent
            };
            this.props.createPost(payload, this.props.history, redirectTo);
        }
    }

    isValid() {
        return this.state.topic !== '' && this.state.postTitle !== '';
    }

    render() {
        return (
            <div className="postForm-container">                
                <Form {...layout} ref={this.formRef} name="postForm" 
                    className="postForm" validateMessages={validateMessages}>
                    <Form.Item name='topic' label="Topic" rules={[{ required: true }]}>
                        <Select
                            showSearch
                            placeholder="Please Select a Topic"
                            optionFilterProp="label"
                            onChange={(val) => this.onFormChange(val, 'topic')}
                            options={this.props.topic.topicList.map((topic) => { return { label: topic.name, value: topic.name } })}                                         
                            filterOption={(input, option) => option.label.toLowerCase().indexOf(input.toLowerCase()) >= 0}>                            
                        </Select>
                    </Form.Item>
                    <Form.Item name='postTitle' label="Title" rules={[{ required: true }]}>
                        <Input onChange={(e) => this.onFormChange(e.target.value, 'postTitle')}/>
                    </Form.Item>
                    <Form.Item name='uploadImgUrl' label="Image URL">
                        <Input onChange={(e) => this.onFormChange(e.target.value, 'uploadImgUrl')} />
                    </Form.Item>                                            
                    <Form.Item name='postContent' label="Content">
                        <Input.TextArea onChange={(e) => this.onFormChange(e.target.value, 'postContent')} />
                    </Form.Item>   
                    <Form.Item {...tailLayout}>
                        <Button type="primary" htmlType="submit" onClick={() => this.submitPost()}>Submit</Button>
                        <Button htmlType="button" onClick={this.onReset}>Reset</Button>
                    </Form.Item>
                </Form>   
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return {        
        topic: state.getAllTopicReducer
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        getAllTopics: () => dispatch(topicActions.getTopicList()),
        createPost: (payload, history, redirectTo) => dispatch(postActions.createPost(payload, history, redirectTo))
    };
}

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(CreatePost));
